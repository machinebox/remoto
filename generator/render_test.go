package generator

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	def, err := Parse("testdata/rpc/example")
	is.NoErr(err)
	b, err := ioutil.ReadFile("testdata/templates/list.txt")
	is.NoErr(err)
	var buf bytes.Buffer
	err = Render(&buf, string(b), def)
	is.NoErr(err)
	is.Equal(buf.String(), `package: greeter
	service: GreetFormatter
		structure: GreetFormatRequest
		method: Greet
			field: Format string
			field: Name string
		structure: GreetResponse
		method: Greet
			field: Greeting string
			field: Error string
	service: Greeter
		structure: GreetRequest
		method: Greet
			field: Name string
		structure: GreetResponse
		method: Greet
			field: Greeting string
			field: Error string
`)
}
