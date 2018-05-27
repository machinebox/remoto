package remotohttpjson

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// HandlerFunc is a function that is executed in response to an incoming
// JSON/HTTP request.
type HandlerFunc func(ctx context.Context, w io.Writer, r io.Reader) error

// Server is a remoto JSON/HTTP server.
type Server struct {
	handlers sync.Map

	// NotFound handles 404 responses.
	NotFound http.Handler

	// OnErr is called when an internal server error occurs.
	// Example errors include encoding, reading and writing.
	OnErr func(err error)
}

// NewServer makes a new JSON/HTTP server.
func NewServer() *Server {
	return &Server{
		NotFound: http.NotFoundHandler(),
		OnErr:    func(err error) { fmt.Printf("remoto: %v\n", err) },
	}
}

// Register registers the Path with the HandlerFunc.
func (s *Server) Register(path string, fn HandlerFunc) {
	s.handlers.Store(path, fn)
}

// Describe an overview of the endpoints to the specified io.Writer.
func (s *Server) Describe(w io.Writer, prefix string) error {
	var err error
	s.handlers.Range(func(k, v interface{}) bool {
		if _, err = io.WriteString(w, prefix+" "+k.(string)+"\n"); err != nil {
			return false
		}
		return true
	})
	return err
}

// ServeHTTP servers the JSON/HTTP request registered to the path.
// Calls Server.NotFound if no path matches.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.NotFound.ServeHTTP(w, r)
		return
	}
	fnval, ok := s.handlers.Load(r.URL.Path)
	if !ok {
		s.NotFound.ServeHTTP(w, r)
		return
	}
	fn := fnval.(HandlerFunc)
	w.Header().Set("Content-Type", "application/json; chatset=utf-8")
	if err := fn(r.Context(), w, r.Body); err != nil {
		s.OnErr(err)
	}
}
