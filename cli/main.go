package main

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
	"sync"

	parser "github.com/lopespt/obsidian-tasks.git/parser"
)

func taskWorker(tasksChannel chan parser.Task, mut *sync.Mutex, allTasks *[]parser.Task, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	go func() {
		for task := range tasksChannel {
			mut.Lock()
			*allTasks = append(*allTasks, task)
			mut.Unlock()
		}
	}()
}

func getAllTasks(vaultPath string) []parser.Task {
	allTasks := []parser.Task{}

	p := parser.NewParser()
	blocker := &sync.Mutex{}
	tasksChannel := make(chan parser.Task)
	wg := sync.WaitGroup{}

	for range 10 {
		taskWorker(tasksChannel, blocker, &allTasks, &wg)
	}

	parseVault(vaultPath, p, tasksChannel)
	wg.Wait()
	return allTasks
}

func filter(tasks *[]parser.Task, args Args) {
	//filter not statuses
	if len(args.FilterNotStatuses) > 0 {
		*tasks = slices.DeleteFunc(*tasks, func(task parser.Task) bool {
			return slices.Contains(args.FilterNotStatuses, task.Status)
		})
	}

	//filter statuses
	if len(args.FilterStatuses) > 0 {
		*tasks = slices.DeleteFunc(*tasks, func(task parser.Task) bool {
			return !slices.Contains(args.FilterStatuses, task.Status)
		})
	}
}

func sort(tasks *[]parser.Task, args Args) {
	slices.SortFunc(*tasks, func(t1 parser.Task, t2 parser.Task) int {
		return strings.Compare(t1.Description, t2.Description)
	})
	slices.SortStableFunc(*tasks, func(t1 parser.Task, t2 parser.Task) int {
		return parser.PriorityToInt[t1.Priority] - parser.PriorityToInt[t2.Priority]
	})
	switch args.SortBy {
	case "anyDate":
		slices.SortStableFunc(*tasks, func(t1 parser.Task, t2 parser.Task) int {
			d1 := t1.ScheduledDate
			if d1 == nil || (t1.DueDate != nil && d1.After(*t1.DueDate)) {
				d1 = t1.DueDate
			}
			d2 := t2.ScheduledDate
			if d2 == nil || (t2.DueDate != nil && d2.After(*t2.DueDate)) {
				d2 = t2.DueDate
			}

			if d1 != nil && d2 != nil {
				if d1.Before(*d2) {
					return -1
				}
				if d1.After(*d2) {
					return 1
				}
			}
			if d1 == nil {
				return 1
			}
			return -1
		})

	case "scheduledDate":
		slices.SortStableFunc(*tasks, func(t1 parser.Task, t2 parser.Task) int {
			if t1.ScheduledDate != nil && t2.ScheduledDate != nil {
				if t1.ScheduledDate.Before(*t2.ScheduledDate) {
					return -1
				}
				if t1.ScheduledDate.After(*t2.ScheduledDate) {
					return 1
				}
			}
			if t1.ScheduledDate == nil {
				return 1
			}
			return -1
		})
	case "dueDate":
		slices.SortStableFunc(*tasks, func(t1 parser.Task, t2 parser.Task) int {
			if t1.DueDate != nil && t2.DueDate != nil {
				if t1.DueDate.Before(*t2.DueDate) {
					return -1
				}
				if t1.DueDate.After(*t2.DueDate) {
					return 1
				}
			}
			if t1.DueDate == nil {
				return 1
			}
			return -1
		})
	}
}

func main() {
	args := NewArgs()

	tasks := getAllTasks(*args.VaultPath)
	filter(&tasks, args)
	sort(&tasks, args)

	enc := json.NewEncoder(os.Stdout)
	for _, task := range tasks {
		enc.Encode(task)
	}
}
