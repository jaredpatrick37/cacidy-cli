package generator

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mitchellh/mapstructure"
)

type rendererTemplate struct {
	dest string
	text string
	args map[string]interface{}
}

type renderer struct {
	path      string
	args      GenerateArgs
	templates []rendererTemplate
}

func (r *renderer) addTemplate(dest, text string, args map[string]interface{}) {
	r.templates = append(r.templates, rendererTemplate{dest, text, args})
}

func (r *renderer) render() error {
	for _, tmpl := range r.templates {
		var text bytes.Buffer
		if err := mapstructure.Decode(r.args, &tmpl.args); err != nil {
			return err
		}
		t := template.Must(template.New("tmpl").Parse(tmpl.text))
		if err := t.Execute(&text, tmpl.args); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(r.path, tmpl.dest), text.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func newRenderer(path string, args GenerateArgs) *renderer {
	return &renderer{path: path, args: args}
}
