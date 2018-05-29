package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	example "github.com/machinebox/remoto/example/go/server"
	"github.com/pkg/errors"
)

func main() {
	var (
		addr = flag.String("addr", ":8080", "listen address")
	)
	flag.Parse()
	fmt.Println("Remoto example server")
	fmt.Printf("listening on %v\n", *addr)
	if err := example.Run(*addr, greetFormatter{}, greeter{}); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type greeter struct{}

func (greeter) Greet(ctx context.Context, req *example.GreetRequest) (*example.GreetResponse, error) {
	resp := &example.GreetResponse{
		Greeting: "Hello " + req.Name,
	}
	return resp, nil
}

func (greeter) GreetPhoto(ctx context.Context, req *example.GreetPhotoRequest) (*example.GreetPhotoResponse, error) {
	f, err := req.Photo.Open(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}
	log.Println(string(b))
	resp := &example.GreetPhotoResponse{
		Greeting: "Hello " + req.Name,
	}
	return resp, nil
}

type greetFormatter struct{}

func (greetFormatter) Greet(ctx context.Context, req *example.GreetFormatRequest) (*example.GreetResponse, error) {
	resp := &example.GreetResponse{
		Greeting: fmt.Sprintf(req.Format, req.Name),
	}
	return resp, nil
}
