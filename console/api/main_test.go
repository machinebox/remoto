package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestRenderTemplate(t *testing.T) {
	is := is.New(t)
	q := url.Values{}
	q.Set("definition", readFile(t, "testdata/example.remoto.go"))
	req, err := http.NewRequest(http.MethodPost, "/api/templates/remotohttp/client.go", strings.NewReader(q.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	is.NoErr(err)
	w := httptest.NewRecorder()
	h := handleRenderTemplate()
	h(w, req)
	body := w.Body.String()

	is.Equal(w.Code, http.StatusOK)
	is.True(strings.Contains(body, "type GreeterClient struct"))
	is.True(strings.Contains(body, "type GreetRequest struct"))
	is.True(strings.Contains(body, "type GreetResponse struct"))
}

func readFile(t *testing.T, path string) string {
	is := is.New(t)
	b, err := ioutil.ReadFile(path)
	is.NoErr(err)
	return string(b)
}
