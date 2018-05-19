package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/matryer/remoto/play/pkg/greeter"
)

type service struct{}

func (service) Greet(ctx context.Context, req *greeter.GreetRequest) (greeter.GreetResponse, error) {
	resp := greeter.GreetResponse{
		Greeting: "Hello " + req.Name,
	}
	return resp, nil
}

func main() {
	var service service
	srv := greeter.NewServer(service)
	log.Println("listening on: ", srv.Addr)
	if err := srv.Run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
