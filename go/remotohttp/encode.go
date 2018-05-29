package remotohttp

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Encode writes the response.
func Encode(w http.ResponseWriter, _ *http.Request, status int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "encode json")
	}
	w.Header().Set("Content-Type", "application/json; chatset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}
