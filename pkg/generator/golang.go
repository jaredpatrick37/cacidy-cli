package generator

import (
	_ "embed"
	"path/filepath"
)

const (
	DefaultGoVersion      = "1.22.0"
	golangImageRepository = "cacidy/go-runner"
)

var (
	//go:embed templates/golang/common-go-mod.tmpl
	commonGoMod string
	//go:embed templates/golang/common-main.tmpl
	commonMain string
	//go:embed templates/golang/hello-pipeline-go-mod.tmpl
	helloPipelineGoMod string
	//go:embed templates/golang/hello-pipeline-main.tmpl
	helloPipelineMain string
	//go:embed templates/golang/Dockerfile.tmpl
	dockerfile string
)

func NewGolangProject(path string, args GenerateArgs) error {
	r := newRenderer(path, args)
	r.addTemplate(filepath.Join("common", "go.mod"), commonGoMod, nil)
	r.addTemplate(filepath.Join("common", "common.go"), commonMain, nil)
	r.addTemplate(filepath.Join("pipelines/hello", "go.mod"), helloPipelineGoMod, nil)
	r.addTemplate(filepath.Join("pipelines/hello", "main.go"), helloPipelineMain, nil)
	r.addTemplate("Dockerfile", dockerfile, map[string]interface{}{"Image": golangImageRepository})
	return r.render()
}
