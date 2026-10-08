package observability

import (
	"strings"
	"testing"
)

func TestMetricsCollectorCollect(t *testing.T) {

	collector := NewMetricsCollector(
		false,
		"",
		"",
	)

	metrics, err := collector.Collect("payment-service")

	if err != nil {
		t.Fatalf(
			"expected no error, got: %v",
			err,
		)
	}

	if metrics == "" {
		t.Fatal("expected metrics output, got empty string")
	}

	if !strings.Contains(
		metrics,
		"HTTP 500 Error Rate",
	) {
		t.Error(
			"expected metrics output to contain HTTP 500 Error Rate",
		)
	}

	if !strings.Contains(
		metrics,
		"Database Connection Usage",
	) {
		t.Error(
			"expected metrics output to contain Database Connection Usage",
		)
	}
}

func TestMetricsCollectorEmptyService(t *testing.T) {

	collector := NewMetricsCollector(
		false,
		"",
		"",
	)

	_, err := collector.Collect("")

	if err == nil {
		t.Fatal(
			"expected error for empty service name",
		)
	}
}
