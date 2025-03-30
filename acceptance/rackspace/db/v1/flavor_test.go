// +build acceptance db ttsubo2000

package v1

import (
	os "github.com/ttsubo2000/gophercloud/openstack/db/v1/flavors"
	"github.com/ttsubo2000/gophercloud/pagination"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/db/v1/flavors"
)

func (c context) listFlavors() {
	c.Logf("Listing flavors")

	err := flavors.List(c.client).EachPage(func(page pagination.Page) (bool, error) {
		flavorList, err := os.ExtractFlavors(page)
		c.AssertNoErr(err)

		for _, f := range flavorList {
			c.Logf("Flavor: ID [%s] Name [%s] RAM [%d]", f.ID, f.Name, f.RAM)
		}

		return true, nil
	})

	c.AssertNoErr(err)
}

func (c context) getFlavor() {
	flavor, err := flavors.Get(c.client, "1").Extract()
	c.Logf("Getting flavor %s", flavor.ID)
	c.AssertNoErr(err)
}
