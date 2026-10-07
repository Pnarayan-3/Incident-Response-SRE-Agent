package incident

import (
	"fmt"

	"github.com/Pnarayan-3/Incident-Response-Agent/config"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/logger"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/remediation"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/safety"
)

type Analyzer struct {
	AIClient *ai.Client
	Config   *config.Config
	Policy   *safety.Policy
}

func NewAnalyzer(
	aiClient *ai.Client,
	cfg *config.Config,
) *Analyzer {

	policy := safety.NewPolicy(
		cfg.ConfidenceThreshold,
	)

	return &Analyzer{
		AIClient: aiClient,
		Config:   cfg,
		Policy:   policy,
	}
}

func (a *Analyzer) Analyze(
	incident *Incident,
	logs string,
	metrics string,
) (
	*ai.IncidentAnalysis,
	[]remediation.Recommendation,
	error,
) {

	if incident == nil {
		return nil, nil, fmt.Errorf(
			"incident cannot be nil",
		)
	}

	logger.Info(
		"Sending incident data to AI analyzer",
	)

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
		return nil, nil, fmt.Errorf(
			"incident analysis failed: %w",
			err,
		)
	}

	if err := ValidateAnalysis(result); err != nil {
		return nil, nil, fmt.Errorf(
			"incident analysis validation failed: %w",
			err,
		)
	}

	logger.Info(
		"AI incident analysis validated successfully",
	)

	recommendations := remediation.BuildRecommendations(
		result,
	)

	result.HumanReviewRequired =
		a.Policy.RequiresHumanReview(
			result,
			recommendations,
		)

	if result.HumanReviewRequired {
		logger.Warn(
			"Safety policy requires human review",
		)
	} else {
		logger.Info(
			"Incident passed the safety review policy",
		)
	}

	return result, recommendations, nil
}

