package app

import (
	"fmt"
	"strconv"
	"strings"
)

type Worksheet struct {
	Questions []Question
}

type Question struct {
	Input     float64
	InputUoM  string
	TargetUoM string

	CorrectAnswer *float64 // nil for an invalid question
}

const QuestionLength = 3

func NewWorksheet(data [][]string) (Worksheet, error) {
	ws := Worksheet{}

	for _, row := range data {
		question, err := buildQuestion(row)
		if err != nil {
			return Worksheet{}, err
		}

		ws.Questions = append(ws.Questions, question)
	}

	return ws, nil
}

func buildQuestion(data []string) (Question, error) {
	q := Question{}

	if len(data) != QuestionLength {
		return Question{},
			fmt.Errorf("invalid question provided: %s", strings.Join(data, ","))
	}

	input, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return Question{},
			fmt.Errorf("invalid input number(s) given: %s", strings.Join(data, ","))
	}
	q.Input = input

	inputUom := data[1]
	q.InputUoM = strings.TrimSpace(strings.ToLower(inputUom))

	targetUoM := data[2]
	q.TargetUoM = strings.TrimSpace(strings.ToLower(targetUoM))

	answer, err := ConvertUnits(q.InputUoM, q.TargetUoM, q.Input)
	if err == nil {
		q.CorrectAnswer = &answer
	}

	return q, nil
}

func (ws Worksheet) Key() []*float64 {
	key := make([]*float64, 0, len(ws.Questions))
	for _, q := range ws.Questions {
		key = append(key, q.CorrectAnswer)
	}
	return key
}

func (q *Question) ToGrid() []string {
	correctStr := ""
	if q.CorrectAnswer != nil {
		correctStr = strconv.FormatFloat(*q.CorrectAnswer, 'f', -1, 64)
	}
	return []string{strconv.FormatFloat(q.Input, 'f', -1, 64), q.InputUoM, q.TargetUoM, correctStr}
}
