package parser

import (
	"bufio"
	"io"
	"regexp"
	"strings"
	"time"
)

var icons []string = []string{
	icon_created_date,
	icon_scheduled_date,
	icon_start_date,
	icon_due_date,
	icon_done_date,
	icon_cancelled_date,
	icon_recurrence,
}

var statuses map[string]string = map[string]string{
	" ": "open",
	"X": "closed",
	"O": "open",
	"o": "open",
	"<": "scheduling",
	">": "forwarded",
	"?": "question",
}

type Parser interface {
	Parse(io.Reader, TaskContext, chan Task)
}

type ParserImpl struct {
}

func NewParser() Parser {
	return &ParserImpl{}
}

func extractDate(icon_date string, line *string) *time.Time {
	pattern := `(` + icon_date + `\s*(\d{4}-\d{2}-\d{2})\s*)`
	dateRegex := regexp.MustCompile(pattern)
	matches := dateRegex.FindStringSubmatch(*line)
	if len(matches) == 3 {
		date, _ := time.Parse("2006-01-02", matches[2])
		*line = dateRegex.ReplaceAllString(*line, "")
		return &date
	}
	return nil
}

func extractPriority(line *string) string {
	for icon, priority := range icons_priority {
		if strings.Contains(*line, icon) {
			*line = strings.ReplaceAll(*line, icon, "")
			return priority
		}
	}
	return "normal"
}

func (p *ParserImpl) Parse(r io.Reader, context TaskContext, out chan Task) {
	scanner := bufio.NewScanner(r)
	pattern := `^\s*-\s*\[(.)\]\s*(.*)$`
	taskRegex := regexp.MustCompile(pattern)
	//println(pattern)
	lnum := 1
	for scanner.Scan() {
		line := scanner.Text()
		matches := taskRegex.FindStringSubmatch(line)

		if len(matches) == 3 {
			status := matches[1]
			taskLine := matches[2]
			cancelledDate := extractDate(icon_cancelled_date, &taskLine)
			createdDate := extractDate(icon_created_date, &taskLine)
			doneDate := extractDate(icon_done_date, &taskLine)
			dueDate := extractDate(icon_due_date, &taskLine)
			scheduledDate := extractDate(icon_scheduled_date, &taskLine)
			startDate := extractDate(icon_start_date, &taskLine)
			priority := extractPriority(&taskLine)
			thisLine := lnum
			context.Lnum = &thisLine

			task := Task{Description: taskLine,
				Status:        status,
				Priority:      priority,
				CancelledDate: cancelledDate,
				CreateDate:    createdDate,
				DoneDate:      doneDate,
				DueDate:       dueDate,
				ScheduledDate: scheduledDate,
				StartDate:     startDate,
				Context:       context,
			}

			out <- task
		}
		lnum += 1
	}
}
