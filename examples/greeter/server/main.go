package main

import (
	"context"
	"fmt"
	"os"

	"github.com/machinebox/remoto/examples/greeter/server/greeter"
)

func main() {
	addr := "0.0.0.0:8080"
	fmt.Println("listening on", addr)
	if err := greeter.Run(addr, service{}); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type service struct{}

func (service) Greet(ctx context.Context, req *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	resp := &greeter.GreetResponse{
		Greeting: "Hello " + req.Name,
	}
	return resp, nil
}
