package testdata

type Greeter interface {
	Greet(GreetRequest) GreetResponse
}

type GreetRequest struct {
	Name int32
}

type GreetResponse struct {
	Greeting string
}
