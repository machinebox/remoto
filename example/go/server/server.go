// Code generated by Remoto; DO NOT EDIT.

// Package example contains the HTTP server for example services.
package example

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/pkg/errors"
	
	"github.com/machinebox/remoto/remototypes"	
	
)

// Run is the simplest way to run the services.
func Run(addr string,
	greetFormatter GreetFormatter,
	greeter Greeter,
) error {
	server := New(
		greetFormatter,
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
	greetFormatter GreetFormatter,
	greeter Greeter,
) *remotohttp.Server {
	server := &remotohttp.Server{
		OnErr: func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Fprintf(os.Stderr, "%s %s: %s\n", r.Method, r.URL.Path, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
	
	RegisterGreetFormatterServer(server, greetFormatter)
	RegisterGreeterServer(server, greeter)
	return server
}

// GreetRequest is the request for Greeter.Greet.
type GreetRequest struct {
	// 
	Name string `json:"name"`
	
}

// 
type GreetPhotoRequest struct {
	// 
	Photo remototypes.File `json:"photo"`
	// 
	Name string `json:"name"`
	
}

// 
type GreetPhotoResponse struct {
	// 
	Greeting string `json:"greeting"`
	// Error is an error message if one occurred.
	Error string `json:"error"`
	
}

// GreetFormatRequest is the request for GreetFormatter.Greet.
type GreetFormatRequest struct {
	// 
	Format string `json:"format"`
	// 
	Name string `json:"name"`
	
}

// GreetResponse is the response for Greeter.Greet and GreetFormatter.Greet.
type GreetResponse struct {
	// 
	Greeting string `json:"greeting"`
	// Error is an error message if one occurred.
	Error string `json:"error"`
	
}



// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	// 
	Greet(context.Context, *GreetFormatRequest) (*GreetResponse, error)

}

// RegisterGreetFormatterServer registers a GreetFormatter with a remotohttp.Server.
func RegisterGreetFormatterServer(server *remotohttp.Server, service GreetFormatter) {
	srv := &httpGreetFormatterServer{
		service: service,
		server: server,
	}
	
	server.Register("/remoto/GreetFormatter.Greet", http.HandlerFunc(srv.Greet))
	
}

type httpGreetFormatterServer struct {
	// service is the GreetFormatter being exposed by this
	// server.
	service GreetFormatter
	// server is the remotohttp.Server that this server is
	// registered with.
	server *remotohttp.Server
}


// Greet is an http.Handler wrapper for GreetFormatter.Greet.
func (srv *httpGreetFormatterServer) Greet(w http.ResponseWriter, r *http.Request) {
	var reqs []*GreetFormatRequest
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

// Greeter provides greeting services.
type Greeter interface {
	// 
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
// 
	GreetPhoto(context.Context, *GreetPhotoRequest) (*GreetPhotoResponse, error)

}

// RegisterGreeterServer registers a Greeter with a remotohttp.Server.
func RegisterGreeterServer(server *remotohttp.Server, service Greeter) {
	srv := &httpGreeterServer{
		service: service,
		server: server,
	}
	
	server.Register("/remoto/Greeter.Greet", http.HandlerFunc(srv.Greet))
	
	server.Register("/remoto/Greeter.GreetPhoto", http.HandlerFunc(srv.GreetPhoto))
	
}

type httpGreeterServer struct {
	// service is the Greeter being exposed by this
	// server.
	service Greeter
	// server is the remotohttp.Server that this server is
	// registered with.
	server *remotohttp.Server
}


// Greet is an http.Handler wrapper for Greeter.Greet.
func (srv *httpGreeterServer) Greet(w http.ResponseWriter, r *http.Request) {
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
// GreetPhoto is an http.Handler wrapper for Greeter.GreetPhoto.
func (srv *httpGreeterServer) GreetPhoto(w http.ResponseWriter, r *http.Request) {
	var reqs []*GreetPhotoRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}
	resps := make([]GreetPhotoResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.GreetPhoto(r.Context(), reqs[i])
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

