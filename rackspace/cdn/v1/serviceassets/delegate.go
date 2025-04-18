package serviceassets

import (
	"github.com/ttsubo2000/gophercloud"

	os "github.com/ttsubo2000/gophercloud/openstack/cdn/v1/serviceassets"
)

// Delete accepts a unique ID and deletes the CDN service asset associated with
// it.
func Delete(c *gophercloud.ServiceClient, id string, opts os.DeleteOptsBuilder) os.DeleteResult {
	return os.Delete(c, id, opts)
}
