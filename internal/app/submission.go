package app

import (
	"errors"
	"strconv"
)

type Submission struct {
	StudentName string
	Responses   []*float64
	Decisions   []Decision
}

type Decision string

const (
	Correct   Decision = "Correct"
	Incorrect Decision = "Incorrect"
	Invalid   Decision = "Invalid"
)

func NewSubmissionList(data [][]string) ([]Submission, error) {
	submissions := make([]Submission, 0)

	for _, row := range data {
		submission, err := newSubmission(row)
		if err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}

	return submissions, nil
}

func newSubmission(data []string) (Submission, error) {
	resp := Submission{}

	resp.StudentName = data[0]

	for _, val := range data[1:] {
		if val == "" {
			resp.Responses = append(resp.Responses, nil)
			continue
		}
		valNumber, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return Submission{}, errors.New("invalid input number given")
		}
		resp.Responses = append(resp.Responses, &valNumber)
	}

	return resp, nil
}

func (s *Submission) Grade(ws Worksheet) {
	answerKey := ws.Key()
	for idx := range s.Responses {
		var decision Decision
		if answerKey[idx] == nil {
			decision = Invalid
		} else {
			correct := *answerKey[idx] == roundFunc(*s.Responses[idx])
			if correct {
				decision = Correct
			} else {
				decision = Incorrect
			}
		}
		s.Decisions = append(s.Decisions, decision)
	}
}

func (s *Submission) ToGrid(questionIdx int) []string {
	responseStr := ""
	if s.Responses[questionIdx] != nil {
		responseStr = strconv.FormatFloat(*s.Responses[questionIdx], 'f', -1, 64)
	}
	return []string{responseStr, string(s.Decisions[questionIdx])}
}
