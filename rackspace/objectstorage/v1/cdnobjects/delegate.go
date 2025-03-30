package cdnobjects

import (
	"github.com/ttsubo2000/gophercloud"
	os "github.com/ttsubo2000/gophercloud/openstack/objectstorage/v1/objects"
)

// Delete is a function that deletes an object from the CDN.
func Delete(c *gophercloud.ServiceClient, containerName, objectName string, opts os.DeleteOptsBuilder) os.DeleteResult {
	return os.Delete(c, containerName, objectName, nil)
}
