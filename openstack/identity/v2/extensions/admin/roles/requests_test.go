package roles

import (
	"testing"

	"github.com/ttsubo2000/gophercloud/pagination"
	th "github.com/ttsubo2000/gophercloud/testhelper"
	"github.com/ttsubo2000/gophercloud/testhelper/client"
)

func TestRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListRoleResponse(t)

	count := 0

	err := List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractRoles(page)
		if err != nil {
			t.Errorf("Failed to extract users: %v", err)
			return false, err
		}

		expected := []Role{
			Role{
				ID:          "123",
				Name:        "compute:admin",
				Description: "Nova Administrator",
			},
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, count)
}

func TestAddUserRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockAddUserRoleResponse(t)

	err := AddUserRole(client.ServiceClient(), "{tenant_id}", "{user_id}", "{role_id}").ExtractErr()

	th.AssertNoErr(t, err)
}

func TestDeleteUserRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteUserRoleResponse(t)

	err := DeleteUserRole(client.ServiceClient(), "{tenant_id}", "{user_id}", "{role_id}").ExtractErr()

	th.AssertNoErr(t, err)
}
