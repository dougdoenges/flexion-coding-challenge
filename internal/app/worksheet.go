package app

import (
	"errors"
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
		return Question{}, errors.New("invalid question provided")
	}

	input, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return Question{}, errors.New("invalid input number given")
	}
	q.Input = input

	inputUom := data[1]
	q.InputUoM = strings.TrimSpace(strings.ToLower(inputUom))

	targetUoM := data[2]
	q.TargetUoM = strings.TrimSpace(strings.ToLower(targetUoM))

	converter, err := NewConverter(q.InputUoM, q.TargetUoM)
	if err == nil {
		answer := converter.Convert(q.Input)
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
	return []string{strconv.FormatFloat(q.Input, 'f', -1, 64), q.InputUoM, q.TargetUoM, correctStr, ""}
}
