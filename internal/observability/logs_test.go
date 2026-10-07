package observability

import (
	"strings"
	"testing"
)

func TestLogCollectorCollect(t *testing.T) {

	collector := NewLogCollector()

	logs, err := collector.Collect(
		"payment-service",
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got: %v",
			err,
		)
	}

	if logs == "" {
		t.Fatal(
			"expected logs, got empty string",
		)
	}

	expectedValues := []string{
		"payment-service",
		"database connection timeout",
		"failed to acquire database connection",
		"status 500",
		"connection pool exhausted",
	}

	for _, value := range expectedValues {

		if !strings.Contains(
			logs,
			value,
		) {
			t.Fatalf(
				"expected logs to contain %q",
				value,
			)
		}
	}
}

func TestLogCollectorEmptyService(t *testing.T) {

	collector := NewLogCollector()

	logs, err := collector.Collect("")

	if err == nil {
		t.Fatal(
			"expected error for empty service name",
		)
	}

	if logs != "" {
		t.Fatalf(
			"expected empty logs, got %q",
			logs,
		)
	}
}

