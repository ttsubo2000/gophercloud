package startstop

import (
	"testing"

	th "github.com/ttsubo2000/gophercloud/testhelper"
	"github.com/ttsubo2000/gophercloud/testhelper/client"
)

const serverID = "{serverId}"

func TestStart(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStartServerResponse(t, serverID)

	err := Start(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestStop(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStopServerResponse(t, serverID)

	err := Stop(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
