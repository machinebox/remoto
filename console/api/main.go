package api

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/machinebox/remoto/generator"
	"github.com/machinebox/remoto/generator/definition"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/api/templates", handleTemplates())
	http.HandleFunc("/api/all.zip", handleRenderAllTemplates())
	http.HandleFunc("/api/templates/", handleRenderTemplate())
	http.HandleFunc("/api/define", handleDefinitionDefine())
}

// handleTemplates gets a list of available templates.
//	GET /api/templates
func handleTemplates() http.HandlerFunc {
	var init sync.Once
	var err error
	type template struct {
		Name         string   `json:"name"`
		Experimental bool     `json:"x"`
		Dirs         []string `json:"dirs"`
		Label        string   `json:"label"`
	}
	var response struct {
		Templates []template `json:"templates"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			err = filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil // skip directories
				}
				path, err = filepath.Rel("templates", path)
				if err != nil {
					return err
				}
				if !strings.Contains(path, "/") {
					return nil // skip root files
				}
				if filepath.Ext(path) == ".plush" {
					path = path[0 : len(path)-len(".plush")]
				}
				response.Templates = append(response.Templates, template{
					Name:         path,
					Label:        filepath.Base(path),
					Dirs:         strings.Split(filepath.Dir(path), string(filepath.Separator)),
					Experimental: strings.Contains(path, "x/"),
				})
				return nil
			})
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleRenderTemplate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		templatePath := r.URL.Path[len("/api/"):] + ".plush"
		tplBytes, err := ioutil.ReadFile(templatePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		def, err := generator.Parse(strings.NewReader(r.FormValue("definition")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := filepath.Base(templatePath)
		filename = def.PackageName + "." + filename[0:len(filename)-len(filepath.Ext(templatePath))]
		if r.URL.Query().Get("dl") == "1" {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename=%q`, filename))
		}
		if err := generator.Render(w, templatePath, string(tplBytes), def); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleRenderAllTemplates() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		def, err := generator.Parse(strings.NewReader(r.FormValue("definition")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/zip, application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename=%s.remoto.zip`, def.PackageName))
		ww := zip.NewWriter(w)
		defer ww.Close()
		err = filepath.Walk("templates", func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() { // skip dirs
				return nil
			}
			if filepath.Ext(p) != ".plush" {
				return nil // skip non .plush files
			}
			templatePath := p
			p = p[0 : len(p)-len(".plush")] // trim .plush off
			zipPath, err := filepath.Rel("templates", p)
			if err != nil {
				return nil
			}
			zipPath = path.Join(def.PackageName, zipPath)
			f, err := ww.Create(zipPath)
			if err != nil {
				return err
			}
			tplBytes, err := ioutil.ReadFile(templatePath)
			if err != nil {
				log.Warningf(ctx, "skipping %q %v", templatePath, err)
				return nil
			}
			err = generator.Render(f, templatePath, string(tplBytes), def)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Errorf(ctx, "%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// handleDefinitionDefine checks the definition file, returning a JSON object
// with ok and error fields.
func handleDefinitionDefine() http.HandlerFunc {
	type response struct {
		OK         bool                   `json:"ok"`
		Error      string                 `json:"error,omitempty"`
		Definition *definition.Definition `json:"definition,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		source := r.FormValue("definition")
		if source == "" {
			http.Error(w, "definition missing", http.StatusInternalServerError)
			return
		}
		var response response
		response.OK = true
		def, err := generator.Parse(strings.NewReader(source))
		if err != nil {
			response.OK = false
			response.Error = err.Error()
		} else {
			response.Definition = &def
			if err := response.Definition.Valid(); err != nil {
				response.OK = false
				response.Error = err.Error()
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
