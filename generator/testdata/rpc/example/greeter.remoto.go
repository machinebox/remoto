// Package greeter is a sweet API that greets people.
package greeter

import "context"

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	// Greet generates a greeting.
	Greet(context.Context, *GreetFormatRequest) (*GreetResponse, error)
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
}

// GreetFormatRequest is the request for Greeter.GreetRequest.
type GreetFormatRequest struct {
	Format string
	Names  []string
}
