package v2

import (
	"os"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

var Client *gophercloud.ServiceClient

func NewClient() (*gophercloud.ServiceClient, error) {
	opts, err := ttsubo2000.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	provider, err := ttsubo2000.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	return ttsubo2000.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Name:   "cloudNetworks",
		Region: os.Getenv("RS_REGION"),
	})
}

func Setup(t *testing.T) {
	client, err := NewClient()
	th.AssertNoErr(t, err)
	Client = client
}

func Teardown() {
	Client = nil
}
