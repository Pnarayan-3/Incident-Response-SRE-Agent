package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	GeminiModel         string
	MaxRetries          int
	RetryDelaySeconds   int
	ConfidenceThreshold float64
	PrometheusEnabled   bool
	PrometheusURL       string
	PrometheusQuery     string
	SlackEnabled        bool
	SlackWebhookURL     string
}

func Load() (*Config, error) {

	model := os.Getenv("GEMINI_MODEL")
	if model == "" {
		model = "gemini-3.5-flash-lite"
	}

	maxRetries := 3

	if value := os.Getenv("GEMINI_MAX_RETRIES"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid GEMINI_MAX_RETRIES: %w",
				err,
			)
		}

		maxRetries = parsed
	}

	if maxRetries < 1 {
		return nil, fmt.Errorf(
			"GEMINI_MAX_RETRIES must be at least 1",
		)
	}

	retryDelay := 1

	if value := os.Getenv("GEMINI_RETRY_DELAY_SECONDS"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid GEMINI_RETRY_DELAY_SECONDS: %w",
				err,
			)
		}

		retryDelay = parsed
	}

	if retryDelay < 0 {
		return nil, fmt.Errorf(
			"GEMINI_RETRY_DELAY_SECONDS cannot be negative",
		)
	}

	confidenceThreshold := 0.80

	if value := os.Getenv("INCIDENT_CONFIDENCE_THRESHOLD"); value != "" {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid INCIDENT_CONFIDENCE_THRESHOLD: %w",
				err,
			)
		}

		confidenceThreshold = parsed
	}

	if confidenceThreshold < 0 || confidenceThreshold > 1 {
		return nil, fmt.Errorf(
			"INCIDENT_CONFIDENCE_THRESHOLD must be between 0 and 1",
		)
	}

	prometheusEnabled := false

	if value := os.Getenv("PROMETHEUS_ENABLED"); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid PROMETHEUS_ENABLED: %w",
				err,
			)
		}

		prometheusEnabled = parsed
	}

	prometheusURL := os.Getenv("PROMETHEUS_URL")
	if prometheusURL == "" {
		prometheusURL = "http://localhost:9090"
	}

	prometheusQuery := os.Getenv("PROMETHEUS_QUERY")
	if prometheusQuery == "" {
		prometheusQuery = "up"
	}

	slackEnabled := false

	if value := os.Getenv("SLACK_ENABLED"); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf(
				"invalid SLACK_ENABLED: %w",
				err,
			)
		}

		slackEnabled = parsed
	}

	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")

	return &Config{
		GeminiModel:         model,
		MaxRetries:          maxRetries,
		RetryDelaySeconds:   retryDelay,
		ConfidenceThreshold: confidenceThreshold,
		PrometheusEnabled:   prometheusEnabled,
		PrometheusURL:       prometheusURL,
		PrometheusQuery:     prometheusQuery,
		SlackEnabled:        slackEnabled,
		SlackWebhookURL:     slackWebhookURL,
	}, nil
}
