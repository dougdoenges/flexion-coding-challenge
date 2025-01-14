package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResultsToGridDisplay(t *testing.T) {
	testWs := [][]string{
		{"1.1", "liters", "cups"},
	}
	testSubmissions := [][]string{
		{"Doug Doenges", "84.2"},
	}
	ws, _ := NewWorksheet(testWs)
	submissions, _ := NewSubmissionList(testSubmissions)

	res := GetResults(ws, submissions)
	assert.NotEqual(t, res, Results{})

	gridDisplay := res.ToGridDisplay()
	testDisplay := [][]string{
		{"Input", "From Unit", "To Unit", "Correct Answer", "", testSubmissions[0][0], ""},
		{testWs[0][0], testWs[0][1], testWs[0][2], "4.6", "", testSubmissions[0][1], "Incorrect"},
	}
	assert.Equal(t, testDisplay, gridDisplay)
}
