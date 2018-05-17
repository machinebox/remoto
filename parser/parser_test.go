package parser

import (
	"log"
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestParser(t *testing.T) {
	is := is.New(t)

	def, err := Parse("testdata/nested")
	is.NoErr(err)

	log.Println(def)

}

func TestErrors(t *testing.T) {
	is := is.New(t)
	tests := map[string]string{
		"testdata/errors/too-many-args":        "greeter.rpc.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/bad-first-arg":        "greeter.rpc.go:4:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/too-few-return-args":  "greeter.rpc.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/bad-return-args":      "greeter.rpc.go:6:2: service methods must have signature (context.Context, *Request) (*Response, error)",
		"testdata/errors/non-pointer-request":  "greeter.rpc.go:6:25: request object must be a pointer to a struct",
		"testdata/errors/non-pointer-response": "greeter.rpc.go:6:41: response object must be a pointer to a struct",
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
