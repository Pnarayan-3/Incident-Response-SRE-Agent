package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Pnarayan-3/Incident-Response-Agent/config"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/incident"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/observability"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/remediation"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	fmt.Println("Starting Incident Response Agent...")
	fmt.Printf("AI Model: %s\n", cfg.GeminiModel)

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
	metricsCollector := observability.NewMetricsCollector()

	alert := &observability.Alert{
		ID:          "ALERT-001",
		Name:        "Payment API High Error Rate",
		Description: "The payment API has exceeded the configured HTTP 500 error threshold.",
		Service:     "payment-service",
		Severity:    "HIGH",
		Source:      "CloudWatch",
		Timestamp:   "2026-10-06T20:30:00Z",
	}

	currentIncident, err := observability.ConvertToIncident(alert)
	if err != nil {
		log.Fatalf("failed to create incident: %v", err)
	}

	logs, err := logCollector.Collect(
		currentIncident.Service,
	)
	if err != nil {
		log.Fatalf("failed to collect logs: %v", err)
	}

	metrics, err := metricsCollector.Collect(
		currentIncident.Service,
	)
	if err != nil {
		log.Fatalf("failed to collect metrics: %v", err)
	}

	result, err := analyzer.Analyze(
		currentIncident,
		logs,
		metrics,
	)

	recommendations := remediation.BuildRecommendations(result)

	fmt.Println("\nRemediation Recommendations:")

	for _, recommendation := range recommendations {

		fmt.Printf(
			"- %s | Risk: %s | Approval Required: %t\n",
			recommendation.Action,
			recommendation.Risk,
			recommendation.RequiresApproval,
		)
	}

	if err != nil {
		log.Fatalf("incident analysis failed: %v", err)
	}

	output, err := json.MarshalIndent(
		result,
		"",
		"  ",
	)
	if err != nil {
		log.Fatalf("failed to format result: %v", err)
	}

	fmt.Println("\nIncident Analysis:")
	fmt.Println(string(output))
}