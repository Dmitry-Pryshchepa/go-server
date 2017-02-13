package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateHandler struct {
	once     sync.Once
	Filename string
	Templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.Templ = template.Must(template.ParseFiles(filepath.Join("templates/html", t.Filename)))
	})
	t.Templ.Execute(w, nil)
}
