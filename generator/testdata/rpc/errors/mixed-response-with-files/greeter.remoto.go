// Package greeter is a sweet API that greets people.
package greeter

import (
	"context"

	"github.com/machinebox/remoto/remototypes"
)

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
	File     remototypes.File
}
