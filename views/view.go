package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir -> layout directory
	LayoutDir = "views/layouts/*"
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

// Render -> rendering template
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// NewView -> compile template
func NewView(layout string, files ...string) *View {
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
