package view

import (
	"fmt"
	"io"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer implements echo.Renderer.
type TemplateRenderer struct {
	tpls *template.Template
}

// New parses *.tpl files under dir and returns a renderer.
// one call in main
func New(dir string) (*TemplateRenderer, error) {
	pattern := filepath.Join(dir, "*.tpl")
	t, tplErr := template.New("").ParseGlob(pattern)
	if tplErr != nil {
		return nil, fmt.Errorf("parse templates: %w", tplErr)
	}
	return &TemplateRenderer{tpls: t}, nil
}

// Render makes TemplateRenderer satisfy echo.Renderer.
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.tpls.ExecuteTemplate(w, name, data)
}
