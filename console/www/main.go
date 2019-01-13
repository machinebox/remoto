package www

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gobuffalo/plush"
	"github.com/machinebox/remoto/generator"
	"github.com/oxtoacart/bpool"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	s := &server{
		buffers: bpool.NewBufferPool(32),
	}
	http.HandleFunc("/definition", s.handleDefinitionSave())
	http.HandleFunc("/definition/", s.handleDefinitionView())
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
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
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
			tpl, err = plush.NewTemplate(b)
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

// handleDefinitionView loads the definition with the ID from the path,
// and renders a page with generators.
func (s *server) handleDefinitionView() http.HandlerFunc {
	var init sync.Once
	var err error
	var tpl *plush.Template
	type template struct {
		Name         string   `json:"name"`
		Experimental bool     `json:"x"`
		Dirs         []string `json:"dirs"`
		Label        string   `json:"label"`
	}
	var templates struct {
		Templates []template `json:"templates"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		init.Do(func() {
			var b string
			b, err = s.readFiles(
				"templates/definition.plush.html",
				"templates/_layout.plush.html",
			)
			if err != nil {
				return
			}
			tpl, err = plush.NewTemplate(b)
			if err != nil {
				return
			}
			client := urlfetch.Client(ctx)
			var resp *http.Response
			var apihost string
			apihost, err := appengine.ModuleHostname(ctx, "api", "", "")
			if err != nil {
				return
			}
			resp, err = client.Get("http://" + apihost + "/api/templates")
			if err != nil {
				return
			}
			defer resp.Body.Close()
			err = json.NewDecoder(resp.Body).Decode(&templates)
			if err != nil {
				return
			}
			log.Debugf(ctx, "%v", templates)
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sourceHash := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
		log.Debugf(ctx, "sourceHash: %v", sourceHash)
		entityKey := datastore.NewKey(ctx, kindDefinition, sourceHash, 0, nil)
		var entity entityDefinition
		err := datastore.Get(ctx, entityKey, &entity)
		if err == datastore.ErrNoSuchEntity {
			http.NotFound(w, r)
			return
		}
		def, err := generator.Parse(bytes.NewReader(entity.Source))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		plushCtx := plush.NewContextWithContext(ctx)
		generator.AddTemplateHelpers(plushCtx)
		plushCtx.Set("source", string(entity.Source))
		plushCtx.Set("def", def)
		plushCtx.Set("templates", templates.Templates)
		out, err := tpl.Exec(plushCtx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := io.WriteString(w, out); err != nil {
			log.Errorf(ctx, "%s", err)
		}
	}
}

// handleDefinitionSave saves an incoming definition, and redirects to the
// view page for that definition.
func (s *server) handleDefinitionSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		source := r.FormValue("definition")
		if source == "" {
			http.Error(w, "definition missing", http.StatusInternalServerError)
			return
		}
		_, err := generator.Parse(strings.NewReader(source))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		source = strings.TrimSpace(source)
		entity := entityDefinition{
			Source: []byte(source),
		}
		sourceHash := fmt.Sprintf("%x", md5.Sum([]byte(source)))
		entityKey := datastore.NewKey(ctx, kindDefinition, sourceHash, 0, nil)
		if _, err := datastore.Put(ctx, entityKey, &entity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/definition/"+sourceHash, http.StatusFound)
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

type entityDefinition struct {
	Key    *datastore.Key `datastore:"-"`
	Source []byte         `datastore:",noindex"`
}

var kindDefinition = "Definition"
