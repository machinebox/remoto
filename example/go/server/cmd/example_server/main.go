package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	example "github.com/matryer/remoto/example/go/server"
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

type greetFormatter struct{}

func (greetFormatter) Greet(ctx context.Context, req *example.GreetFormatRequest) (*example.GreetResponse, error) {
	resp := &example.GreetResponse{
		Greeting: fmt.Sprintf(req.Format, req.Name),
	}
	return resp, nil
}
