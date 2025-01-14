package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorksheet(t *testing.T) {
	testData := [][]string{
		{"1.10", "liters", "cups"},
		{"2.20", "tablespoons", "cubic inches"},
		{"3.30", "gallons", "liters"},
		{"4.40", "cups", "liters"},
		{"5.50", "cubic feet", "gallons"},
	}

	ws, err := NewWorksheet(testData)
	assert.NoError(t, err)
	assert.NotEqual(t, ws, Worksheet{})
	assert.Len(t, ws.Questions, len(testData))

	testData[0][0] = "NOT A NUMBER"
	ws, err = NewWorksheet(testData)
	assert.Error(t, err)
	assert.Equal(t, ws, Worksheet{})

	testData[0] = append(testData[0], "invalid question part.")
	ws, err = NewWorksheet(testData)
	assert.Error(t, err)
	assert.Equal(t, ws, Worksheet{})
}

func TestKey(t *testing.T) {
	testData := [][]string{
		{"1.10", "liters", "cups"},
		{"2.20", "tablespoons", "cubic inches"},
		{"3.30", "gallons", "liters"},
		{"4.40", "cups", "liters"},
		{"5.50", "cubic feet", "gallons"},
	}

	ws, _ := NewWorksheet(testData)
	assert.NotEqual(t, ws, Worksheet{})

	answerKey := ws.Key()
	assert.Len(t, answerKey, 5)
	for _, a := range answerKey {
		assert.NotNil(t, a)
	}

	testData = [][]string{
		{"1.1", "not a unit", "liters"},
		{"1.2", "also not a unit", "Kelvin"},
		{"1.3", "celsius", "kelvin"},
	}
	ws, _ = NewWorksheet(testData)
	assert.NotEqual(t, ws, Worksheet{})
	answerKey = ws.Key()
	assert.Len(t, answerKey, 3)
	assert.Nil(t, answerKey[0])
	assert.Nil(t, answerKey[1])
	assert.NotNil(t, answerKey[2])
}

func TestWorksheetToGrid(t *testing.T) {
	q := Question{
		Input:     123.123,
		InputUoM:  "kelvin",
		TargetUoM: "celsius",
	}
	gridQ := q.ToGrid()
	assert.Len(t, gridQ, 4)

	q = Question{
		Input:     123.123,
		InputUoM:  "not a unit",
		TargetUoM: "also not a unit",
	}
	gridQ = q.ToGrid()
	assert.Len(t, gridQ, 4)
}
