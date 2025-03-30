// +build acceptance identity roles

package v2

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/openstack/identity/v2/extensions/admin/roles"
	"github.com/ttsubo2000/gophercloud/pagination"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestRoles(t *testing.T) {
	client := authenticatedClient(t)

	tenantID := findTenant(t, client)
	userID := createUser(t, client, tenantID)
	roleID := listRoles(t, client)

	addUserRole(t, client, tenantID, userID, roleID)

	deleteUserRole(t, client, tenantID, userID, roleID)

	deleteUser(t, client, userID)
}

func listRoles(t *testing.T, client *gophercloud.ServiceClient) string {
	var roleID string

	err := roles.List(client).EachPage(func(page pagination.Page) (bool, error) {
		roleList, err := roles.ExtractRoles(page)
		th.AssertNoErr(t, err)

		for _, role := range roleList {
			t.Logf("Listing role: ID [%s] Name [%s]", role.ID, role.Name)
			roleID = role.ID
		}

		return true, nil
	})

	th.AssertNoErr(t, err)

	return roleID
}

func addUserRole(t *testing.T, client *gophercloud.ServiceClient, tenantID, userID, roleID string) {
	err := roles.AddUserRole(client, tenantID, userID, roleID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Added role %s to user %s", roleID, userID)
}

func deleteUserRole(t *testing.T, client *gophercloud.ServiceClient, tenantID, userID, roleID string) {
	err := roles.DeleteUserRole(client, tenantID, userID, roleID).ExtractErr()
	th.AssertNoErr(t, err)
	t.Logf("Removed role %s from user %s", roleID, userID)
}
