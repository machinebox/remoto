package generator

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestParser(t *testing.T) {
	is := is.New(t)

	def, err := Parse("testdata/rpc/example")
	is.NoErr(err)

	is.Equal(len(def.Services), 2)
	is.Equal(def.PackageName, "greeter")
	out := def.String()
	is.Equal(out, `package greeter

// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	Greet(context.Context, *GreetFormatRequest) (*GreetResponse, error)
}

// GreetFormatRequest is the request for Greeter.GreetRequest.
type GreetFormatRequest struct {
	Format string
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
	Error string
}

// Greeter provides greeting services.
type Greeter interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	Greeting string
	Error string
}

`)

	is.Equal(def.PackageName, "greeter")
	is.Equal(def.Services[0].Name, "GreetFormatter")
	is.Equal(def.Services[0].Comment, "GreetFormatter provides formattable greeting services.")
	is.Equal(def.Services[0].Structures[0].Name, "GreetFormatRequest")
	is.Equal(def.Services[0].Structures[1].Name, "GreetResponse")
	is.Equal(def.Services[1].Name, "Greeter")
	is.Equal(def.Services[1].Comment, "Greeter provides greeting services.")
	is.Equal(def.Services[1].Structures[0].Name, "GreetRequest")
	is.Equal(def.Services[1].Structures[1].Name, "GreetResponse")
}

func TestErrors(t *testing.T) {
	is := is.New(t)
	tests := map[string]string{
		"testdata/rpc/errors/too-many-args":        "greeter.remoto.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/rpc/errors/no-variadic":          "greeter.remoto.go:8:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/rpc/errors/bad-first-arg":        "greeter.remoto.go:4:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/rpc/errors/too-few-return-args":  "greeter.remoto.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/rpc/errors/bad-return-args":      "greeter.remoto.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/rpc/errors/non-pointer-request":  "greeter.remoto.go:6:25: request object must be a pointer to a struct",
		"testdata/rpc/errors/non-pointer-response": "greeter.remoto.go:6:41: response object must be a pointer to a struct",
		"testdata/rpc/errors/bad-type":             "greeter.remoto.go:10:2: type int not supported: use explicitly sized types int32 or int64",
		"testdata/rpc/errors/unexported-fields":    "greeter.remoto.go:13:2: field name: must be exported",
		"testdata/rpc/errors/unexported-methods":   "greeter.remoto.go:8:2: method greet: must be exported",
	}
	pwd, err := os.Getwd()
	is.NoErr(err)
	for path, expectedErr := range tests {
		t.Run(path, func(t *testing.T) {
			is := is.New(t)
			os.Chdir(path)
			defer os.Chdir(pwd)
			_, err := Parse(".")
			is.True(err != nil) // must be an error
			is.Equal(err.Error(), expectedErr)
		})
	}
}
