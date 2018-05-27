package testdata

type Greeter interface {
	Greet(string, *GreetRequest) (*GreetResponse, error)
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
