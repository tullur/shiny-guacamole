package views

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir -> layout directory
	LayoutDir = "views/layouts/*"
	// TemplateDir -> template Dir
	TemplateDir = "views/"
	// TemplateExt -> template extension
	TemplateExt = ".gohtml"
)

// View -> template struct
type View struct {
	Template *template.Template
	Layout   string
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	return files
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExtension(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

// Render -> rendering template
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Conent-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		log.Fatalln(err)
	}
}

// NewView -> compile template
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExtension(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}
