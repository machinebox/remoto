package remotohttp

import (
	"context"
	"net/http"
	"sync"
)

// Handler handles requests.
type Handler func(ctx context.Context) error

// Server is a Remoto HTTP server.
type Server struct {
	handlers sync.Map
}

// NewContext makes a new Context for the specified request.
func NewContext(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	return context.WithValue(ctx, contextKeyRemotoRequest, &remotoRequest{w: w, r: r})
}

// remotoRequest holds contextual information about the HTTP request.
type remotoRequest struct {
	w http.ResponseWriter
	r *http.Request
}

var (
	// contextKeyRemotoRequest is the context key for a Remoto HTTP request.
	contextKeyRemotoRequest = contextKey("contextKeyRemotoRequest")
)

type contextKey string

func (c contextKey) String() string {
	return "remoto context key: " + string(c)
}
