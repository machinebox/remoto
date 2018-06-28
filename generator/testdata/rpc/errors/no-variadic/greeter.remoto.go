package testdata

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(...GreetRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Names string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greetings string
}
