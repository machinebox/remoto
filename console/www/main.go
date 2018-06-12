package www

import (
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gobuffalo/plush"
	"github.com/oxtoacart/bpool"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	s := &server{
		buffers: bpool.NewBufferPool(32),
	}
	http.HandleFunc("/", s.handleIndex())
}

type server struct {
	buffers *bpool.BufferPool
}

func (s *server) handleIndex() http.HandlerFunc {
	var init sync.Once
	var err error
	var tpl *plush.Template
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		init.Do(func() {
			var b string
			b, err = s.readFiles(
				"templates/index.plush.html",
				"templates/_layout.plush.html",
			)
			if err != nil {
				return
			}
			tpl, err = plush.NewTemplate(string(b))
			if err != nil {
				return
			}
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		out, err := tpl.Exec(plush.NewContext())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.WriteString(w, out); err != nil {
			log.Errorf(ctx, "%s", err)
		}
	}
}

func errorHandler(err error, status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, err.Error(), status)
	}
}

// readFiles loads one or more files into a string.
func (s *server) readFiles(filenames ...string) (string, error) {
	buf := s.buffers.Get()
	defer s.buffers.Put(buf)
	for _, file := range filenames {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		if _, err := buf.Write(b); err != nil {
			return "", err
		}
		if _, err := buf.WriteString("\n"); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
