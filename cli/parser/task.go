package parser

import (
	"time"
)

type Task struct {
	Description   string     `json:"description,omitempty"`
	Status        string     `json:"status,omitempty"`
	Priority      string     `json:"priority,omitempty"`
	CreateDate    *time.Time `json:"createDate,omitempty"`
	ScheduledDate *time.Time `json:"scheduledDate,omitempty"`
	StartDate     *time.Time `json:"startDate,omitempty"`
	DueDate       *time.Time `json:"dueDate,omitempty"`
	DoneDate      *time.Time `json:"doneDate,omitempty"`
	CancelledDate *time.Time `json:"cancelledDate,omitempty"`
}