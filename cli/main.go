package main

import (
	"encoding/json"
	"os"
	"sync"

	parser "github.com/lopespt/obsidian-tasks.git/parser"
)

func taskWorker(tasksChannel chan parser.Task, mut *sync.Mutex) {
	jenc := json.NewEncoder(os.Stdout)
	go func() {
		for task := range tasksChannel {
			mut.Lock()
			jenc.Encode(task)
			mut.Unlock()
		}
	}()
}

func main() {
	p := parser.NewParser()
	blocker := &sync.Mutex{}
	tasksChannel := make(chan parser.Task)
	for range 10 {
		taskWorker(tasksChannel, blocker)
	}

	parseVault("/Users/gwachs/obsidian/gwachs", p, tasksChannel)
}
