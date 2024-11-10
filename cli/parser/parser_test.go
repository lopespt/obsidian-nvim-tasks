package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractDate(t *testing.T) {
	s := CreateTaskLine()
	date := extractDate(icon_scheduled_date, &s)
	year, month, day := date.Date()
	if year != 2024 || month != 11 || day != 14 {
		t.Errorf("Expected 2024-11-14, got %s", date)
	}
	assert.Equal(t, "- [] Marcar entrevista com Ana Claudia sobre Programação Dinâmica e Avisar (julio.brazil@uber.com) dia ✅ 2024-11-04", s)

	date = extractDate(icon_done_date, &s)
	year, month, day = date.Date()
	if year != 2024 || month != 11 || day != 4 {
		t.Errorf("Expected 2024-11-04, got %s", date)
	}
	assert.Equal(t, "- [] Marcar entrevista com Ana Claudia sobre Programação Dinâmica e Avisar (julio.brazil@uber.com) dia ", s)
}

func CreateTaskLine() string {
	return "- [] Marcar entrevista com Ana Claudia sobre Programação Dinâmica e Avisar (julio.brazil@uber.com) dia ⏳ 2024-11-14 ✅ 2024-11-04"
}
