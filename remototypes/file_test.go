package remototypes

import (
	"context"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestFileString(t *testing.T) {
	is := is.New(t)
	is.Equal(NewFile("file").String(), "<remoto.File:file>")
	is.Equal(NewFileURL("https://machinebox.io/").String(), "<remoto.URL:https://machinebox.io/>")
}

func TestFileOpen(t *testing.T) {
	is := is.New(t)
	rc := ioutil.NopCloser(strings.NewReader("source"))
	opener := func(file File) (io.ReadCloser, error) {
		return rc, nil
	}
	ctx := context.Background()
	ctx = WithFileOpenContext(ctx, opener)
	f := NewFile("file")
	r, err := f.Open(ctx)
	is.NoErr(err)
	is.Equal(r, rc)
}
