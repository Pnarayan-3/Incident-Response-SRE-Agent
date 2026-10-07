package remediation

import (
	"testing"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
)

func TestBuildRecommendations(t *testing.T) {

	tests := []struct {
		name             string
		action           string
		expectedRisk     string
		expectedApproval bool
	}{
		{
			name:             "delete database",
			action:           "Delete the failed database",
			expectedRisk:     "HIGH",
			expectedApproval: true,
		},
		{
			name:             "restart service",
			action:           "Restart the payment service",
			expectedRisk:     "HIGH",
			expectedApproval: true,
		},
		{
			name:             "rollback deployment",
			action:           "Rollback the latest deployment",
			expectedRisk:     "HIGH",
			expectedApproval: true,
		},
		{
			name:             "scale down service",
			action:           "Scale down the service",
			expectedRisk:     "HIGH",
			expectedApproval: true,
		},
		{
			name:             "increase replicas",
			action:           "Scale the service to five replicas",
			expectedRisk:     "MEDIUM",
			expectedApproval: true,
		},
		{
			name:             "modify configuration",
			action:           "Modify the database connection configuration",
			expectedRisk:     "MEDIUM",
			expectedApproval: true,
		},
		{
			name:             "review logs",
			action:           "Review application logs for additional evidence",
			expectedRisk:     "LOW",
			expectedApproval: false,
		},
		{
			name:             "check metrics",
			action:           "Check service metrics and error rates",
			expectedRisk:     "LOW",
			expectedApproval: false,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result := &ai.IncidentAnalysis{
				RecommendedActions: []string{
					test.action,
				},
			}

			recommendations := BuildRecommendations(result)

			if len(recommendations) != 1 {
				t.Fatalf(
					"expected 1 recommendation, got %d",
					len(recommendations),
				)
			}

			recommendation := recommendations[0]

			if recommendation.Risk != test.expectedRisk {
				t.Fatalf(
					"expected risk %s, got %s",
					test.expectedRisk,
					recommendation.Risk,
				)
			}

			if recommendation.RequiresApproval !=
				test.expectedApproval {

				t.Fatalf(
					"expected approval %t, got %t",
					test.expectedApproval,
					recommendation.RequiresApproval,
				)
			}
		})
	}
}

func TestBuildRecommendationsNilResult(t *testing.T) {

	recommendations := BuildRecommendations(nil)

	if recommendations != nil {
		t.Fatalf(
			"expected nil recommendations, got %v",
			recommendations,
		)
	}
}

func TestBuildRecommendationsMultipleActions(t *testing.T) {

	result := &ai.IncidentAnalysis{
		RecommendedActions: []string{
			"Review application logs",
			"Increase the service replicas",
			"Restart the payment service",
		},
	}

	recommendations := BuildRecommendations(result)

	if len(recommendations) != 3 {
		t.Fatalf(
			"expected 3 recommendations, got %d",
			len(recommendations),
		)
	}

	if recommendations[0].Risk != "LOW" {
		t.Fatalf(
			"expected first recommendation to be LOW risk, got %s",
			recommendations[0].Risk,
		)
	}

	if recommendations[1].Risk != "MEDIUM" {
		t.Fatalf(
			"expected second recommendation to be MEDIUM risk, got %s",
			recommendations[1].Risk,
		)
	}

	if recommendations[2].Risk != "HIGH" {
		t.Fatalf(
			"expected third recommendation to be HIGH risk, got %s",
			recommendations[2].Risk,
		)
	}
}

