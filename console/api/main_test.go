package api

import (
	"io/ioutil"
	"log"
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
	is.NoErr(err)
	w := httptest.NewRecorder()
	h := handleRenderTemplate()
	h(w, req)
	is.Equal(w.Code, http.StatusOK)
	body := w.Body.String()
	log.Println(body)
	is.True(strings.Contains(body, "aaa"))
}

func readFile(t *testing.T, path string) string {
	is := is.New(t)
	b, err := ioutil.ReadFile(path)
	is.NoErr(err)
	return string(b)
}
