package fortios

import (
	"context"
	"fmt"
	"net/http"
)

type ManagedSwitchResponse struct {
	HTTPMethod string   `json:"http_method"`
	Results    []Switch `json:"results"`
	Vdom       string   `json:"vdom"`
	Path       string   `json:"path"`
	Name       string   `json:"name"`
	Action     string   `json:"action"`
	Status     string   `json:"status"`
	Serial     string   `json:"serial"`
	Version    string   `json:"version"`
	Build      int      `json:"build"`
}
type Transceiver struct {
	VendorPartNumber string `json:"vendor_part_number"`
	Vendor           string `json:"vendor"`
}
type IgmpSnoopingGroup struct {
	GroupCount int `json:"group_count"`
}
type DhcpSnooping struct {
	Untrusted int `json:"untrusted"`
}
type Port struct {
	Interface         string            `json:"interface"`
	FortilinkPort     bool              `json:"fortilink_port"`
	Vlan              string            `json:"vlan"`
	FgtPeerPortName   string            `json:"fgt_peer_port_name"`
	FgtPeerDeviceName string            `json:"fgt_peer_device_name"`
	IslPeerDeviceName string            `json:"isl_peer_device_name"`
	IslPeerPortName   string            `json:"isl_peer_port_name"`
	IslPeerTrunkName  string            `json:"isl_peer_trunk_name"`
	MclagIcl          bool              `json:"mclag_icl"`
	Mclag             bool              `json:"mclag"`
	Status            string            `json:"status"`
	Duplex            string            `json:"duplex"`
	Speed             int               `json:"speed"`
	PoeCapable        bool              `json:"poe_capable"`
	PortPower         float64           `json:"port_power"`
	PowerStatus       int               `json:"power_status"`
	Transceiver       Transceiver       `json:"transceiver,omitempty"`
	StpStatus         string            `json:"stp_status,omitempty"`
	IgmpSnoopingGroup IgmpSnoopingGroup `json:"igmp_snooping_group"`
	DhcpSnooping      DhcpSnooping      `json:"dhcp_snooping"`
}
type Switch struct {
	Status                string `json:"status"`
	OsVersion             string `json:"os_version"`
	ConnectingFrom        string `json:"connecting_from,omitempty"`
	JoinTime              string `json:"join_time"`
	ImageDownloadProgress int    `json:"image_download_progress"`
	Name                  string `json:"name"`
	Serial                string `json:"serial"`
	FgtPeerIntfName       string `json:"fgt_peer_intf_name"`
	State                 string `json:"state"`
	Ports                 []Port `json:"ports"`
	MaxPoeBudget          int    `json:"max_poe_budget"`
	IgmpSnoopingSupported bool   `json:"igmp_snooping_supported"`
	DhcpSnoopingSupported bool   `json:"dhcp_snooping_supported"`
	McLagSupported        bool   `json:"mc_lag_supported"`
	LedBlinkSupported     bool   `json:"led_blink_supported"`
	Type                  string `json:"type"`
	FaceplateXML          string `json:"faceplate_xml"`
	Vdom                  string `json:"vdom"`
	IsL3                  bool   `json:"is_l3"`
}

func (c *Client) GetManagedSwitch(ctx context.Context) (*ManagedSwitchResponse, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/monitor/switch-controller/managed-switch/status", c.baseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	var res ManagedSwitchResponse
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
