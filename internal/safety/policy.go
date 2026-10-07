package safety

import (
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
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
) bool {

	if result == nil {
		return true
	}

	if result.Severity == "CRITICAL" {
		return true
	}

	if result.Confidence < p.ConfidenceThreshold {
		return true
	}

	return false
}
