package incident

import (
	"fmt"

	"github.com/Pnarayan-3/Incident-Response-Agent/config"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
)

type Analyzer struct {
	AIClient *ai.Client
	Config   *config.Config
}

func NewAnalyzer(
	aiClient *ai.Client,
	cfg *config.Config,
) *Analyzer {
	return &Analyzer{
		AIClient: aiClient,
		Config:   cfg,
	}
}

func (a *Analyzer) Analyze(
	incident *Incident,
	logs string,
	metrics string,
) (*ai.IncidentAnalysis, error) {

	if incident == nil {
		return nil, fmt.Errorf("incident cannot be nil")
	}

	incidentData := fmt.Sprintf(
		`ID: %s
Title: %s
Description: %s
Service: %s
Severity: %s
Status: %s
Source: %s
Timestamp: %s`,
		incident.ID,
		incident.Title,
		incident.Description,
		incident.Service,
		incident.Severity,
		incident.Status,
		incident.Source,
		incident.Timestamp,
	)

	result, err := a.AIClient.Analyze(
		incidentData,
		logs,
		metrics,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"incident analysis failed: %w",
			err,
		)
	}

	if err := ValidateAnalysis(result); err != nil {
		return nil, err
	}

	result.HumanReviewRequired =
		result.Confidence < a.Config.ConfidenceThreshold

	return result, nil
}