package notification

import (
	"fmt"
	"strings"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/ai"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/incident"
	"github.com/Pnarayan-3/Incident-Response-Agent/internal/remediation"
)

func BuildIncidentMessage(
	currentIncident *incident.Incident,
	analysis *ai.IncidentAnalysis,
	recommendations []remediation.Recommendation,
) string {

	var message strings.Builder

	message.WriteString("🚨 *Incident Response Alert*\n\n")

	fmt.Fprintf(
		&message,
		"*Incident:* %s\n",
		currentIncident.ID,
	)

	fmt.Fprintf(
		&message,
		"*Title:* %s\n",
		currentIncident.Title,
	)

	fmt.Fprintf(
		&message,
		"*Service:* %s\n",
		currentIncident.Service,
	)

	fmt.Fprintf(
		&message,
		"*Severity:* %s\n",
		analysis.Severity,
	)

	fmt.Fprintf(
		&message,
		"*Confidence:* %.2f\n",
		analysis.Confidence,
	)

	fmt.Fprintf(
		&message,
		"*Human Review Required:* %t\n\n",
		analysis.HumanReviewRequired,
	)

	message.WriteString("*Root Cause*\n")
	fmt.Fprintf(
		&message,
		"%s\n\n",
		analysis.RootCause,
	)

	message.WriteString("*Evidence*\n")

	for _, evidence := range analysis.Evidence {
		fmt.Fprintf(
			&message,
			"• %s\n",
			evidence,
		)
	}

	message.WriteString("\n*Recommended Actions*\n")

	for _, recommendation := range recommendations {
		fmt.Fprintf(
			&message,
			"• %s — Risk: %s",
			recommendation.Action,
			recommendation.Risk,
		)

		if recommendation.RequiresApproval {
			message.WriteString(" — Approval Required")
		}

		message.WriteString("\n")
	}

	return message.String()
}
