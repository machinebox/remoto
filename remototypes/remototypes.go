package remototypes

import (
	"context"
	"errors"
	"io"
)

// File describes a binary file.
// This type is only allowed in requests, for responses RPC methods should
// return a FileResponse.
type File struct {
	Fieldname string `json:"fieldname"`
	Filename  string `json:"filename"`
}

// Open opens the file as an io.ReadCloser.
// Callers must close the file.
func (f File) Open(ctx context.Context) (io.ReadCloser, error) {
	opener, ok := ctx.Value(contextKeyFileOpener).(Opener)
	if !ok {
		return nil, errors.New("opener missing from context")
	}
	return opener(ctx, f)
}

// FileResponse is response type for a file.
type FileResponse struct {
	Filename      string    `json:"filename"`
	ContentType   string    `json:"contentType"`
	ContentLength int       `json:"contentLength"`
	Data          io.Reader `json:"-"`
	Error         string    `json:"error"`
}

// Opener is a function that knows how to open files.
type Opener func(ctx context.Context, file File) (io.ReadCloser, error)

// WithOpener gets a new context.Context with the specified Opener.
func WithOpener(ctx context.Context, opener Opener) context.Context {
	return context.WithValue(ctx, contextKeyFileOpener, opener)
}

// contextKey is a local context key type.
// see https://medium.com/@matryer/context-keys-in-go-5312346a868d
type contextKey string

func (c contextKey) String() string {
	return "remototypes context key: " + string(c)
}

// contextKeyFileOpener is the context key for a function capable of
// opening files.
var contextKeyFileOpener = contextKey("files")
