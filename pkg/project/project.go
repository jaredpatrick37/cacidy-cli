package project

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mitchellh/mapstructure"
)

const (
	GoSDK     = "go"
	NodeSDK   = "node"
	PythonSDK = "python"
)

var SDKS = []string{GoSDK, NodeSDK, PythonSDK}

type NewProjectArgs struct {
	GoVersion string
}

func New(path, sdk string, args NewProjectArgs) error {
	switch sdk {
	case GoSDK:
		return NewGolangProject(path, args)
	}
	return nil
}

type rendererTemplate struct {
	dest string
	text string
	args map[string]interface{}
}

type renderer struct {
	path      string
	args      NewProjectArgs
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
		fmt.Printf("%+v\n", tmpl.args)
		if err := os.WriteFile(filepath.Join(r.path, tmpl.dest), text.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func newRenderer(path string, args NewProjectArgs) *renderer {
	return &renderer{path: path, args: args}
}

func initializeProject(path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return errors.New("the provided path already exists")
	}
	for _, dir := range []string{"bin", "common", "pipelines/hello"} {
		if err := os.MkdirAll(filepath.Join(path, dir), 0755); err != nil {
			return err
		}
	}
	return nil
}
