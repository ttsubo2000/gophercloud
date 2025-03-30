// +build acceptance identity

package v2

import (
	"strconv"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	os "github.com/ttsubo2000/gophercloud/openstack/identity/v2/users"
	"github.com/ttsubo2000/gophercloud/pagination"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/identity/v2/users"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestUsers(t *testing.T) {
	client := authenticatedClient(t)

	userID := createUser(t, client)

	listUsers(t, client)

	getUser(t, client, userID)

	updateUser(t, client, userID)

	resetApiKey(t, client, userID)

	deleteUser(t, client, userID)
}

func createUser(t *testing.T, client *gophercloud.ServiceClient) string {
	t.Log("Creating user")

	opts := users.CreateOpts{
		Username: tools.RandomString("user_", 5),
		Enabled:  os.Disabled,
		Email:    "new_user@foo.com",
	}

	user, err := users.Create(client, opts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created user %s", user.ID)

	return user.ID
}

func listUsers(t *testing.T, client *gophercloud.ServiceClient) {
	err := users.List(client).EachPage(func(page pagination.Page) (bool, error) {
		userList, err := os.ExtractUsers(page)
		th.AssertNoErr(t, err)

		for _, user := range userList {
			t.Logf("Listing user: ID [%s] Username [%s] Email [%s] Enabled? [%s]",
				user.ID, user.Username, user.Email, strconv.FormatBool(user.Enabled))
		}

		return true, nil
	})

	th.AssertNoErr(t, err)
}

func getUser(t *testing.T, client *gophercloud.ServiceClient, userID string) {
	_, err := users.Get(client, userID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Getting user %s", userID)
}

func updateUser(t *testing.T, client *gophercloud.ServiceClient, userID string) {
	opts := users.UpdateOpts{Username: tools.RandomString("new_name", 5), Email: "new@foo.com"}
	user, err := users.Update(client, userID, opts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Updated user %s: Username [%s] Email [%s]", userID, user.Username, user.Email)
}

func deleteUser(t *testing.T, client *gophercloud.ServiceClient, userID string) {
	res := users.Delete(client, userID)
	th.AssertNoErr(t, res.Err)
	t.Logf("Deleted user %s", userID)
}

func resetApiKey(t *testing.T, client *gophercloud.ServiceClient, userID string) {
	key, err := users.ResetAPIKey(client, userID).Extract()
	th.AssertNoErr(t, err)

	if key.APIKey == "" {
		t.Fatal("failed to reset API key for user")
	}

	t.Logf("Reset API key for user %s to %s", key.Username, key.APIKey)
}
