package app

import (
	"fmt"
)

type Results struct {
	input             Worksheet
	gradedSubmissions []Submission
}

func GetResults(ws Worksheet, submissions []Submission) Results {
	for idx := range submissions {
		submissions[idx].Grade(ws)
	}
	return Results{
		input:             ws,
		gradedSubmissions: submissions,
	}
}

func (r *Results) ToGridDisplay() [][]string {
	// todo: create grid with values
	gridDisplay := make([][]string, 0, len(r.input.Questions)+1)

	colsPerStudent := 2
	headerCount := 5
	numCols := len(r.gradedSubmissions)*colsPerStudent + headerCount

	// make header row
	headerRow := make([]string, 0, numCols)
	headerRow = append(headerRow, []string{"Input", "From Unit", "To Unit", "Correct Answer", ""}...)
	for _, submission := range r.gradedSubmissions {
		headerRow = append(headerRow, submission.StudentName)
		for range colsPerStudent - 1 {
			headerRow = append(headerRow, "")
		}
	}
	gridDisplay = append(gridDisplay, headerRow)

	// populate questions and answers
	for idx := range r.input.Questions {
		row := make([]string, 0, numCols)
		row = append(row, r.input.Questions[idx].ToGrid()...)
		for _, submission := range r.gradedSubmissions {
			row = append(row, submission.ToGrid(idx)...)
		}
		gridDisplay = append(gridDisplay, row)
	}

	return gridDisplay
}

// TODO
func (r *Results) ToString() string {
	const ROW_DIVIDER = "-"
	const COL_DIVIDER = "|"

	fmt.Println(ROW_DIVIDER)
	fmt.Println(COL_DIVIDER)

	return ""
}

// Options for exporting to csv, json, etc

// pull file logic out of package code
// Finalize calculation logic
// reorganize errors and test
