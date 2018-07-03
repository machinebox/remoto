package generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestParser(t *testing.T) {
	is := is.New(t)

	def, err := ParseDir("testdata/rpc/example")
	is.NoErr(err)

	is.Equal(len(def.Services), 2)
	is.Equal(def.PackageName, "greeter")
	is.Equal(def.PackageComment, "Package greeter is a sweet API that greets people.")
	is.Equal(def.Services[0].Name, "GreetFormatter")
	is.Equal(def.Services[0].Comment, "GreetFormatter provides formattable greeting services.")

	greetFormatRequest := def.Structure("GreetFormatRequest")
	is.True(greetFormatRequest != nil)
	is.Equal(greetFormatRequest.Name, "GreetFormatRequest")
	is.Equal(greetFormatRequest.Fields[1].Name, "Names")
	is.Equal(greetFormatRequest.Fields[1].Comment, "Names is one or more names of people to greet.")
	is.Equal(greetFormatRequest.Fields[1].Type.Name, "string")
	is.Equal(greetFormatRequest.Fields[1].Type.IsMultiple, true)

	is.Equal(def.Services[1].Name, "Greeter")
	is.Equal(def.Services[1].Comment, "Greeter provides greeting services.")

	is.Equal(def.Services[1].Methods[0].Name, "Greet")
	is.Equal(def.Services[1].Methods[0].Comment, "Greet generates a greeting.")

	greetRequest := def.Structure("GreetRequest")
	is.Equal(greetRequest.Name, "GreetRequest")
	is.Equal(greetRequest.Comment, "GreetRequest is the request for Greeter.GreetRequest.")
	is.Equal(greetRequest.Fields[0].Name, "Name")
	is.Equal(greetRequest.Fields[0].Comment, "Name is the name of the person to greet.")

	is.Equal(def.Services[1].Structures[1].Name, "GreetResponse")
	is.Equal(def.Services[1].Structures[1].Comment, "GreetResponse is the response for Greeter.GreetRequest.")

	greetingFormat := def.Structure("GreetingFormat")
	is.Equal(greetingFormat.Comment, "GreetingFormat describes the format of a greeting.")
	is.Equal(greetingFormat.Fields[0].Name, "Format")
	is.Equal(greetingFormat.Fields[0].Comment, "Format is a Go-style format string describing the greeting.\n%s will be replaced with the name of the person.")
	is.Equal(greetingFormat.Fields[0].Type.Name, "string")
	is.Equal(greetingFormat.Fields[1].Name, "AllCaps")
	is.Equal(greetingFormat.Fields[1].Comment, "AllCaps is whether to convert the greeting to all caps.")
	is.Equal(greetingFormat.Fields[1].Type.Name, "bool")

	out := def.String()
	is.Equal(out, `// Package greeter is a sweet API that greets people.
package greeter

// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	// Greet generates a greeting.
	Greet(GreetFormatRequest) GreetResponse
}

// GreetingFormat describes the format of a greeting.
type GreetingFormat struct {
	// Format is a Go-style format string describing the greeting.
	// %s will be replaced with the name of the person.
	Format string
	// AllCaps is whether to convert the greeting to all caps.
	AllCaps bool
}

// GreetFormatRequest is the request for Greeter.GreetRequest.
type GreetFormatRequest struct {
	// Format is the GreetingFormat describing the format
	// of the greetings.
	Format GreetingFormat
	// Names is one or more names of people to greet.
	Names []string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	// Greeting is the personalized greeting.
	Greeting string
	// Error is an error message if one occurred.
	Error string
}

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	// Name is the name of the person to greet.
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	// Greeting is the personalized greeting.
	Greeting string
	// Error is an error message if one occurred.
	Error string
}

`)
}

func TestParseReader(t *testing.T) {
	is := is.New(t)

	src := strings.NewReader(exampleRemotoDefinition)
	def, err := Parse(src)
	is.NoErr(err)

	is.Equal(len(def.Services), 2)
	is.Equal(def.PackageName, "greeter")
	is.Equal(def.PackageComment, "Package greeter is a sweet API that greets people.")
	is.Equal(def.Services[0].Name, "GreetFormatter")
	is.Equal(def.Services[0].Comment, "GreetFormatter provides formattable greeting services.")

	greetFormatRequest := def.Structure("GreetFormatRequest")
	is.Equal(greetFormatRequest.Name, "GreetFormatRequest")
	is.Equal(greetFormatRequest.Fields[1].Name, "Names")
	is.Equal(greetFormatRequest.Fields[1].Comment, "Names is one or more names of people to greet.")
	is.Equal(greetFormatRequest.Fields[1].Type.Name, "string")
	is.Equal(greetFormatRequest.Fields[1].Type.IsMultiple, true)

	is.Equal(def.Services[1].Name, "Greeter")
	is.Equal(def.Services[1].Comment, "Greeter provides greeting services.")

	is.Equal(def.Services[1].Methods[0].Name, "Greet")
	is.Equal(def.Services[1].Methods[0].Comment, "Greet generates a greeting.")

	greetRequest := def.Structure("GreetRequest")
	is.Equal(greetRequest.Name, "GreetRequest")
	is.Equal(greetRequest.Comment, "GreetRequest is the request for Greeter.GreetRequest.")
	is.Equal(greetRequest.Fields[0].Name, "Name")
	is.Equal(greetRequest.Fields[0].Comment, "Name is the name of the person to greet.")

	is.Equal(def.Services[1].Structures[1].Name, "GreetResponse")
	is.Equal(def.Services[1].Structures[1].Comment, "GreetResponse is the response for Greeter.GreetRequest.")

	greetingFormat := def.Structure("GreetingFormat")
	is.Equal(greetingFormat.Comment, "GreetingFormat describes the format of a greeting.")
	is.Equal(greetingFormat.Fields[0].Name, "Format")
	is.Equal(greetingFormat.Fields[0].Comment, "Format is a Go-style format string describing the greeting.\n%s will be replaced with the name of the person.")
	is.Equal(greetingFormat.Fields[0].Type.Name, "string")
	is.Equal(greetingFormat.Fields[1].Name, "AllCaps")
	is.Equal(greetingFormat.Fields[1].Comment, "AllCaps is whether to convert the greeting to all caps.")
	is.Equal(greetingFormat.Fields[1].Type.Name, "bool")

	out := def.String()
	is.Equal(out, `// Package greeter is a sweet API that greets people.
package greeter

// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	// Greet generates a greeting.
	Greet(GreetFormatRequest) GreetResponse
}

// GreetingFormat describes the format of a greeting.
type GreetingFormat struct {
	// Format is a Go-style format string describing the greeting.
	// %s will be replaced with the name of the person.
	Format string
	// AllCaps is whether to convert the greeting to all caps.
	AllCaps bool
}

// GreetFormatRequest is the request for Greeter.GreetRequest.
type GreetFormatRequest struct {
	// Format is the GreetingFormat describing the format
	// of the greetings.
	Format GreetingFormat
	// Names is one or more names of people to greet.
	Names []string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	// Greeting is the personalized greeting.
	Greeting string
	// Error is an error message if one occurred.
	Error string
}

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	// Name is the name of the person to greet.
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	// Greeting is the personalized greeting.
	Greeting string
	// Error is an error message if one occurred.
	Error string
}

`)
}

