package testdata

type Greeter interface {
	Greet(GreetRequest) (GreetResponse, error)
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
