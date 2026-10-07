package incident

import (
	"testing"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
)

func validAnalysis() *ai.IncidentAnalysis {
	return &ai.IncidentAnalysis{
		Severity:   "HIGH",
		RootCause:  "Database connection pool was exhausted.",
		Confidence: 0.92,
		Evidence: []string{
			"Database connection usage reached 100%.",
			"Application logs show database connection timeouts.",
		},
		Impact: "Payment requests are returning HTTP 500 errors.",
		RecommendedActions: []string{
			"Review database connection pool configuration.",
			"Investigate database connection saturation.",
		},
		HumanReviewRequired: false,
	}
}

func TestValidateAnalysisValid(t *testing.T) {

	result := validAnalysis()

	err := ValidateAnalysis(result)

	if err != nil {
		t.Fatalf(
			"expected valid analysis, got error: %v",
			err,
		)
	}
}

func TestValidateAnalysisNil(t *testing.T) {

	err := ValidateAnalysis(nil)

	if err == nil {
		t.Fatal(
			"expected error for nil analysis",
		)
	}
}

func TestValidateAnalysisInvalidSeverity(t *testing.T) {

	result := validAnalysis()
	result.Severity = "URGENT"

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for invalid severity",
		)
	}
}

func TestValidateAnalysisEmptyRootCause(t *testing.T) {

	result := validAnalysis()
	result.RootCause = ""

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for empty root cause",
		)
	}
}

func TestValidateAnalysisEmptyImpact(t *testing.T) {

	result := validAnalysis()
	result.Impact = ""

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for empty impact",
		)
	}
}

func TestValidateAnalysisInvalidConfidence(t *testing.T) {

	tests := []struct {
		name       string
		confidence float64
	}{
		{
			name:       "confidence below zero",
			confidence: -0.1,
		},
		{
			name:       "confidence above one",
			confidence: 1.1,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result := validAnalysis()
			result.Confidence = test.confidence

			err := ValidateAnalysis(result)

			if err == nil {
				t.Fatalf(
					"expected error for confidence %.2f",
					test.confidence,
				)
			}
		})
	}
}

func TestValidateAnalysisEmptyEvidence(t *testing.T) {

	result := validAnalysis()
	result.Evidence = nil

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for empty evidence",
		)
	}
}

func TestValidateAnalysisEmptyEvidenceItem(t *testing.T) {

	result := validAnalysis()
	result.Evidence = []string{
		"Database connection usage reached 100%.",
		"",
	}

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for empty evidence item",
		)
	}
}

func TestValidateAnalysisEmptyRecommendations(t *testing.T) {

	result := validAnalysis()
	result.RecommendedActions = nil

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for empty recommendations",
		)
	}
}

func TestValidateAnalysisEmptyRecommendationItem(t *testing.T) {

	result := validAnalysis()
	result.RecommendedActions = []string{
		"Review database configuration.",
		"",
	}

	err := ValidateAnalysis(result)

	if err == nil {
		t.Fatal(
			"expected error for empty recommendation",
		)
	}
}
