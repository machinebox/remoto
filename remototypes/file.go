package remototypes

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
)

// TODO: change this so it will work for clients too... where you can set an io.Reader for the file

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

// Open gets the file from the http.Request, or uses the client to download it.
// Callers must close the returned io.ReadCloser.
func (f File) Open(ctx context.Context, r *http.Request, client *http.Client) (io.ReadCloser, error) {
	s := string(f)
	switch {
	case strings.HasPrefix(s, "<remoto.File:"):
		file, _, err := r.FormFile(s[len("<remoto.File:") : len(s)-1])
		if err != nil {
			return nil, err
		}
		return file, nil
	case strings.HasPrefix(s, "<remoto.URL:"):
		url := s[len("<remoto.URL:") : len(s)-1]
		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			resp.Body.Close()
			return nil, errors.New("GET " + url + ": " + resp.Status)
		}
		return resp.Body, nil
	}
	return nil, errors.New("bad File value")
}
