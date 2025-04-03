package render

import (
	"html/template"
	"io/fs"
)

type Renderer struct {
	fs fs.FS
}

func NewRenderer(fs fs.FS) *Renderer {
	return &Renderer{
		fs: fs,
	}
}

func (r *Renderer) Render(names ...string) (*template.Template, error) {
	return template.ParseFS(r.fs, names...)
}
