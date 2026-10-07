package observability

import (
	"strings"
	"testing"
)

func TestMetricsCollectorCollect(t *testing.T) {

	collector := NewMetricsCollector()

	metrics, err := collector.Collect(
		"payment-service",
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got: %v",
			err,
		)
	}

	if metrics == "" {
		t.Fatal(
			"expected metrics, got empty string",
		)
	}

	expectedValues := []string{
		"payment-service",
		"HTTP 500 Error Rate: 18%",
		"API Latency: 4.2 seconds",
		"Database Connection Usage: 100%",
		"CPU Usage: 42%",
		"Memory Usage: 61%",
	}

	for _, value := range expectedValues {

		if !strings.Contains(
			metrics,
			value,
		) {
			t.Fatalf(
				"expected metrics to contain %q",
				value,
			)
		}
	}
}

func TestMetricsCollectorEmptyService(t *testing.T) {

	collector := NewMetricsCollector()

	metrics, err := collector.Collect("")

	if err == nil {
		t.Fatal(
			"expected error for empty service name",
		)
	}

	if metrics != "" {
		t.Fatalf(
			"expected empty metrics, got %q",
			metrics,
		)
	}
}

