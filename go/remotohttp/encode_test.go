package remotohttp_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/matryer/is"
)

func TestEncode(t *testing.T) {
	is := is.New(t)
	data := struct {
		Greeting string `json:"greeting"`
	}{
		Greeting: "Hi there",
	}
	w := httptest.NewRecorder()
	err := remotohttp.Encode(w, nil, http.StatusOK, data)
	is.NoErr(err)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `{"greeting":"Hi there"}`)
	is.Equal(w.HeaderMap.Get("Content-Type"), "application/json; chatset=utf-8")
}
