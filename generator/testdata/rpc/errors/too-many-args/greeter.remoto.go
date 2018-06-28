package testdata

type Greeter interface {
	Greet(GreetRequest, string) (GreetResponse, error)
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
