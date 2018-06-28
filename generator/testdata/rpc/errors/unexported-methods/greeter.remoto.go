package testdata

// Greeter provides greeting services.
type Greeter interface {
	// greet generates a greeting.
	greet(GreetRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
}
