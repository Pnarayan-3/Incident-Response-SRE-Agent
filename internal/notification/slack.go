package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackNotifier struct {
	Enabled    bool
	WebhookURL string
	HTTPClient *http.Client
}

type slackPayload struct {
	Text string `json:"text"`
}

func NewSlackNotifier(
	enabled bool,
	webhookURL string,
) *SlackNotifier {

	return &SlackNotifier{
		Enabled:    enabled,
		WebhookURL: webhookURL,
		HTTPClient: &http.Client{},
	}
}

func (s *SlackNotifier) Notify(
	message string,
) error {

	if !s.Enabled {
		return nil
	}

	if s.WebhookURL == "" {
		return fmt.Errorf(
			"Slack webhook URL is not configured",
		)
	}

	payload := slackPayload{
		Text: message,
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf(
			"failed to encode Slack payload: %w",
			err,
		)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		s.WebhookURL,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return fmt.Errorf(
			"failed to create Slack request: %w",
			err,
		)
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf(
			"Slack notification failed: %w",
			err,
		)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf(
			"Slack returned status %d",
			resp.StatusCode,
		)
	}

	return nil
}
