package generator

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

// func TestParser(t *testing.T) {
// 	is := is.New(t)

// 	def, err := Parse("testdata")
// 	is.NoErr(err)

// 	log.Println(def)

// }

func TestErrors(t *testing.T) {
	is := is.New(t)
	tests := map[string]string{
		"testdata/errors/too-many-args":        "greeter.rpc.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/no-variadic":          "greeter.rpc.go:8:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/bad-first-arg":        "greeter.rpc.go:4:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/too-few-return-args":  "greeter.rpc.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/bad-return-args":      "greeter.rpc.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/non-pointer-request":  "greeter.rpc.go:6:25: request object must be a pointer to a struct",
		"testdata/errors/non-pointer-response": "greeter.rpc.go:6:41: response object must be a pointer to a struct",
		"testdata/errors/bad-type":             "greeter.rpc.go:10:2: type int not supported: use explicitly sized types int32 or int64",
		"testdata/errors/unexported-fields":    "greeter.rpc.go:13:2: field name: must be exported",
		"testdata/errors/unexported-methods":   "greeter.rpc.go:8:2: method greet: must be exported",
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
