# remoto
Simple binary RPC ecosystem.

## Definition

Definition files are Go source with `.rpc.go` file extension.

An example definition looks like this:

```go
package project

import "context"

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
}
```

### Rules

* Each service is an `interface`
* Each method is an endpoint
* Methods must take `context.Context` as first argument, and pointer to request object as second argument
* Methods must return a pointer to the response object as the first argument, and an error as the second
* Only a subset of Go types are supported: `string`, `float64`, `int64`, `bool`, and `struct` types
* Comments describe the service
