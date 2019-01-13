package remotohttp

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/machinebox/remoto/go/remotohttp/remototypes"
)

// Server is an HTTP server for serving Remoto requests.
type Server struct {
	handlers sync.Map

	// NotFound handles 404 responses.
	NotFound http.Handler

	// OnErr is called when there has been a system level error,
	// like encoding/decoding.
	OnErr func(w http.ResponseWriter, r *http.Request, err error)

	// NewClient gets a new http.Client. By default,
	// returns http.DefaultClient.
	NewClient func() *http.Client
}

// Register registers the path with the http.Handler.
func (srv *Server) Register(path string, fn http.Handler) {
	srv.handlers.Store(path, fn)
}

// ServeHTTP calls the registered handler
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if srv.NotFound != nil {
			srv.NotFound.ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
		return
	}
	h, ok := srv.handlers.Load(r.URL.Path)
	if !ok {
		srv.NotFound.ServeHTTP(w, r)
		return
	}
	handler, ok := h.(http.Handler)
	if !ok {
		panic("remotohttp: handler is the wrong type")
	}
	opener := func(_ context.Context, file remototypes.File) (io.ReadCloser, error) {
		f, _, err := r.FormFile(file.Fieldname)
		return f, err
	}
	r = r.WithContext(remototypes.WithOpener(r.Context(), opener))
	handler.ServeHTTP(w, r)
}

// Describe an overview of the endpoints to the specified io.Writer.
func (srv *Server) Describe(w io.Writer) error {
	var err error
	srv.handlers.Range(func(k, v interface{}) bool {
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
