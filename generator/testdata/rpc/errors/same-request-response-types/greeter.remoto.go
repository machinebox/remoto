// Package greeter is a sweet API that greets people.
package greeter

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(GreetRequest) GreetRequest
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}
