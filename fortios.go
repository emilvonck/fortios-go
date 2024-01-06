package fortios

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(appkey string, baseurl string) *Client {
	return &Client{
		baseURL: baseurl,
		apiKey:  appkey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

type errorResponse struct {
	Status     string `json:"status"`
	HTTPStatus int    `json:"http_status"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Serial     string `json:"serial"`
	Version    string `json:"version"`
	Build      int    `json:"build"`
	HTTPMethod string `json:"http_method"`
}

type successResponse struct {
	Results interface{} `json:"data"`
}

// Content-type and body should be already added to req
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Status)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}
	var parsedJsonReturn map[string]interface{}
	json.Unmarshal(responseData, &parsedJsonReturn)
	modifiedJsonReturn := map[string]interface{}{
		"data": parsedJsonReturn,
	}

	newReturn, err := json.Marshal(modifiedJsonReturn)
	if err != nil {
		return err
	}

	// Unmarshall and populate v
	fullResponse := successResponse{
		Results: v,
	}
	if err = json.Unmarshal(newReturn, &fullResponse); err != nil {
		return err
	}

	return nil
}
