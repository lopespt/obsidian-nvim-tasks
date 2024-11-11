package main

import (
	"flag"
	"strings"
)

type Args struct {
	VaultPath         *string
	FilterNotStatuses []string
	FilterStatuses    []string
	SortBy            string
}

func NewArgs() Args {
	args := Args{}

	vaultPath := flag.String("v", "", "Path to the obsidian vault")

	filterNotStatuses := flag.String("ns", "", "Filter out tasks with these statuses. Comma separated list")
	filtertStatuses := flag.String("s", "", "Filter tasks with these statuses. Comma separated list")
	sortByDue := flag.Bool("SortDueDate", false, "Sort by due date")
	sortByAny := flag.Bool("SortAnyDate", false, "Sort by due or scheduled date, what comes first")
	sortByScheduled := flag.Bool("SortScheduledDate", false, "Sort by scheduled date")

	flag.Parse()
	args.VaultPath = vaultPath
	args.FilterNotStatuses = strings.Split(*filterNotStatuses, ",")
	if args.FilterNotStatuses[0] == "" {
		args.FilterNotStatuses = []string{}
	}
	args.FilterStatuses = strings.Split(*filtertStatuses, ",")
	if args.FilterStatuses[0] == "" {
		args.FilterStatuses = []string{}
	}
	if *sortByDue {
		args.SortBy = "dueDate"
	}
	if *sortByAny {
		args.SortBy = "anyDate"
	}
	if *sortByScheduled {
		args.SortBy = "scheduledDate"
	}

	return args
}
