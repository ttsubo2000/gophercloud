// +build acceptance

package v2

import (
	"os"
	"testing"

	"github.com/ttsubo2000/gophercloud/ttsubo2000/identity/v2/tokens"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestTokenAuth(t *testing.T) {
	authedClient := createClient(t, true)
	token := authedClient.TokenID

	tenantID := os.Getenv("RS_TENANT_ID")
	if tenantID == "" {
		t.Skip("You must set RS_TENANT_ID environment variable to run this test")
	}

	authOpts := tokens.AuthOptions{}
	authOpts.TenantID = tenantID
	authOpts.TokenID = token

	_, err := tokens.Create(authedClient, authOpts).ExtractToken()
	th.AssertNoErr(t, err)
}
