package keypairs

import (
	"testing"

	os "github.com/ttsubo2000/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/ttsubo2000/gophercloud/pagination"
	th "github.com/ttsubo2000/gophercloud/testhelper"
	"github.com/ttsubo2000/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleListSuccessfully(t)

	count := 0
	err := List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractKeyPairs(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, os.ExpectedKeyPairSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleCreateSuccessfully(t)

	actual, err := Create(client.ServiceClient(), os.CreateOpts{
		Name: "createdkey",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &os.CreatedKeyPair, actual)
}

func TestImport(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleImportSuccessfully(t)

	actual, err := Create(client.ServiceClient(), os.CreateOpts{
		Name:      "importedkey",
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDx8nkQv/zgGgB4rMYmIf+6A4l6Rr+o/6lHBQdW5aYd44bd8JttDCE/F/pNRr0lRE+PiqSPO8nDPHw0010JeMH9gYgnnFlyY3/OcJ02RhIPyyxYpv9FhY+2YiUkpwFOcLImyrxEsYXpD/0d3ac30bNH6Sw9JD9UZHYcpSxsIbECHw== Generated by Nova",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &os.ImportedKeyPair, actual)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleGetSuccessfully(t)

	actual, err := Get(client.ServiceClient(), "firstkey").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &os.FirstKeyPair, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleDeleteSuccessfully(t)

	err := Delete(client.ServiceClient(), "deletedkey").ExtractErr()
	th.AssertNoErr(t, err)
}
