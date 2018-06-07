package remotohttp

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/machinebox/remoto/remototypes"
	"github.com/pkg/errors"
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
	opener := func(file remototypes.File) (io.ReadCloser, error) {
		s := string(file)
		switch {
		case strings.HasPrefix(s, "<remoto.File:"):
			file, _, err := r.FormFile(s[len("<remoto.File:") : len(s)-1])
			if err != nil {
				return nil, err
			}
			return file, nil
		case strings.HasPrefix(s, "<remoto.URL:"):
			client := srv.NewClient()
			if client == nil {
				client = http.DefaultClient
			}
			url := s[len("<remoto.URL:") : len(s)-1]
			resp, err := client.Get(url)
			if err != nil {
				return nil, err
			}
			if resp.StatusCode < 200 || resp.StatusCode >= 400 {
				resp.Body.Close()
				return nil, errors.New("GET " + url + ": " + resp.Status)
			}
			return resp.Body, nil
		}
		return nil, errors.New("bad File value")
	}
	ctx := remototypes.WithFileOpenContext(r.Context(), opener)
	r = r.WithContext(ctx)
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
