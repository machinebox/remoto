package testdata

type Greeter interface {
	Greet(GreetRequest) string
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
