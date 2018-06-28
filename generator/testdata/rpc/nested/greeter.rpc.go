package testdata

type Greeter interface {
	Greet(GreetRequest) GreetResponse
}

type GreetRequest struct {
	Name Name
}

type GreetResponse struct {
	Greeting string
}

type Name struct {
	First string
	Last  string
}
