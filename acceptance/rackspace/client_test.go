// +build acceptance

package ttsubo2000

import (
	"testing"

	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestAuthenticatedClient(t *testing.T) {
	// Obtain credentials from the environment.
	ao, err := ttsubo2000.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := ttsubo2000.AuthenticatedClient(tools.OnlyRS(ao))
	if err != nil {
		t.Fatalf("Unable to authenticate: %v", err)
	}

	if client.TokenID == "" {
		t.Errorf("No token ID assigned to the client")
	}

	t.Logf("Client successfully acquired a token: %v", client.TokenID)
}
