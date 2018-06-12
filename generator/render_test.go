package generator

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	def, err := ParseDir("testdata/rpc/example")
	is.NoErr(err)
	b, err := ioutil.ReadFile("testdata/templates/list.txt")
	is.NoErr(err)
	var buf bytes.Buffer
	err = Render(&buf, "", string(b), def)
	is.NoErr(err)
	is.Equal(buf.String(), `package: greeter
	service: GreetFormatter
		method: Greet
			request: GreetFormatRequest
			response: GreetResponse
		structure: GreetingFormat
			field: Format string
			field: AllCaps bool
		structure: GreetFormatRequest
			field: Format GreetingFormat
			field: Names string
		structure: GreetResponse
			field: Greeting string
			field: Error string
	service: Greeter
		method: Greet
			request: GreetRequest
			response: GreetResponse
		structure: GreetRequest
			field: Name string
		structure: GreetResponse
			field: Greeting string
			field: Error string
`)
}
