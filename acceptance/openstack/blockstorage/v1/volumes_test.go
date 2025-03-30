// +build acceptance blockstorage

package v1

import (
	"os"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/openstack"
	"github.com/ttsubo2000/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/ttsubo2000/gophercloud/pagination"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func newClient(t *testing.T) (*gophercloud.ServiceClient, error) {
	ao, err := openstack.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := openstack.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	return openstack.NewBlockStorageV1(client, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
}

func TestVolumes(t *testing.T) {
	client, err := newClient(t)
	th.AssertNoErr(t, err)

	cv, err := volumes.Create(client, &volumes.CreateOpts{
		Size: 1,
		Name: "gophercloud-test-volume",
	}).Extract()
	th.AssertNoErr(t, err)
	defer func() {
		err = volumes.WaitForStatus(client, cv.ID, "available", 60)
		th.AssertNoErr(t, err)
		err = volumes.Delete(client, cv.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	_, err = volumes.Update(client, cv.ID, &volumes.UpdateOpts{
		Name: "gophercloud-updated-volume",
	}).Extract()
	th.AssertNoErr(t, err)

	v, err := volumes.Get(client, cv.ID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Got volume: %+v\n", v)

	if v.Name != "gophercloud-updated-volume" {
		t.Errorf("Unable to update volume: Expected name: gophercloud-updated-volume\nActual name: %s", v.Name)
	}

	err = volumes.List(client, &volumes.ListOpts{Name: "gophercloud-updated-volume"}).EachPage(func(page pagination.Page) (bool, error) {
		vols, err := volumes.ExtractVolumes(page)
		th.CheckEquals(t, 1, len(vols))
		return true, err
	})
	th.AssertNoErr(t, err)
}
