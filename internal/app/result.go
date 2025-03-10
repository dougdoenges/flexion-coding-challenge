package app

type Results struct {
	input             Worksheet
	gradedSubmissions []Submission
}

func GetResults(ws Worksheet, submissions []Submission) Results {
	for idx := range submissions {
		submissions[idx].Grade(ws.Key())
	}
	return Results{
		input:             ws,
		gradedSubmissions: submissions,
	}
}

func (r *Results) ToGridDisplay() [][]string {
	gridDisplay := make([][]string, 0, len(r.input.Questions)+1)

	const spacer = ""
	const colsPerStudent = 2
	const headerCount = 5
	numCols := len(r.gradedSubmissions)*colsPerStudent + headerCount

	// make header row
	headerRow := make([]string, 0, numCols)
	headerRow = append(headerRow,
		[]string{"Input", "From Unit", "To Unit", "Correct Answer"}...)
	headerRow = append(headerRow, spacer)
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
		row = append(row, spacer)
		for _, submission := range r.gradedSubmissions {
			row = append(row, submission.ToGrid(idx)...)
		}
		gridDisplay = append(gridDisplay, row)
	}

	return gridDisplay
}
