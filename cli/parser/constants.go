package parser

const (
	icon_created_date   = "➕"
	icon_scheduled_date = "⏳"
	icon_start_date     = "🛫"
	icon_due_date       = "📅"
	icon_done_date      = "✅"
	icon_cancelled_date = "❌"

	icon_lowest_priority  = "⏬"
	icon_low_priority     = "🔽"
	icon_medium_priority  = "🔼"
	icon_high_priority    = "⏫"
	icon_highest_priority = "🔺"

	icon_recurrence = "🔁"
)

var icons_priority map[string]string = map[string]string{
	icon_lowest_priority:  "lowest",
	icon_low_priority:     "low",
	icon_medium_priority:  "medium",
	icon_high_priority:    "high",
	icon_highest_priority: "highest",
}

var PriorityToInt map[string]int = map[string]int{
	"lowest":  0,
	"low":     1,
	"normal":  2,
	"medium":  3,
	"high":    4,
	"highest": 5,
}
