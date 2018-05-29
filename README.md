# remoto
Ultra-simple but complete RPC ecosystem.

## Definition

Definition files are Go source with `.remoto.go` file extension.

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
* Any slices (boundless arrays) of the supported types are also allowed (e.g. `[]string`, `[]bool`, etc.)
* Comments describe the service

### Special types

Remoto provides the following special types:

* `remototypes.File` - A local or remote file

### Tips

* For simplicity, don't import common types - just describe all the requireds types in a single file
