package incident

import (
	"fmt"
	"strings"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
)

func ValidateAnalysis(
	result *ai.IncidentAnalysis,
) error {

	if result == nil {
		return fmt.Errorf("incident analysis is nil")
	}

	switch result.Severity {
	case "LOW", "MEDIUM", "HIGH", "CRITICAL":
	default:
		return fmt.Errorf(
			"invalid severity: %s",
			result.Severity,
		)
	}

	if strings.TrimSpace(result.RootCause) == "" {
		return fmt.Errorf(
			"root cause cannot be empty",
		)
	}

	if strings.TrimSpace(result.Impact) == "" {
		return fmt.Errorf(
			"impact cannot be empty",
		)
	}

	if result.Confidence < 0 || result.Confidence > 1 {
		return fmt.Errorf(
			"confidence must be between 0 and 1: %.2f",
			result.Confidence,
		)
	}

	if len(result.Evidence) == 0 {
		return fmt.Errorf(
			"at least one evidence item is required",
		)
	}

	if len(result.RecommendedActions) == 0 {
		return fmt.Errorf(
			"at least one recommended action is required",
		)
	}

	for _, evidence := range result.Evidence {
		if strings.TrimSpace(evidence) == "" {
			return fmt.Errorf(
				"evidence cannot be empty",
			)
		}
	}

	for _, action := range result.RecommendedActions {
		if strings.TrimSpace(action) == "" {
			return fmt.Errorf(
				"recommended action cannot be empty",
			)
		}
	}

	return nil
}
