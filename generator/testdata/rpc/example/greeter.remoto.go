// Package greeter is a sweet API that greets people.
package greeter

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(*GreetRequest) *GreetResponse
}

// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	// Greet generates a greeting.
	Greet(*GreetFormatRequest) *GreetResponse
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

// GreetFormatRequest is the request for Greeter.GreetRequest.
type GreetFormatRequest struct {
	// Format is a Go-style format string describing the greeting.
	// `%s` will be replaced with the name of the person.
	Format string
	// Names is one or more names of people to greet.
	Names []string
}
