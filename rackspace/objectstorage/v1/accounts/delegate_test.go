package accounts

import (
	"testing"

	os "github.com/ttsubo2000/gophercloud/openstack/objectstorage/v1/accounts"
	th "github.com/ttsubo2000/gophercloud/testhelper"
	fake "github.com/ttsubo2000/gophercloud/testhelper/client"
)

func TestUpdateAccounts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	os.HandleUpdateAccountSuccessfully(t)

	options := &UpdateOpts{Metadata: map[string]string{"gophercloud-test": "accounts"}}
	res := Update(fake.ServiceClient(), options)
	th.CheckNoErr(t, res.Err)
}

func TestGetAccounts(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	os.HandleGetAccountSuccessfully(t)

	expected := map[string]string{"Foo": "bar"}
	actual, err := Get(fake.ServiceClient()).ExtractMetadata()
	th.CheckNoErr(t, err)
	th.CheckDeepEquals(t, expected, actual)
}
