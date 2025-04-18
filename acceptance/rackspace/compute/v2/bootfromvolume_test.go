// +build acceptance

package v2

import (
	"testing"

	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	osBFV "github.com/ttsubo2000/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/compute/v2/bootfromvolume"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/compute/v2/servers"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestBootFromVolume(t *testing.T) {
	client, err := newClient()
	th.AssertNoErr(t, err)

	if testing.Short() {
		t.Skip("Skipping test that requires server creation in short mode.")
	}

	options, err := optionsFromEnv()
	th.AssertNoErr(t, err)

	name := tools.RandomString("Gophercloud-", 8)
	t.Logf("Creating server [%s].", name)

	bd := []osBFV.BlockDevice{
		osBFV.BlockDevice{
			UUID:       options.imageID,
			SourceType: osBFV.Image,
			VolumeSize: 10,
		},
	}

	server, err := bootfromvolume.Create(client, servers.CreateOpts{
		Name:        name,
		FlavorRef:   "performance1-1",
		BlockDevice: bd,
	}).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created server: %+v\n", server)
	defer deleteServer(t, client, server)

	getServer(t, client, server)

	listServers(t, client)
}