func TestErrors(t *testing.T) {
	is := is.New(t)
	tests := map[string]string{
		"testdata/rpc/errors/too-many-args":               "greeter.remoto.go:4:2: service methods must have signature (*Request) *Response",
		"testdata/rpc/errors/no-variadic":                 "greeter.remoto.go:6:2: service methods must have signature (*Request) *Response",
		"testdata/rpc/errors/bad-first-arg":               "greeter.remoto.go:4:2: service methods must have signature (*Request) *Response",
		"testdata/rpc/errors/too-few-return-args":         "greeter.remoto.go:4:2: service methods must have signature (*Request) *Response",
		"testdata/rpc/errors/bad-return-args":             "greeter.remoto.go:4:22: response object must be a named struct",
		"testdata/rpc/errors/pointer-request":             "greeter.remoto.go:4:8: request object must be a named struct (not a pointer - remove the *)",
		"testdata/rpc/errors/pointer-response":            "greeter.remoto.go:4:22: response object must be a named struct (not a pointer - remove the *)",
		"testdata/rpc/errors/bad-type":                    "greeter.remoto.go:8:2: type int32 not supported: use int",
		"testdata/rpc/errors/unexported-fields":           "greeter.remoto.go:11:2: field name: must be exported",
		"testdata/rpc/errors/unexported-methods":          "greeter.remoto.go:6:2: method greet: must be exported",
		"testdata/rpc/errors/same-request-response-types": "greeter.remoto.go:7:2: service methods must use different types for request and response objects",
		"testdata/rpc/errors/other-imports":               "import not allowed: context",
	}
	pwd, err := os.Getwd()
	is.NoErr(err)
	for path, expectedErr := range tests {
		t.Run(path, func(t *testing.T) {
			is := is.New(t)
			err := os.Chdir(path)
			is.NoErr(err)
			defer func() {
				err = os.Chdir(pwd)
				is.NoErr(err)
			}()
			_, err = ParseDir(".")
			is.True(err != nil) // must be an error
			is.Equal(err.Error(), expectedErr)
		})
	}
}

func TestParserImports(t *testing.T) {
	is := is.New(t)

	b, err := ioutil.ReadFile("testdata/rpc/import/greeter.remoto.go")
	is.NoErr(err)

	def, err := Parse(bytes.NewReader(b))
	is.NoErr(err)

	is.Equal(len(def.Services), 1)
	is.Equal(def.PackageName, "greeter")
	out := def.String()
	is.True(strings.Contains(out, `Photo remototypes.File`))
}

// exampleRemotoDefinition is a copy of testdata/rpc/example/greeter.remoto.go
// used for testing the io.Reader version of the parser.
var exampleRemotoDefinition = `// Package greeter is a sweet API that greets people.
package greeter

// Greeter provides greeting services.
type Greeter interface {
	// Greet generates a greeting.
	Greet(GreetRequest) GreetResponse
}

// GreetFormatter provides formattable greeting services.
type GreetFormatter interface {
	// Greet generates a greeting.
	Greet(GreetFormatRequest) GreetResponse
}

// GreetRequest is the request for Greeter.GreetRequest.
type GreetRequest struct {
	// Name is the name of the person to greet.
	Name string
}

// GreetResponse is the response for Greeter.GreetRequest.
type GreetResponse struct {
	// Greeting is the personalized greeting.
	Greeting string
}

// GreetFormatRequest is the request for Greeter.GreetRequest.
type GreetFormatRequest struct {
	// Format is the GreetingFormat describing the format
	// of the greetings.
	Format GreetingFormat
	// Names is one or more names of people to greet.
	Names []string
}

// GreetingFormat describes the format of a greeting.
type GreetingFormat struct {
	// Format is a Go-style format string describing the greeting.
	// %s will be replaced with the name of the person.
	Format string
	// AllCaps is whether to convert the greeting to all caps.
	AllCaps bool
}
`
