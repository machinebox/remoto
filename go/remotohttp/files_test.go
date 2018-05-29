package remotohttp_test

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/machinebox/remoto/remototypes"
	"github.com/matryer/is"
)

func TestOpenFile(t *testing.T) {
	is := is.New(t)
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	f1, err := w.CreateFormFile("files[0]", "mat.jpg")
	is.NoErr(err)
	f1.Write([]byte("mat-photo-binary-data"))
	f2, err := w.CreateFormFile("files[1]", "david.jpg")
	is.NoErr(err)
	f2.Write([]byte("david-photo-binary-data"))
	f3, err := w.CreateFormFile("files[2]", "aaron.jpg")
	is.NoErr(err)
	f3.Write([]byte("aaron-photo-binary-data"))
	is.NoErr(w.Close())
	req, err := http.NewRequest(http.MethodPost, "/service/method", &buf)
	is.NoErr(err)
	req.Header.Set("Content-Type", w.FormDataContentType())
	f, err := remotohttp.OpenFile(context.Background(), remototypes.NewFile("files[0]"), req, http.DefaultClient)
	is.NoErr(err)
	defer f.Close()
	matPhotoData, err := ioutil.ReadAll(f)
	is.NoErr(err)
	is.Equal(string(matPhotoData), `mat-photo-binary-data`)
}

func TestOpenURL(t *testing.T) {
	is := is.New(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "mat-photo-binary-data")
	}))
	defer srv.Close()
	f, err := remotohttp.OpenFile(context.Background(), remototypes.NewFileURL(srv.URL), nil, http.DefaultClient)
	is.NoErr(err)
	defer f.Close()
	matPhotoData, err := ioutil.ReadAll(f)
	is.NoErr(err)
	is.Equal(string(matPhotoData), `mat-photo-binary-data`)
}
