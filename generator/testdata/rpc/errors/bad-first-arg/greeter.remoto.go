package testdata

type Greeter interface {
	Greet() GreetResponse
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
