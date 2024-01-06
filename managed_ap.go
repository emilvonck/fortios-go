package fortios

import (
	"context"
	"fmt"
	"net/http"
)

type ManagedApResponse struct {
	Vdom       string        `json:"vdom"`
	Path       string        `json:"path"`
	Name       string        `json:"name"`
	Action     string        `json:"action"`
	Status     string        `json:"status"`
	Serial     string        `json:"serial"`
	Version    string        `json:"version"`
	Build      int           `json:"build"`
	HTTPMethod string        `json:"http_method"`
	Results    []AccessPoint `json:"results"`
}

type AccessPoint struct {
	Name                string `json:"name"`
	IsLocal             bool   `json:"is_local"`
	Vdom                string `json:"vdom"`
	Eos                 bool   `json:"eos"`
	Serial              string `json:"serial"`
	ApProfile           string `json:"ap_profile"`
	BleProfile          string `json:"ble_profile"`
	State               string `json:"state"`
	ConnectingFrom      string `json:"connecting_from"`
	ConnectingInterface string `json:"connecting_interface"`
	Status              string `json:"status"`
	WtpID               string `json:"wtp_id"`
	Clients             int    `json:"clients"`
}

func (c *Client) GetManagedAp(ctx context.Context) (*ManagedApResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/monitor/wifi/managed_ap", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res ManagedApResponse
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
