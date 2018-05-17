package testdata

import "context"

type Greeter interface {
	Greet(context.Context, *GreetRequest) *GreetResponse
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
