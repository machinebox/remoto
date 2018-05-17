package testdata

import "context"

// Greeter provides greeting services.
type Greeter interface {
	// greet generates a greeting.
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
}
