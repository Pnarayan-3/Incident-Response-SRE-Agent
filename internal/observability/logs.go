package observability

import (
	"fmt"
)

type LogCollector struct{}

func NewLogCollector() *LogCollector {
	return &LogCollector{}
}

func (l *LogCollector) Collect(
	service string,
) (string, error) {

	if service == "" {
		return "", fmt.Errorf(
			"service name cannot be empty",
		)
	}

	// Temporary implementation.
	// This will later be replaced with a real
	// CloudWatch / log provider integration.

	logs := fmt.Sprintf(`
Service: %s

2026-10-06 20:28:12 ERROR
database connection timeout

2026-10-06 20:28:15 ERROR
failed to acquire database connection

2026-10-06 20:29:01 ERROR
request failed with status 500

2026-10-06 20:29:15 ERROR
database connection pool exhausted
`, service)

	return logs, nil
}
