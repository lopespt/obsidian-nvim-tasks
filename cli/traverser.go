package main

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/lopespt/obsidian-tasks.git/parser"
)

func parseVault(path string, p parser.Parser, out chan parser.Task) {
	// Open the vault
	wg := sync.WaitGroup{}

	filepath.Walk(path, func(path string, info os.FileInfo, readDirErr error) error {
		wg.Add(1)
		if readDirErr != nil {
			println(readDirErr.Error())
		}
		go func(path string) {
			defer wg.Done()
			f, err := os.Open(path)
			if err != nil {
				println(err.Error())
			}
			defer f.Close()
			context := parser.TaskContext{
				Filename: path,
			}
			p.Parse(f, context, out)
		}(path)
		return nil
	})

	wg.Wait()
	close(out)
}
