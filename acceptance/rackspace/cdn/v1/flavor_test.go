// +build acceptance

package v1

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	os "github.com/ttsubo2000/gophercloud/openstack/cdn/v1/flavors"
	"github.com/ttsubo2000/gophercloud/pagination"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/cdn/v1/flavors"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestFlavor(t *testing.T) {
	client := newClient(t)

	t.Log("Listing Flavors")
	id := testFlavorsList(t, client)

	t.Log("Retrieving Flavor")
	testFlavorGet(t, client, id)
}

func testFlavorsList(t *testing.T, client *gophercloud.ServiceClient) string {
	var id string
	err := flavors.List(client).EachPage(func(page pagination.Page) (bool, error) {
		flavorList, err := os.ExtractFlavors(page)
		th.AssertNoErr(t, err)

		for _, flavor := range flavorList {
			t.Logf("Listing flavor: ID [%s] Providers [%+v]", flavor.ID, flavor.Providers)
			id = flavor.ID
		}

		return true, nil
	})

	th.AssertNoErr(t, err)
	return id
}

func testFlavorGet(t *testing.T, client *gophercloud.ServiceClient, id string) {
	flavor, err := flavors.Get(client, id).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Retrieved Flavor: %+v", *flavor)
}
