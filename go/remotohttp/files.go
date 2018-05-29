package remotohttp

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/machinebox/remoto/remototypes"
	"github.com/pkg/errors"
)

// OpenFile gets the file from the http.Request, or uses the client to download it.
// Callers must close the returned io.ReadCloser.
func OpenFile(ctx context.Context, file remototypes.File, r *http.Request, client *http.Client) (io.ReadCloser, error) {
	s := string(file)
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
