package virtualinterfaces

import "github.com/ttsubo2000/gophercloud"

func listURL(c *gophercloud.ServiceClient, instanceID string) string {
	return c.ServiceURL("servers", instanceID, "os-virtual-interfacesv2")
}

func createURL(c *gophercloud.ServiceClient, instanceID string) string {
	return c.ServiceURL("servers", instanceID, "os-virtual-interfacesv2")
}

func deleteURL(c *gophercloud.ServiceClient, instanceID, interfaceID string) string {
	return c.ServiceURL("servers", instanceID, "os-virtual-interfacesv2", interfaceID)
}
