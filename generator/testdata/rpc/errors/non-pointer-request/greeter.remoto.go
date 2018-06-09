package testdata

type Greeter interface {
	Greet(GreetRequest) *GreetResponse
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
