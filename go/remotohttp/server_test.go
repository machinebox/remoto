package remotohttp_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/matryer/is"
)

func TestServerServeHTTP(t *testing.T) {
	is := is.New(t)
	type greetRequest struct {
		Name string `json:"name"`
	}
	type greetResponse struct {
		Greeting string `json:"greeting"`
	}
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqs []greetRequest
		err := remotohttp.Decode(r, &reqs)
		is.NoErr(err)
		resps := make([]greetResponse, len(reqs))
		for i := range reqs {
			resps[i].Greeting = "Hello " + reqs[i].Name
		}
		err = remotohttp.Encode(w, r, http.StatusOK, resps)
		is.NoErr(err)
	})
	srv := &remotohttp.Server{}
	srv.Register("/remoto/Greeter.Greet", h)
	var reqs = []greetRequest{
		{Name: "Mat"},
		{Name: "David"},
		{Name: "Aaron"},
	}
	b, err := json.Marshal(reqs)
	is.NoErr(err)
	req, err := http.NewRequest(http.MethodPost, "/remoto/Greeter.Greet", bytes.NewReader(b))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.Body.String(), `[{"greeting":"Hello Mat"},{"greeting":"Hello David"},{"greeting":"Hello Aaron"}]`)
}

func TestServerDescribe(t *testing.T) {
	is := is.New(t)
	nilHandler := http.HandlerFunc(nil)
	srv := &remotohttp.Server{}
	srv.Register("/remoto/Service1.Method1", nilHandler)
	srv.Register("/remoto/Service2.Method2", nilHandler)
	srv.Register("/remoto/Service3.Method3", nilHandler)
	var buf bytes.Buffer
	err := srv.Describe(&buf)
	is.NoErr(err)
	s := buf.String()
	is.True(strings.Contains(s, "endpoint: /remoto/Service1.Method1"))
	is.True(strings.Contains(s, "endpoint: /remoto/Service2.Method2"))
	is.True(strings.Contains(s, "endpoint: /remoto/Service3.Method3"))
}
