// Package greeter is a sweet API that greets people.
package greeter

import "context"

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(context.Context, *GreetRequest) (*GreetRequest, error)
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}
