package generator

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	GoSDK     = "go"
	NodeSDK   = "node"
	PythonSDK = "python"
)

var SDKS = []string{GoSDK, NodeSDK, PythonSDK}

type GenerateArgs struct {
	GoVersion string
}

func New(path, sdk string, args GenerateArgs) error {
	if path == "" {
		return errors.New("project path is required")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return errors.New("the provided path already exists")
	}
	for _, dir := range []string{"bin", "common", "pipelines/hello"} {
		if err := os.MkdirAll(filepath.Join(path, dir), 0755); err != nil {
			return err
		}
	}
	switch sdk {
	case GoSDK:
		return NewGolangProject(path, args)
	}
	return nil
}
