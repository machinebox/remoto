package remotohttp_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestDecodeJSON(t *testing.T) {
	is := is.New(t)

	type r struct {
		Name string
	}

	j := `[
		{"name": "Mat"},
		{"name": "David"},
		{"name": "Aaron"}
	]`
	req, err := http.NewRequest(http.MethodPost, "/service/method", strings.NewReader(j))
	is.NoErr(err)

	var requestObjects []r
	err = remotohttp.Decode(req, &requestObjects)
	is.NoErr(err)

	is.Equal(len(requestObjects), 3)
	is.Equal(requestObjects[0].Name, "Mat")
	is.Equal(requestObjects[1].Name, "David")
	is.Equal(requestObjects[2].Name, "Aaron")

}
