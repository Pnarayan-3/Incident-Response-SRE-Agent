# Incident Response Agent

AI-powered SRE agent for incident analysis, root-cause detection, remediation recommendations, and human-approved incident response.

## Overview

The **Incident Response Agent** is a Go-based AI-powered SRE tool designed to assist engineers during production incidents.

The agent accepts an incident alert, collects logs and observability metrics, sends the collected context to Configurable Generative AI Model for analysis, validates the AI response, evaluates remediation risk, and determines whether human review is required.

The agent can also integrate with **Prometheus** for metrics collection and **Slack** for incident notifications.

The system is designed around a safety-first principle:

> AI can analyze incidents and recommend actions, but production-impacting remediation should require human approval.

---

## Problem

Production incidents often require engineers to quickly correlate:

* Monitoring alerts
* Application logs
* Service metrics
* Error rates
* Latency
* Resource utilization
* Possible root causes
* Remediation options

Manually correlating this information can increase the time required to understand and respond to an incident.

This project demonstrates how an AI-assisted SRE workflow can reduce that analysis effort while keeping human approval in the loop for risky decisions.

---

## Features

### Incident Processing
### Observability
### AI Analysis
### Safety Controls
### Remediation Risk Classification

Recommended actions are classified as:

| Risk   | Example                    | Approval     |
| ------ | -------------------------- | ------------ |
| LOW    | Review logs                | Not required |
| MEDIUM | Modify configuration       | Required     |
| HIGH   | Restart service / rollback | Required     |

The classification is intentionally conservative.

---

## Slack Notifications

The agent can optionally send incident analysis notifications to Slack.

Example:

```text
🚨 Incident Response Alert

Incident: ALERT-001
Title: Payment API High Error Rate
Service: payment-service
Severity: HIGH
Confidence: 0.91
Human Review Required: true

Root Cause
Database connection pool exhaustion.

Evidence
• Database connection timeout
• Connection pool usage reached 100%
• HTTP 500 error rate reached 18%
• API latency increased to 4.2 seconds

Recommended Actions
• Investigate database connection saturation — Risk: LOW
• Increase database connection capacity — Risk: MEDIUM — Approval Required
```

Slack notifications are optional and disabled by default.

---

## Prometheus Integration
## CI/CD

GitHub Actions validates the project automatically.

The CI pipeline performs:

```text
Checkout
   ↓
Setup Go 1.25
   ↓
Formatting Check
   ↓
Unit Tests
   ↓
Coverage
   ↓
Build
   ↓
Docker Build
```

Coverage is also generated as a CI artifact.

---

## Docker

The application uses a multi-stage Docker build.

```text
Go Builder Image
       ↓
Compile static binary
       ↓
Minimal Alpine Runtime
       ↓
Non-root user
       ↓
Incident Agent
```

Build the image with:

```bash
docker build -t incident-response-agent .
```

The container runs the compiled agent binary as a non-root user.

---
