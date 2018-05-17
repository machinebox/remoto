package testdata

import "context"

type Greeter interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

type GreetRequest struct {
	Name int
}

type GreetResponse struct {
	Greeting string
}
