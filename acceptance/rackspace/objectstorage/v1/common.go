// +build acceptance ttsubo2000 objectstorage v1

package v1

import (
	"os"
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

func createClient(t *testing.T, cdn bool) (*gophercloud.ServiceClient, error) {
	region := os.Getenv("RS_REGION")
	if region == "" {
		t.Fatal("Please provide a ttsubo2000 region as RS_REGION")
	}

	ao := ttsubo2000AuthOptions(t)

	provider, err := ttsubo2000.NewClient(ao.IdentityEndpoint)
	th.AssertNoErr(t, err)

	err = ttsubo2000.Authenticate(provider, ao)
	th.AssertNoErr(t, err)

	if cdn {
		return ttsubo2000.NewObjectCDNV1(provider, gophercloud.EndpointOpts{
			Region: region,
		})
	}

	return ttsubo2000.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: region,
	})
}
