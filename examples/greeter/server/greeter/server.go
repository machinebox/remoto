// Code generated by Remoto; DO NOT EDIT.

// Package greeter contains the HTTP server for greeter services.
package greeter

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/machinebox/remoto/remototypes"
	"github.com/pkg/errors"
)

type Greeter interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// Run is the simplest way to run the services.
func Run(addr string,
	greeter Greeter,
) error {
	server := New(
		greeter,
	)
	if err := server.Describe(os.Stdout); err != nil {
		return errors.Wrap(err, "describe service")
	}
	if err := http.ListenAndServe(addr, server); err != nil {
		return err
	}
	return nil
}

// New makes a new remotohttp.Server with the specified services
// registered.
func New(
	greeter Greeter,
) *remotohttp.Server {
	server := &remotohttp.Server{
		OnErr: func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Fprintf(os.Stderr, "%s %s: %s\n", r.Method, r.URL.Path, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
		NotFound: http.NotFoundHandler(),
	}

	RegisterGreeterServer(server, greeter)
	return server
}

// RegisterGreeterServer registers a Greeter with a remotohttp.Server.
func RegisterGreeterServer(server *remotohttp.Server, service Greeter) {
	srv := &httpGreeterServer{
		service: service,
		server:  server,
	}
	server.Register("/remoto/Greeter.Greet", http.HandlerFunc(srv.handleGreet))

}

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Greeting string `json:"greeting"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// httpGreeterServer is an internal type that provides an
// HTTP wrapper around Greeter.
type httpGreeterServer struct {
	// service is the Greeter being exposed by this
	// server.
	service Greeter
	// server is the remotohttp.Server that this server is
	// registered with.
	server *remotohttp.Server
}

// handleGreet is an http.Handler wrapper for Greeter.Greet.
func (srv *httpGreeterServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	var reqs []*GreetRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]GreetResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.Greet(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// this is here so we don't get a compiler complaints.
func init() {
	var _ = remototypes.File{}
	var _ = strconv.Itoa(0)
	var _ = io.EOF
}
