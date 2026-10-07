package incident

import (
	"fmt"
	"strings"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/remediation"
)

func BuildReport(
	incidentData *Incident,
	analysis *ai.IncidentAnalysis,
	recommendations []remediation.Recommendation,
) string {

	var report strings.Builder

	report.WriteString("========================================\n")
	report.WriteString("        INCIDENT RESPONSE REPORT\n")
	report.WriteString("========================================\n\n")

	fmt.Fprintf(
		&report,
		"Incident ID: %s\n",
		incidentData.ID,
	)

	fmt.Fprintf(
		&report,
		"Title: %s\n",
		incidentData.Title,
	)

	fmt.Fprintf(
		&report,
		"Service: %s\n",
		incidentData.Service,
	)

	fmt.Fprintf(
		&report,
		"Source: %s\n",
		incidentData.Source,
	)

	fmt.Fprintf(
		&report,
		"Timestamp: %s\n",
		incidentData.Timestamp,
	)

	fmt.Fprintf(
		&report,
		"Severity: %s\n",
		analysis.Severity,
	)

	fmt.Fprintf(
		&report,
		"Confidence: %.2f\n",
		analysis.Confidence,
	)

	fmt.Fprintf(
		&report,
		"Human Review Required: %t\n",
		analysis.HumanReviewRequired,
	)

	report.WriteString("\n----------------------------------------\n")
	report.WriteString("ROOT CAUSE\n")
	report.WriteString("----------------------------------------\n")
	report.WriteString(analysis.RootCause)
	report.WriteString("\n")

	report.WriteString("\n----------------------------------------\n")
	report.WriteString("EVIDENCE\n")
	report.WriteString("----------------------------------------\n")

	for _, evidence := range analysis.Evidence {
		fmt.Fprintf(
			&report,
			"- %s\n",
			evidence,
		)
	}

	report.WriteString("\n----------------------------------------\n")
	report.WriteString("IMPACT\n")
	report.WriteString("----------------------------------------\n")
	report.WriteString(analysis.Impact)
	report.WriteString("\n")

	report.WriteString("\n----------------------------------------\n")
	report.WriteString("REMEDIATION\n")
	report.WriteString("----------------------------------------\n")

	for _, recommendation := range recommendations {

		fmt.Fprintf(
			&report,
			"- %s\n",
			recommendation.Action,
		)

		fmt.Fprintf(
			&report,
			"  Risk: %s\n",
			recommendation.Risk,
		)

		fmt.Fprintf(
			&report,
			"  Approval Required: %t\n",
			recommendation.RequiresApproval,
		)
	}

	report.WriteString("\n========================================\n")

	return report.String()
}
