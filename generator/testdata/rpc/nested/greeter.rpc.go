package testdata

import "context"

type Greeter interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

type GreetRequest struct {
	Name Name
}

type GreetResponse struct {
	Greeting string
}

type Name struct {
	First string
	Last  string
}
