package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	httpClient *http.Client
	config     Config
}

type Config struct {
	URL       string
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func NewClient(c Config) *Client {
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
		return errors.Wrap(err, "failed to build request")
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("failed to send messages. status code: %s", res.Status)
	}

	return nil
}
