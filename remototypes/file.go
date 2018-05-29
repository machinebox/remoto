package remototypes

import (
	"context"
	"errors"
	"io"
)

// File describes a binary file.
// Can be either:
//  <remoto.File:fieldname>
//  <remoto.URL:https://machinebox.io/>
// File: The file should be in the http.Request under fieldname.
// URL: The file will be downloaded from the URL.
type File string

// String gets a string representation of File.
func (f File) String() string {
	return string(f)
}

// NewFile creates a new File with the given fieldname.
func NewFile(fieldname string) File {
	return File("<remoto.File:" + fieldname + ">")
}

// NewFileURL create a new File with the given URL.
func NewFileURL(url string) File {
	return File("<remoto.URL:" + url + ">")
}

// Open opens the file.
func (f File) Open(ctx context.Context) (io.ReadCloser, error) {
	opener, ok := ctx.Value(contextKeyOpener).(OpenerFunc)
	if !ok {
		return nil, errors.New("unable to open (no opener found in context)")
	}
	return opener(f)
}

// OpenerFunc is a function capable of opening files.
// It is stored in the context and called by Open.
type OpenerFunc func(File) (io.ReadCloser, error)

// WithFileOpenContext specifies the opener function in the context.
func WithFileOpenContext(ctx context.Context, opener OpenerFunc) context.Context {
	return context.WithValue(ctx, contextKeyOpener, opener)
}

type contextKey string

func (c contextKey) String() string {
	return "remoto context key: " + string(c)
}

const contextKeyOpener = contextKey("opener")
