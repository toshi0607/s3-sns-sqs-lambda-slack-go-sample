package main

import (
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	inputJson := readJsonFromFile(t, "../../testdata/s3event.json")
	sqsEvent := events.SQSEvent{
		Records: []events.SQSMessage{
			{
				Body: string(inputJson),
			},
		},
	}

	if err := handler(sqsEvent); err != nil {
		t.Errorf("error: %s", err)
	}
}

func readJsonFromFile(t *testing.T, inputFile string) []byte {
	inputJson, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJson
}
