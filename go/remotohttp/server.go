package remotohttp

import (
	"io"
	"net/http"
	"sync"
)

// Server is an HTTP server for serving Remoto requests.
type Server struct {
	handlers sync.Map

	// NotFound handles 404 responses.
	NotFound http.Handler

	// OnErr is called when there has been a system level error,
	// like encoding/decoding.
	OnErr func(w http.ResponseWriter, r *http.Request, err error)
}

// Register registers the path with the http.Handler.
func (s *Server) Register(path string, fn http.Handler) {
	s.handlers.Store(path, fn)
}

// ServeHTTP calls the registered handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, ok := s.handlers.Load(r.URL.Path)
	if !ok {
		s.NotFound.ServeHTTP(w, r)
		return
	}
	handler, ok := h.(http.Handler)
	if !ok {
		panic("remotohttp: handler is the wrong type")
	}
	handler.ServeHTTP(w, r)
}

// Describe an overview of the endpoints to the specified io.Writer.
func (s *Server) Describe(w io.Writer) error {
	var err error
	s.handlers.Range(func(k, v interface{}) bool {
		if _, err = io.WriteString(w, "endpoint: "+k.(string)+"\n"); err != nil {
			return false
		}
		return true
	})
	return err
}

// Error is an error wrapper for repsonses.
type Error struct {
	Err error `json:"error"`
}
