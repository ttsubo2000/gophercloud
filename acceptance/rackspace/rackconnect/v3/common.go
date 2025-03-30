// +build acceptance

package v3

import (
	"os"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func newClient(t *testing.T) *gophercloud.ServiceClient {
	ao, err := ttsubo2000.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := ttsubo2000.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	c, err := ttsubo2000.NewRackConnectV3(client, gophercloud.EndpointOpts{
		Region: os.Getenv("RS_REGION_NAME"),
	})
	th.AssertNoErr(t, err)
	return c
}
