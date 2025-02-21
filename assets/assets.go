package assets

import (
	"embed"
	"fmt"
	"github.com/pborges/errs"
	"io/fs"
)

// All provides the file systems for each folder within the "assets" directory.  For easier access, please use assets.Get.
var All = []embed.FS{Shaders, Audio}

// Audio provides an embed.FS of the "audio" directory.  For easier access, please use assets.Get.
//
//go:embed audio/*
var Audio embed.FS

// AudioFiles returns all files in the "audio" directory.
func (g get) AudioFiles() []string {
	return g.Contents("audio")
}

// AudioFile returns back a specific audio file rooted in the "audio" directory.
func (g get) AudioFile(path string) []byte {
	data, err := Audio.ReadFile("audio/" + path)
	if err != nil {
		err = errs.Wrap(err, fmt.Errorf("failed to read audio file: %s", path))
		panic(errs.Detailed(err))
	}
	return data
}

// Shaders provides an embed.FS of the "shaders" directory.  For easier access, please use assets.Get.
//
//go:embed shaders/*
var Shaders embed.FS

// Shaders returns all files in the "shaders" directory.
func (g get) Shaders() []string {
	return g.Contents("shaders")
}

// Shader returns back a specific shader file rooted in the "shaders" directory.
func (g get) Shader(path string) []byte {
	data, err := Shaders.ReadFile("shaders/" + path)
	if err != nil {
		err = errs.Wrap(err, fmt.Errorf("failed to read shader: %s", path))
		panic(errs.Detailed(err))
	}
	return data
}

type get int

// Get provides a fluent API into the local Glitter assets.
var Get get

// Contents returns back the files contained in a particular subfolder by passing root along to fs.WalkDir
// and filtering out only the files found.
func (g get) Contents(root string) []string {
	var output []string
	for _, fileSystem := range All {
		fs.WalkDir(fileSystem, root, func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() {
				output = append(output, path)
			}
			return nil
		})
	}
	return output
}
