package remotohttp

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Encode writes the response.
func Encode(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "encode json")
	}
	var out io.Writer = w
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gzw := gzip.NewWriter(w)
		out = gzw
		defer gzw.Close()
	}
	w.Header().Set("Content-Type", "application/json; chatset=utf-8")
	w.WriteHeader(status)
	if _, err := out.Write(b); err != nil {
		return err
	}
	return nil
}

// EncodeErr writes an error response with http.StatusInternalServerError.
func EncodeErr(w http.ResponseWriter, r *http.Request, err error) error {
	// returns [{"error":"message"}]
	e := []struct {
		Error string `json:"error"`
	}{
		{
			Error: err.Error(),
		},
	}
	return Encode(w, r, http.StatusInternalServerError, e)
}
