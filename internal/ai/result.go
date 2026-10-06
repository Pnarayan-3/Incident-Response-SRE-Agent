package ai

type IncidentAnalysis struct {
	Severity            string   `json:"severity"`
	RootCause           string   `json:"root_cause"`
	Evidence             []string `json:"evidence"`
	Confidence           float64  `json:"confidence"`
	Impact              string   `json:"impact"`
	RecommendedActions  []string `json:"recommended_actions"`
	HumanReviewRequired bool     `json:"human_review_required"`
}