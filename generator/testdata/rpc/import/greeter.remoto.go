package greeter

import (
	"github.com/machinebox/remoto/remototypes"
)

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name  string
	Photo remototypes.File
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
}
