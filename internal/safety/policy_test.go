package safety

import (
	"testing"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/remediation"
)

func TestRequiresHumanReview(t *testing.T) {

	policy := NewPolicy(0.80)

	tests := []struct {
		name            string
		result          *ai.IncidentAnalysis
		recommendations []remediation.Recommendation
		expected        bool
	}{
		{
			name: "high confidence low risk incident",
			result: &ai.IncidentAnalysis{
				Severity:   "LOW",
				Confidence: 0.95,
			},
			recommendations: []remediation.Recommendation{
				{
					Action:           "Review application logs",
					Risk:             "LOW",
					RequiresApproval: false,
				},
			},
			expected: false,
		},
		{
			name: "low confidence incident",
			result: &ai.IncidentAnalysis{
				Severity:   "MEDIUM",
				Confidence: 0.60,
			},
			recommendations: []remediation.Recommendation{
				{
					Action:           "Review application logs",
					Risk:             "LOW",
					RequiresApproval: false,
				},
			},
			expected: true,
		},
		{
			name: "critical incident",
			result: &ai.IncidentAnalysis{
				Severity:   "CRITICAL",
				Confidence: 0.98,
			},
			recommendations: []remediation.Recommendation{
				{
					Action:           "Review application logs",
					Risk:             "LOW",
					RequiresApproval: false,
				},
			},
			expected: true,
		},
		{
			name: "medium risk remediation",
			result: &ai.IncidentAnalysis{
				Severity:   "MEDIUM",
				Confidence: 0.95,
			},
			recommendations: []remediation.Recommendation{
				{
					Action:           "Increase connection pool size",
					Risk:             "MEDIUM",
					RequiresApproval: true,
				},
			},
			expected: true,
		},
		{
			name: "high risk remediation",
			result: &ai.IncidentAnalysis{
				Severity:   "HIGH",
				Confidence: 0.95,
			},
			recommendations: []remediation.Recommendation{
				{
					Action:           "Restart the service",
					Risk:             "HIGH",
					RequiresApproval: true,
				},
			},
			expected: true,
		},
		{
			name:     "nil analysis",
			result:   nil,
			recommendations: []remediation.Recommendation{},
			expected: true,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			actual := policy.RequiresHumanReview(
				test.result,
				test.recommendations,
			)

			if actual != test.expected {
				t.Fatalf(
					"expected human review %t, got %t",
					test.expected,
					actual,
				)
			}
		})
	}
}

