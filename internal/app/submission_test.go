package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSubmissionList(t *testing.T) {
	testData := [][]string{
		{"Doug Doenges", "84.2", "45"},
		{"Hi Flexion", "84.2", "40"},
		{"Can I", "84.2", "45"},
		{"Have a job?", "84.2", "45"},
	}

	submissions, err := NewSubmissionList(testData)
	assert.NoError(t, err)
	assert.Len(t, submissions, len(testData))
	for _, sb := range submissions {
		assert.NotEqual(t, "", sb.StudentName)
		assert.Len(t, sb.Responses, 2)
	}
}

func TestGradedSubmissionList(t *testing.T) {
	testData := [][]string{
		{"Test Name", "84.2", "45", "123"},
		{"Another Name", "84.2", "40", "123"},
		{"Iwanttowork AtFlexion", "84.2", "45", "123"},
	}
	answerKey := make([]*float64, 0, 2)
	a := 84.2
	b := 43.
	answerKey = append(answerKey, &a, &b, nil)

	submissions, err := NewSubmissionList(testData)
	assert.NoError(t, err)

	for _, s := range submissions {
		s.Grade(answerKey)
		assert.Len(t, s.Decisions, 3)
		assert.Equal(t, s.Decisions[0], Correct)
		assert.Equal(t, s.Decisions[1], Incorrect)
		assert.Equal(t, s.Decisions[2], Invalid)
	}
}

func TestSubmissionToGrid(t *testing.T) {
	testData := [][]string{
		{"Doug Doenges", "84.2", "45"},
		{"Doug Doenges", "", "45"},
	}
	submissions, _ := NewSubmissionList(testData)

	gridS := submissions[0].ToGrid(0)
	assert.Len(t, gridS, 2)
	assert.Equal(t, gridS[0], "84.2")

	gridS = submissions[1].ToGrid(0)
	assert.Len(t, gridS, 2)
	assert.Equal(t, gridS[0], "")
}
