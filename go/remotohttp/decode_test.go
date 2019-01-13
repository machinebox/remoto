package remotohttp_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/machinebox/remoto/go/remotohttp"
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
	req.Header.Set("Content-Type", "application/json")
	var requestObjects []r
	err = remotohttp.Decode(req, &requestObjects)
	is.NoErr(err)
	is.Equal(len(requestObjects), 3)
	is.Equal(requestObjects[0].Name, "Mat")
	is.Equal(requestObjects[1].Name, "David")
	is.Equal(requestObjects[2].Name, "Aaron")
}

func TestDecodeFormData(t *testing.T) {
	is := is.New(t)
	type r struct {
		Name string
	}
	j := `[
		{"name": "Mat"},
		{"name": "David"},
		{"name": "Aaron"}
	]`
	data := url.Values{}
	data.Set("json", j)
	req, err := http.NewRequest(http.MethodPost, "/service/method", strings.NewReader(data.Encode()))
	is.NoErr(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var requestObjects []r
	err = remotohttp.Decode(req, &requestObjects)
	is.NoErr(err)
	is.Equal(len(requestObjects), 3)
	is.Equal(requestObjects[0].Name, "Mat")
	is.Equal(requestObjects[1].Name, "David")
	is.Equal(requestObjects[2].Name, "Aaron")
}

func TestDecodeMultipartForm(t *testing.T) {
	is := is.New(t)
	type r struct {
		Name string
	}
	j := `[
		{"name": "Mat"},
		{"name": "David"},
		{"name": "Aaron"}
	]`
	data := url.Values{}
	data.Set("json", j)
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	err := w.WriteField("json", j)
	is.NoErr(err)
	err = w.Close()
	is.NoErr(err)
	req, err := http.NewRequest(http.MethodPost, "/service/method", &buf)
	is.NoErr(err)
	req.Header.Set("Content-Type", w.FormDataContentType())
	var requestObjects []r
	err = remotohttp.Decode(req, &requestObjects)
	is.NoErr(err)
	is.Equal(len(requestObjects), 3)
	is.Equal(requestObjects[0].Name, "Mat")
	is.Equal(requestObjects[1].Name, "David")
	is.Equal(requestObjects[2].Name, "Aaron")
}

// func TestDecodeFile(t *testing.T) {
//	is := is.New(t)
//	type r struct {
//		Name  string
//		Photo remototypes.File
//	}
//	j := string(`[
//		{"name": "Mat", "photo": "` + remototypes.NewFile("files[0]") + `"},
//		{"name": "David", "photo": "` + remototypes.NewFile("files[1]") + `"},
//		{"name": "Aaron", "photo": "` + remototypes.NewFile("files[2]") + `"}
//	]`)
//	var buf bytes.Buffer
//	w := multipart.NewWriter(&buf)
//	w.WriteField("json", j)
//	f1, err := w.CreateFormFile("files[0]", "mat.jpg")
//	is.NoErr(err)
//	f1.Write([]byte("mat-photo-binary-data"))
//	f2, err := w.CreateFormFile("files[1]", "david.jpg")
//	is.NoErr(err)
//	f2.Write([]byte("david-photo-binary-data"))
//	f3, err := w.CreateFormFile("files[2]", "aaron.jpg")
//	is.NoErr(err)
//	f3.Write([]byte("aaron-photo-binary-data"))
//	is.NoErr(w.Close())
//	req, err := http.NewRequest(http.MethodPost, "/service/method", &buf)
//	is.NoErr(err)
//	req.Header.Set("Content-Type", w.FormDataContentType())
//	var requestObjects []r
//	err = remotohttp.Decode(req, &requestObjects)
//	is.NoErr(err)
//	is.Equal(len(requestObjects), 3)
//	is.Equal(requestObjects[0].Name, "Mat")
//	is.Equal(requestObjects[1].Name, "David")
//	is.Equal(requestObjects[2].Name, "Aaron")
//	// ctx := context.Background()
//	// f, err := requestObjects[0].Photo.Open(ctx)
//	// is.NoErr(err)
//	// defer f.Close()
//	// matPhotoData, err := ioutil.ReadAll(f)
//	// is.NoErr(err)
//	// is.Equal(string(matPhotoData), `mat-photo-binary-data`)
// }
