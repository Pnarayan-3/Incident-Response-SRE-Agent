package observability

import (
	"fmt"
)

type MetricsCollector struct{}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{}
}

func (m *MetricsCollector) Collect(
	service string,
) (string, error) {

	if service == "" {
		return "", fmt.Errorf(
			"service name cannot be empty",
		)
	}

	// Temporary implementation.
	// This will later be replaced with a real
	// Prometheus / CloudWatch metrics integration.

	metrics := fmt.Sprintf(`
Service: %s

HTTP 500 Error Rate: 18%%
API Latency: 4.2 seconds
Database Connection Usage: 100%%
CPU Usage: 42%%
Memory Usage: 61%%
`, service)

	return metrics, nil
}