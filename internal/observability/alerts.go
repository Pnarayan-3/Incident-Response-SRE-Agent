package observability

import (
	"fmt"

	"github.com/Pnarayan-3/Incident-Response-Agent/internal/incident"
)

type Alert struct {
	ID          string
	Name        string
	Description string
	Service     string
	Severity    string
	Source      string
	Timestamp   string
}

func ConvertToIncident(
	alert *Alert,
) (*incident.Incident, error) {

	if alert == nil {
		return nil, fmt.Errorf("alert cannot be nil")
	}

	if alert.ID == "" {
		return nil, fmt.Errorf("alert ID cannot be empty")
	}

	if alert.Name == "" {
		return nil, fmt.Errorf("alert name cannot be empty")
	}

	if alert.Service == "" {
		return nil, fmt.Errorf("alert service cannot be empty")
	}

	return &incident.Incident{
		ID:          alert.ID,
		Title:       alert.Name,
		Description: alert.Description,
		Service:     alert.Service,
		Severity:    alert.Severity,
		Status:      "OPEN",
		Source:      alert.Source,
		Timestamp:   alert.Timestamp,
	}, nil
}
