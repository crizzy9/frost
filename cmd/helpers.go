package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func resolveFilePathWithEnvVar(path string) (resolvedPath string) {
	switch t := strings.Split(path, "/"); path[0] {
	case '$':
		fmt.Println("Resolving environment variable: ", t[0])
		resolvedPath = filepath.Join(os.Getenv(strings.Trim(t[0], "$")), filepath.Join(t[1:]...))
	case '~':
		resolvedPath = filepath.Join(os.Getenv("HOME"), filepath.Join(t[1:]...))
	case '/':
		resolvedPath = path
	default:
		wd, _ := os.Getwd()
		resolvedPath = filepath.Join(wd, strings.Trim(path, "./"))
	}

	return
}
