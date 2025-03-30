package bulk

import "github.com/ttsubo2000/gophercloud"

func deleteURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint + "?bulk-delete"
}

func extractURL(c *gophercloud.ServiceClient, ext string) string {
	return c.Endpoint + "?extract-archive=" + ext
}
