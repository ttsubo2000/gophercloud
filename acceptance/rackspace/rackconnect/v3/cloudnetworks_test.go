// +build acceptance

package v3

import (
	"fmt"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/rackconnect/v3/cloudnetworks"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestCloudNetworks(t *testing.T) {
	c := newClient(t)
	cnID := testListNetworks(t, c)
	testGetNetworks(t, c, cnID)
}

func testListNetworks(t *testing.T, c *gophercloud.ServiceClient) string {
	allPages, err := cloudnetworks.List(c).AllPages()
	th.AssertNoErr(t, err)
	allcn, err := cloudnetworks.ExtractCloudNetworks(allPages)
	fmt.Printf("Listing all cloud networks: %+v\n\n", allcn)
	var cnID string
	if len(allcn) > 0 {
		cnID = allcn[0].ID
	}
	return cnID
}

func testGetNetworks(t *testing.T, c *gophercloud.ServiceClient, id string) {
	cn, err := cloudnetworks.Get(c, id).Extract()
	th.AssertNoErr(t, err)
	fmt.Printf("Retrieved cloud network: %+v\n\n", cn)
}
