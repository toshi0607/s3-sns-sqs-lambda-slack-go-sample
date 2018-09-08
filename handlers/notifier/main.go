package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var client *Client

func main() {
	lambda.Start(handler)
}

func init() {
	client = NewClient(
		config{
			URL:       os.Getenv("WEBHOOK_URL"),
			Channel:   os.Getenv("CHANNEL"),
			Username:  os.Getenv("USERNAME"),
			IconEmoji: os.Getenv("ICON"),
		},
	)
}

func handler(snsEvent events.SNSEvent) error {
	record := snsEvent.Records[0]
	snsRecord := snsEvent.Records[0].SNS
	fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

	if err := client.PostMessage(snsRecord.Message); err != nil {
		return err
	}
	return nil
}

type Client struct {
	httpClient *http.Client
	config     config
}

type config struct {
	URL       string
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func NewClient(c config) *Client {

	return &Client{
		httpClient: &http.Client{},
		config:     c,
	}
}

func (c Client) PostMessage(message string) error {
	c.config.Text = message
	p, _ := json.Marshal(c.config)

	req, err := http.NewRequest(
		"POST",
		c.config.URL,
		bytes.NewReader(p),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("failed to send messages: %s", res.Status)
	}

	return nil
}
