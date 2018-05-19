package greeter

import (
	"context"
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Service interface {
	Greet(context.Context, *GreetRequest) (GreetResponse, error)
}

type Server struct {
	mux     *http.ServeMux
	rpc     *rpc.Server
	service Service

	// Addr is the address to listen on.
	// Default is ":8888"
	Addr string
}

func NewServer(service Service) *Server {
	srv := &Server{
		service: service,
		mux:     http.NewServeMux(),
		Addr:    ":8888",
	}
	srv.rpc = rpc.NewServer()
	srv.rpc.RegisterCodec(json.NewCodec(), "application/json")
	srv.rpc.RegisterService(srv, "")
	srv.mux.Handle("/rpc", srv.rpc)
	return srv
}

// Run runs the server.
func (s *Server) Run() error {
	if err := http.ListenAndServe(s.Addr, s.mux); err != nil {
		return err
	}
	return nil
}

func (s *Server) Greet(r *http.Request, req *GreetRequest, resp *GreetResponse) error {
	response, err := s.service.Greet(r.Context(), req)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
