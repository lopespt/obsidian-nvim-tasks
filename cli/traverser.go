package main

import (
	"os"
	"path/filepath"

	"github.com/lopespt/obsidian-tasks.git/parser"
)

func parseVault(path string, p parser.Parser, out chan parser.Task) {
	// Open the vault
	filepath.Walk(path, func(path string, info os.FileInfo, readDirErr error) error {
		go func(path string) {
			f, _ := os.Open(path)
			defer f.Close()
			p.Parse(f, out)
		}(path)
		return nil
	})

	close(out)
}
