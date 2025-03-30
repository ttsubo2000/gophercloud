package datastores

import (
	"github.com/ttsubo2000/gophercloud"
	os "github.com/ttsubo2000/gophercloud/openstack/db/v1/datastores"
	"github.com/ttsubo2000/gophercloud/pagination"
)

// List will list all available flavors.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return os.List(client)
}

// Get retrieves the details for a particular flavor.
func Get(client *gophercloud.ServiceClient, flavorID string) os.GetResult {
	return os.Get(client, flavorID)
}

// ListVersions will list all of the available versions for a specified
// datastore type.
func ListVersions(client *gophercloud.ServiceClient, datastoreID string) pagination.Pager {
	return os.ListVersions(client, datastoreID)
}

// GetVersion will retrieve the details of a specified datastore version.
func GetVersion(client *gophercloud.ServiceClient, datastoreID, versionID string) os.GetVersionResult {
	return os.GetVersion(client, datastoreID, versionID)
}
