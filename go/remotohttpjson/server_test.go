package remotohttpjson_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/matryer/remoto/go/remotohttpjson"
	"github.com/pkg/errors"
)

func TestServer(t *testing.T) {
	is := is.New(t)
	type greetRequest struct {
		Name string `json:"name"`
	}
	type greetResponse struct {
		Greeting string `json:"name"`
		Error    string `json:"_error,omitempty"`
	}
	greet := func(ctx context.Context, r *greetRequest) (*greetResponse, error) {
		return &greetResponse{
			Greeting: fmt.Sprintf("hello %s", r.Name),
		}, nil
	}
	greetWrapper := func(ctx context.Context, w io.Writer, r io.Reader) error {
		var reqs []*greetRequest
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return errors.Wrap(err, "read request")
		}
		err = json.Unmarshal(b, &reqs)
		is.NoErr(err)
		resps := make([]*greetResponse, len(reqs))
		for i := range reqs {
			if resps[i], err = greet(ctx, reqs[i]); err != nil {
				resps[i].Error = err.Error()
			}
		}
		if b, err = json.Marshal(resps); err != nil {
			return errors.Wrap(err, "encode response")
		}
		if _, err = w.Write(b); err != nil {
			return errors.Wrap(err, "write response")
		}
		return nil
	}
	srv := remotohttpjson.NewServer()
	srv.Register("/remoto/greeter.Greet", greetWrapper)
	var buf bytes.Buffer
	is.NoErr(srv.Describe(&buf, "endpoint:"))
	is.Equal(buf.String(), `endpoint: /remoto/greeter.Greet`+"\n")
	in := `[{"name": "Mat"},
{"name": "David"},
{"name": "Aaron"}]`
	r, err := http.NewRequest(http.MethodPost, "/remoto/greeter.Greet", strings.NewReader(in))
	is.NoErr(err) // http.NewRequest
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK)
	is.Equal(w.HeaderMap.Get("Content-Type"), "application/json; chatset=utf-8")
	var responses []*greetResponse
	err = json.NewDecoder(w.Body).Decode(&responses)
	is.NoErr(err)
	is.Equal(len(responses), 3)
	is.Equal(responses[0].Greeting, "hello Mat")
	is.Equal(responses[1].Greeting, "hello David")
	is.Equal(responses[2].Greeting, "hello Aaron")
}

func TestServerErr(t *testing.T) {
	is := is.New(t)
	type greetRequest struct {
		Name string `json:"name"`
	}
	type greetResponse struct {
		Greeting string `json:"name"`
		Error    string `json:"_error,omitempty"`
	}
	greet := func(ctx context.Context, r *greetRequest) (*greetResponse, error) {
		return nil, errors.New("something went wrong " + r.Name)
	}
	greetWrapper := func(ctx context.Context, w io.Writer, r io.Reader) error {
		var reqs []*greetRequest
		err := json.NewDecoder(r).Decode(&reqs)
		is.NoErr(err)
		resps := make([]greetResponse, len(reqs))
		for i := range reqs {
			resp, err := greet(ctx, reqs[i])
			if err != nil {
				resps[i].Error = err.Error()
				continue
			}
			resps[i] = *resp
		}
		out := json.NewEncoder(w)
		out.SetEscapeHTML(false)
		if err = out.Encode(resps); err != nil {
			return err
		}
		return nil
	}
	srv := remotohttpjson.NewServer()
	srv.Register("/remoto/greeter.Greet", greetWrapper)
	in := `[{"name": "Mat"},
{"name": "David"},
{"name": "Aaron"}]`
	r, err := http.NewRequest(http.MethodPost, "/remoto/greeter.Greet", strings.NewReader(in))
	is.NoErr(err) // http.NewRequest
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusOK)
	var responses []*greetResponse
	err = json.NewDecoder(w.Body).Decode(&responses)
	is.NoErr(err)
	is.Equal(len(responses), 3)
	is.Equal(responses[0].Error, "something went wrong Mat")
	is.Equal(responses[1].Error, "something went wrong David")
	is.Equal(responses[2].Error, "something went wrong Aaron")
}
