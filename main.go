package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// walk the directory tree
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			if strings.HasPrefix(path, "pkg/configs/") {
				return nil
			} else {
				println(path)
				os.Remove(path)
			}
		}
		return nil
	})

}
