package testdata

import "context"

type Greeter interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

type GreetRequest struct {
	Name int32
}

type GreetResponse struct {
	Greeting string
}
