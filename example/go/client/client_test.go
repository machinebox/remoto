package example_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	example "github.com/machinebox/remoto/example/go/client"
	exampleserver "github.com/machinebox/remoto/example/go/server"
	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)
	server := exampleserver.New(greetFormatter{}, greeter{})
	srv := httptest.NewServer(server)
	defer srv.Close()
	c := example.NewGreeterClient(srv.URL, http.DefaultClient)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp, err := c.Greet(ctx, &example.GreetRequest{
		Name: "Remoto",
	})
	is.NoErr(err)
	is.Equal(resp.Error, "")
	is.Equal(resp.Greeting, "Hello Remoto")
}

func TestMulti(t *testing.T) {
	is := is.New(t)
	server := exampleserver.New(greetFormatter{}, greeter{})
	srv := httptest.NewServer(server)
	defer srv.Close()
	c := example.NewGreeterClient(srv.URL, http.DefaultClient)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	req1 := &example.GreetRequest{
		Name: "Mat",
	}
	req2 := &example.GreetRequest{
		Name: "David",
	}
	req3 := &example.GreetRequest{
		Name: "Aaron",
	}
	resps, err := c.GreetMulti(ctx, req1, req2, req3)
	is.NoErr(err)
	is.Equal(len(resps), 3)
	is.Equal(resps[0].Greeting, "Hello Mat")
	is.Equal(resps[1].Greeting, "Hello David")
	is.Equal(resps[2].Greeting, "Hello Aaron")
}

type greeter struct{}

func (greeter) Greet(ctx context.Context, req *exampleserver.GreetRequest) (*exampleserver.GreetResponse, error) {
	resp := &exampleserver.GreetResponse{
		Greeting: "Hello " + req.Name,
	}
	return resp, nil
}

type greetFormatter struct{}

func (greetFormatter) Greet(ctx context.Context, req *exampleserver.GreetFormatRequest) (*exampleserver.GreetResponse, error) {
	resp := &exampleserver.GreetResponse{
		Greeting: fmt.Sprintf(req.Format, req.Name),
	}
	return resp, nil
}
