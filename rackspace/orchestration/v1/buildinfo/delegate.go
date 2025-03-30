package buildinfo

import (
	"github.com/ttsubo2000/gophercloud"
	os "github.com/ttsubo2000/gophercloud/openstack/orchestration/v1/buildinfo"
)

// Get retreives build info data for the Heat deployment.
func Get(c *gophercloud.ServiceClient) os.GetResult {
	return os.Get(c)
}
