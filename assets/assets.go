// Package assets provides a standardized way of accessing external data.
package assets

import (
	"embed"
	"fmt"
	"github.com/pborges/errs"
	"io/fs"
)

// All provides the file systems for each folder within the "assets" directory.
//
// For easier access, please use assets.Get.
//
//go:embed content/*
var All embed.FS

// AudioFiles returns all files in the "audio" directory.
func (g get) AudioFiles() []string {
	return g.Files("content/audio")
}

// AudioFile returns back a specific audio file rooted in the "audio" directory.
func (g get) AudioFile(path string) []byte {
	data, err := All.ReadFile("content/audio/" + path)
	if err != nil {
		err = errs.Wrap(err, fmt.Errorf("failed to read audio file: %s", path))
		panic(errs.Detailed(err))
	}
	return data
}

// Shaders returns all files in the "shaders" directory.
func (g get) Shaders() []string {
	return g.Files("content/shaders")
}

// Shader returns back a specific shader file rooted in the "shaders" directory.
func (g get) Shader(path string) string {
	data, err := All.ReadFile("content/shaders/" + path)
	if err != nil {
		err = errs.Wrap(err, fmt.Errorf("failed to read shader: %s", path))
		panic(errs.Detailed(err))
	}
	return string(data)
}

type get int

// Get provides a fluent API into the local glitter assets.
var Get get

// Files returns back the files contained in a particular subfolder by passing root along to fs.WalkDir
// and filtering out only the files found.  If no root is provided, it defaults to "."
func (g get) Files(root ...string) []string {
	if len(root) == 0 {
		root = append(root, ".")
	}
	var output []string
	fs.WalkDir(All, root[0], func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			output = append(output, path)
		}
		return nil
	})
	return output
}
