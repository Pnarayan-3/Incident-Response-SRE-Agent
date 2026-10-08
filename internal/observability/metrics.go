package observability

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type MetricsCollector struct {
	PrometheusEnabled bool
	PrometheusURL     string
	PrometheusQuery   string
	HTTPClient        *http.Client
}

type prometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string             `json:"resultType"`
		Result     []prometheusResult `json:"result"`
	} `json:"data"`
	ErrorType string `json:"errorType"`
	Error     string `json:"error"`
}

type prometheusResult struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
}

func NewMetricsCollector(
	prometheusEnabled bool,
	prometheusURL string,
	prometheusQuery string,
) *MetricsCollector {

	return &MetricsCollector{
		PrometheusEnabled: prometheusEnabled,
		PrometheusURL:     strings.TrimRight(prometheusURL, "/"),
		PrometheusQuery:   prometheusQuery,
		HTTPClient:        &http.Client{},
	}
}

func (m *MetricsCollector) Collect(
	service string,
) (string, error) {

	if service == "" {
		return "", fmt.Errorf(
			"service name cannot be empty",
		)
	}

	if m.PrometheusEnabled {
		return m.collectFromPrometheus(service)
	}

	return m.collectMock(service)
}

func (m *MetricsCollector) collectMock(
	service string,
) (string, error) {

	metrics := fmt.Sprintf(`
Service: %s

HTTP 500 Error Rate: 18%%
API Latency: 4.2 seconds
Database Connection Usage: 100%%
CPU Usage: 42%%
Memory Usage: 61%%
`, service)

	return metrics, nil
}

func (m *MetricsCollector) collectFromPrometheus(
	service string,
) (string, error) {

	if m.PrometheusURL == "" {
		return "", fmt.Errorf(
			"Prometheus URL cannot be empty",
		)
	}

	if m.PrometheusQuery == "" {
		return "", fmt.Errorf(
			"Prometheus query cannot be empty",
		)
	}

	queryURL := m.PrometheusURL + "/api/v1/query"

	requestURL, err := url.Parse(queryURL)
	if err != nil {
		return "", fmt.Errorf(
			"invalid Prometheus URL: %w",
			err,
		)
	}

	query := m.PrometheusQuery

	requestQuery := requestURL.Query()
	requestQuery.Set("query", query)
	requestURL.RawQuery = requestQuery.Encode()

	req, err := http.NewRequest(
		http.MethodGet,
		requestURL.String(),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create Prometheus request: %w",
			err,
		)
	}

	resp, err := m.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf(
			"Prometheus request failed: %w",
			err,
		)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"Prometheus returned status %d",
			resp.StatusCode,
		)
	}

	var result prometheusResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf(
			"failed to parse Prometheus response: %w",
			err,
		)
	}

	if result.Status != "success" {
		return "", fmt.Errorf(
			"Prometheus query failed: %s",
			result.Error,
		)
	}

	var output strings.Builder

	fmt.Fprintf(
		&output,
		"Service: %s\n\n",
		service,
	)

	fmt.Fprintf(
		&output,
		"Prometheus Query: %s\n",
		query,
	)

	fmt.Fprintf(
		&output,
		"Result Type: %s\n\n",
		result.Data.ResultType,
	)

	if len(result.Data.Result) == 0 {
		output.WriteString("No metric results returned.\n")
		return output.String(), nil
	}

	output.WriteString("Results:\n")

	for _, metric := range result.Data.Result {

		fmt.Fprintf(
			&output,
			"- Metric: %v\n",
			metric.Metric,
		)

		if len(metric.Value) >= 2 {
			fmt.Fprintf(
				&output,
				"  Value: %v\n",
				metric.Value[1],
			)
		}
	}

	return output.String(), nil
}
