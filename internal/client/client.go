package client

import (
	"flag"
	"log"

	"github.com/dougdoenges/flexion-coding-challenge/internal/app"
	"github.com/dougdoenges/flexion-coding-challenge/internal/parser/file"
)

func Run() {
	worksheetFile := flag.String("worksheet", "", "Give file path for worksheet (required)")
	responsesFile := flag.String("responses", "", "Give file path for student responses to grade (required)")
	outputLocation := flag.String("output", "", "Give file path and name for output (required)")
	flag.Parse()

	if *worksheetFile == "" {
		log.Fatal("worksheet file path is required")
	}
	if *responsesFile == "" {
		log.Fatal("response file path is required")
	}
	if *outputLocation == "" {
		log.Fatal("output file is required")
	}

	worksheetReader, err := file.NewReader[app.Worksheet](*worksheetFile)
	if err != nil {
		log.Fatal(err)
	}
	worksheet, err := worksheetReader.Read(app.NewWorksheet)
	if err != nil {
		log.Fatal(err)
	}

	submissionReader, err := file.NewReader[[]app.Submission](*responsesFile)
	if err != nil {
		log.Fatal(err)
	}
	submissions, err := submissionReader.Read(app.NewSubmissionList)
	if err != nil {
		log.Fatal(err)
	}

	results := app.GetResults(worksheet, submissions)

	resultData := results.ToGridDisplay()
	outputWriter, err := file.NewWriter(*outputLocation)
	if err != nil {
		log.Fatal(err)
	}
	err = outputWriter.Write(resultData)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Success! Graded results can be found here: %s", *outputLocation)
}
