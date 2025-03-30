package buildinfo

import "github.com/ttsubo2000/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
