package observability

import (
	"testing"
)

func TestConvertToIncident(t *testing.T) {

	alert := &Alert{
		ID:          "ALERT-001",
		Name:        "Payment API High Error Rate",
		Description: "Payment API is returning a high number of HTTP 500 errors.",
		Service:     "payment-service",
		Severity:    "HIGH",
		Source:      "CloudWatch",
		Timestamp:   "2026-10-06T20:30:00Z",
	}

	incident, err := ConvertToIncident(alert)

	if err != nil {
		t.Fatalf(
			"expected no error, got: %v",
			err,
		)
	}

	if incident == nil {
		t.Fatal(
			"expected incident, got nil",
		)
	}

	if incident.ID != alert.ID {
		t.Fatalf(
			"expected ID %s, got %s",
			alert.ID,
			incident.ID,
		)
	}

	if incident.Title != alert.Name {
		t.Fatalf(
			"expected title %s, got %s",
			alert.Name,
			incident.Title,
		)
	}

	if incident.Description != alert.Description {
		t.Fatalf(
			"expected description %s, got %s",
			alert.Description,
			incident.Description,
		)
	}

	if incident.Service != alert.Service {
		t.Fatalf(
			"expected service %s, got %s",
			alert.Service,
			incident.Service,
		)
	}

	if incident.Severity != alert.Severity {
		t.Fatalf(
			"expected severity %s, got %s",
			alert.Severity,
			incident.Severity,
		)
	}

	if incident.Source != alert.Source {
		t.Fatalf(
			"expected source %s, got %s",
			alert.Source,
			incident.Source,
		)
	}

	if incident.Timestamp != alert.Timestamp {
		t.Fatalf(
			"expected timestamp %s, got %s",
			alert.Timestamp,
			incident.Timestamp,
		)
	}

	if incident.Status != "OPEN" {
		t.Fatalf(
			"expected status OPEN, got %s",
			incident.Status,
		)
	}
}

func TestConvertToIncidentNilAlert(t *testing.T) {

	incident, err := ConvertToIncident(nil)

	if err == nil {
		t.Fatal(
			"expected error for nil alert",
		)
	}

	if incident != nil {
		t.Fatalf(
			"expected nil incident, got %+v",
			incident,
		)
	}
}

func TestConvertToIncidentMissingID(t *testing.T) {

	alert := &Alert{
		Name:    "Payment API High Error Rate",
		Service: "payment-service",
	}

	incident, err := ConvertToIncident(alert)

	if err == nil {
		t.Fatal(
			"expected error for missing alert ID",
		)
	}

	if incident != nil {
		t.Fatalf(
			"expected nil incident, got %+v",
			incident,
		)
	}
}

func TestConvertToIncidentMissingName(t *testing.T) {

	alert := &Alert{
		ID:      "ALERT-001",
		Service: "payment-service",
	}

	incident, err := ConvertToIncident(alert)

	if err == nil {
		t.Fatal(
			"expected error for missing alert name",
		)
	}

	if incident != nil {
		t.Fatalf(
			"expected nil incident, got %+v",
			incident,
		)
	}
}

func TestConvertToIncidentMissingService(t *testing.T) {

	alert := &Alert{
		ID:   "ALERT-001",
		Name: "Payment API High Error Rate",
	}

	incident, err := ConvertToIncident(alert)

	if err == nil {
		t.Fatal(
			"expected error for missing service",
		)
	}

	if incident != nil {
		t.Fatalf(
			"expected nil incident, got %+v",
			incident,
		)
	}
}
