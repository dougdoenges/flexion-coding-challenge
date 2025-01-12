package client

import (
	"context"
	"flag"
	"log"

	"github.com/dougdoenges/flexion-coding-challenge/internal/app"
	"github.com/dougdoenges/flexion-coding-challenge/internal/parser/file"
)

func Run(ctx context.Context) {
	worksheetFile := flag.String("worksheet", "", "Give file path for worksheet (required)")
	responsesFile := flag.String("responses", "", "Give file path for student responses to grade (required)")
	outputLocation := flag.String("outputFile", "", "Give file type for output (optional: default in config)")
	flag.Parse()

	if *worksheetFile == "" {
		log.Fatal("worksheet File path is required")
	}
	if *responsesFile == "" {
		log.Fatal("response file path is required")
	}
	if *outputLocation == "" {
		log.Fatal("output location is required")
	}
	// TODO: validate all file types from flags

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
}
