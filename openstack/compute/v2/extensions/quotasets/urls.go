package quotasets

import "github.com/ttsubo2000/gophercloud"

const resourcePath = "os-quota-sets"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func getURL(c *gophercloud.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID)
}
