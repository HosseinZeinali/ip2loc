package http

import (
	"bytes"
	"html/template"
	"io"
)

const layout = "templates/"

type MainView struct {
}

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(layout + "*.html")),
	}

	return renderer
}

func (renderer *TemplateRenderer) Exec(name string, view MainView) (io.WriterTo, error) {
	buf := bytes.NewBuffer([]byte{})
	err := renderer.templates.ExecuteTemplate(buf, name, view)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
