//go:build integration
// +build integration

package fortios

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetManagedAp(t *testing.T) {
	c := NewClient(os.Getenv("FORTIOS_INTEGRATION_API_KEY"), os.Getenv("FORTIOS_INTEGRATION_FIREWALL"))

	ctx := context.Background()
	res, err := c.GetManagedAp(ctx)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "wifi", res.Path, "expecting correct Path")
	}
	if res != nil {
		assert.Equal(t, "GET", res.HTTPMethod, "expecting correct HTTPMethod")
	}
	if res != nil {
		assert.Equal(t, "success", res.Status, "expecting success Status")
	}
}

func TestGetWifiClient(t *testing.T) {
	c := NewClient(os.Getenv("FORTIOS_INTEGRATION_API_KEY"), os.Getenv("FORTIOS_INTEGRATION_FIREWALL"))

	ctx := context.Background()
	res, err := c.GetWifiClient(ctx)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, "wifi", res.Path, "expecting correct Path")
	}
	if res != nil {
		assert.Equal(t, "GET", res.HTTPMethod, "expecting correct HTTPMethod")
	}
	if res != nil {
		assert.Equal(t, "success", res.Status, "expecting success Status")
	}
}
