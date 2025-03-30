// +build acceptance lbs

package v1

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func newProvider() (*gophercloud.ProviderClient, error) {
	opts, err := ttsubo2000.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}
	opts = tools.OnlyRS(opts)

	return ttsubo2000.AuthenticatedClient(opts)
}

func newClient() (*gophercloud.ServiceClient, error) {
	provider, err := newProvider()
	if err != nil {
		return nil, err
	}

	return ttsubo2000.NewLBV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RS_REGION"),
	})
}

func newComputeClient() (*gophercloud.ServiceClient, error) {
	provider, err := newProvider()
	if err != nil {
		return nil, err
	}

	return ttsubo2000.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RS_REGION"),
	})
}

func setup(t *testing.T) *gophercloud.ServiceClient {
	client, err := newClient()
	th.AssertNoErr(t, err)

	return client
}

func intsToStr(ids []int) string {
	strIDs := []string{}
	for _, id := range ids {
		strIDs = append(strIDs, strconv.Itoa(id))
	}
	return strings.Join(strIDs, ", ")
}
