// +build acceptance

package v2

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func ttsubo2000AuthOptions(t *testing.T) gophercloud.AuthOptions {
	// Obtain credentials from the environment.
	options, err := ttsubo2000.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)
	options = tools.OnlyRS(options)

	if options.Username == "" {
		t.Fatal("Please provide a ttsubo2000 username as RS_USERNAME.")
	}
	if options.APIKey == "" {
		t.Fatal("Please provide a ttsubo2000 API key as RS_API_KEY.")
	}

	return options
}

func createClient(t *testing.T, auth bool) *gophercloud.ServiceClient {
	ao := ttsubo2000AuthOptions(t)

	provider, err := ttsubo2000.NewClient(ao.IdentityEndpoint)
	th.AssertNoErr(t, err)

	if auth {
		err = ttsubo2000.Authenticate(provider, ao)
		th.AssertNoErr(t, err)
	}

	return ttsubo2000.NewIdentityV2(provider)
}

func unauthenticatedClient(t *testing.T) *gophercloud.ServiceClient {
	return createClient(t, false)
}

func authenticatedClient(t *testing.T) *gophercloud.ServiceClient {
	return createClient(t, true)
}
