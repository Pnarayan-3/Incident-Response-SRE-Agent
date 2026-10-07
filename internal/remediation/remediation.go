package remediation

import (
	"strings"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
)

type Recommendation struct {
	Action           string
	Risk             string
	RequiresApproval bool
}

func BuildRecommendations(
	result *ai.IncidentAnalysis,
) []Recommendation {

	recommendations := make(
		[]Recommendation,
		0,
		len(result.RecommendedActions),
	)

	for _, action := range result.RecommendedActions {

		risk := determineRisk(action)

		recommendations = append(
			recommendations,
			Recommendation{
				Action:           action,
				Risk:             risk,
				RequiresApproval: risk != "LOW",
			},
		)
	}

	return recommendations
}

func determineRisk(action string) string {

	action = strings.ToLower(action)

	highRiskKeywords := []string{
		"delete",
		"drop",
		"terminate",
		"restart",
		"scale down",
		"shutdown",
		"rollback",
	}

	for _, keyword := range highRiskKeywords {
		if strings.Contains(action, keyword) {
			return "HIGH"
		}
	}

	mediumRiskKeywords := []string{
		"scale",
		"modify",
		"change",
		"increase",
		"decrease",
		"restart",
	}

	for _, keyword := range mediumRiskKeywords {
		if strings.Contains(action, keyword) {
			return "MEDIUM"
		}
	}

	return "LOW"
}
