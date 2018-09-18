package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/satori/go.uuid"
)

func TestHandler(t *testing.T) {
	snsEvent := events.SNSEvent{
		Records: []events.SNSEventRecord{
			{
				SNS: events.SNSEntity{
					MessageID: uuid.Must(uuid.NewV4(), nil).String(),
					Message:   "テストメッセージ",
				},
			},
		},
	}

	if err := handler(snsEvent); err != nil {
		t.Errorf("error: %s", err)
	}
}
