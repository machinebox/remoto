package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/matryer/is"
)

func TestRender(t *testing.T) {
	is := is.New(t)
	def, err := Parse("testdata/rpc/example")
	is.NoErr(err)
	b, err := ioutil.ReadFile("testdata/templates/go-server.plush.go")
	is.NoErr(err)
	var buf bytes.Buffer
	err = Render(&buf, string(b), def)
	is.NoErr(err)
	fmt.Println(buf.String())
	is.Fail()
}
