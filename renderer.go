package main

import (
	"html/template"
	"io"

	"github.com/Sirupsen/logrus"
)

// Renderer provides method `Render` to render the file with given data.
type Renderer struct {
	template *template.Template
}

func NewRenderer(filename string) *Renderer {
	return &Renderer{
		template: template.New(filename).ParseFiles(filename),
	}
}

func (r *Renderer) Render(wr io.Writer, data interface{}) error {
	r.template.Execute(wr, data)
	return nil
}
