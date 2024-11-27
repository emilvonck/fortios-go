package fortios

import (
	"context"
	"fmt"
	"net/http"
)

type Results struct {
	ModelName     string `json:"model_name"`
	ModelNumber   string `json:"model_number"`
	Model         string `json:"model"`
	Hostname      string `json:"hostname"`
	LogDiskStatus string `json:"log_disk_status"`
}

type SystemStatusResponse struct {
	HTTPMethod string  `json:"http_method"`
	Results    Results `json:"results"`
	Vdom       string  `json:"vdom"`
	Path       string  `json:"path"`
	Name       string  `json:"name"`
	Action     string  `json:"action"`
	Status     string  `json:"status"`
	Serial     string  `json:"serial"`
	Version    string  `json:"version"`
	Build      int     `json:"build"`
}

func (c *Client) GetSystemStatus(ctx context.Context) (*SystemStatusResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/monitor/system/status", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res SystemStatusResponse
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
