package main

import (
	"fmt"
	"log"

	"github.com/Pnarayan-3/Incident-Response-Agent/config"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/incident"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/logger"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/observability"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/notification"
)

func main() {

	logger.Info(
		"Starting Incident Response Agent",
	)

	cfg, err := config.Load()
	if err != nil {
		logger.Error(
			"Configuration error: " + err.Error(),
		)

		log.Fatalf(
			"configuration error: %v",
			err,
		)
	}

	fmt.Printf(
		"AI Model: %s\n",
		cfg.GeminiModel,
	)

	aiClient := ai.NewClient(
		cfg.GeminiModel,
		cfg.MaxRetries,
		cfg.RetryDelaySeconds,
	)

	analyzer := incident.NewAnalyzer(
		aiClient,
		cfg,
	)

	logCollector := observability.NewLogCollector()
	metricsCollector := observability.NewMetricsCollector(
		cfg.PrometheusEnabled,
		cfg.PrometheusURL,
		cfg.PrometheusQuery,
	)

	slackNotifier := notification.NewSlackNotifier(
		cfg.SlackEnabled,
		cfg.SlackWebhookURL,
	)


	alert := &observability.Alert{
		ID:          "ALERT-001",
		Name:        "Payment API High Error Rate",
		Description: "The payment API has exceeded the configured HTTP 500 error threshold.",
		Service:     "payment-service",
		Severity:    "HIGH",
		Source:      "CloudWatch",
		Timestamp:   "2026-10-06T20:30:00Z",
	}

	logger.Info(
		"Processing incident alert: " + alert.ID,
	)

	currentIncident, err := observability.ConvertToIncident(
		alert,
	)
	if err != nil {
		logger.Error(
			"Failed to create incident: " + err.Error(),
		)

		log.Fatalf(
			"failed to create incident: %v",
			err,
		)
	}

	logger.Info(
		"Collecting logs for service: " +
			currentIncident.Service,
	)

	logs, err := logCollector.Collect(
		currentIncident.Service,
	)
	if err != nil {
		logger.Error(
			"Failed to collect logs: " + err.Error(),
		)

		log.Fatalf(
			"failed to collect logs: %v",
			err,
		)
	}

	logger.Info(
		"Collecting metrics for service: " +
			currentIncident.Service,
	)

	metrics, err := metricsCollector.Collect(
		currentIncident.Service,
	)
	if err != nil {
		logger.Error(
			"Failed to collect metrics: " + err.Error(),
		)

		log.Fatalf(
			"failed to collect metrics: %v",
			err,
		)
	}

	logger.Info(
		"Starting AI incident analysis",
	)

	result, recommendations, err := analyzer.Analyze(
		currentIncident,
		logs,
		metrics,
	)
	if err != nil {
		logger.Error(
			"Incident analysis failed: " + err.Error(),
		)

		log.Fatalf(
			"incident analysis failed: %v",
			err,
		)
	}

	logger.Info(
		"Incident analysis completed",
	)

	if result.HumanReviewRequired {
		logger.Warn(
			"Human review is required for this incident",
		)
	}

	report := incident.BuildReport(
		currentIncident,
		result,
		recommendations,
	)

	slackMessage := notification.BuildIncidentMessage(
		currentIncident,
		result,
		recommendations,
	)

	if err := slackNotifier.Notify(slackMessage); err != nil {
		logger.Warn(
			"Slack notification failed: " + err.Error(),
		)
	} else if cfg.SlackEnabled {
		logger.Info(
			"Incident notification sent to Slack",
		)
	}

	fmt.Println()
	fmt.Println(report)
}
