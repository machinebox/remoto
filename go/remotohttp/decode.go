package remotohttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Decode extracts the incoming data from the http.Request.
func Decode(r *http.Request, v interface{}) error {
	contentType := strings.ToLower(r.Header.Get("Content-Type"))
	switch {
	case strings.Contains(contentType, "application/json"):
		return decodeJSON(r, v)
	case strings.Contains(contentType, "application/x-www-form-urlencoded"),
		strings.Contains(contentType, "multipart/form-data"):
		return decodeFormdata(r, v)
	}
	return errors.New("unsupported Content-Type (use application/json, application/x-www-form-urlencoded or multipart/form-data)")
}

func decodeJSON(r *http.Request, v interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return errors.Wrap(err, "decode json")
	}
	return nil
}

func decodeFormdata(r *http.Request, v interface{}) error {
	j := r.FormValue("json")
	if j == "" {
		return errors.New("missing field: json")
	}
	if err := json.Unmarshal([]byte(j), v); err != nil {
		return errors.Wrap(err, "decode json")
	}
	return nil
}
