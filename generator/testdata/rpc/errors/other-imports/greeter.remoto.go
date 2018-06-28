// Package greeter is a sweet API that greets people.
package greeter

import (
	"context"
)

var _ = context.Background()

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	// Name is the name of the person to greet.
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	// Greeting is the personalized greeting.
	Greeting string
}
