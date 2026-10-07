package safety

import (
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/remediation"
)

type Policy struct {
	ConfidenceThreshold float64
}

func NewPolicy(
	confidenceThreshold float64,
) *Policy {
	return &Policy{
		ConfidenceThreshold: confidenceThreshold,
	}
}

func (p *Policy) RequiresHumanReview(
	result *ai.IncidentAnalysis,
	recommendations []remediation.Recommendation,
) bool {

	if result == nil {
		return true
	}

	// Critical incidents always require human approval.
	if result.Severity == "CRITICAL" {
		return true
	}

	// Low-confidence analysis requires human approval.
	if result.Confidence < p.ConfidenceThreshold {
		return true
	}

	// Medium- and high-risk remediation actions require approval.
	for _, recommendation := range recommendations {
		if recommendation.RequiresApproval {
			return true
		}
	}

	return false
}
