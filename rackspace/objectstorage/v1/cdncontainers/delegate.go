package cdncontainers

import (
	"github.com/ttsubo2000/gophercloud"
	os "github.com/ttsubo2000/gophercloud/openstack/objectstorage/v1/containers"
	"github.com/ttsubo2000/gophercloud/pagination"
)

// ExtractNames interprets a page of List results when just the container
// names are requested.
func ExtractNames(page pagination.Page) ([]string, error) {
	return os.ExtractNames(page)
}

// ListOpts are options for listing ttsubo2000 CDN containers.
type ListOpts struct {
	EndMarker string `q:"end_marker"`
	Format    string `q:"format"`
	Limit     int    `q:"limit"`
	Marker    string `q:"marker"`
}

// ToContainerListParams formats a ListOpts into a query string and boolean
// representing whether to list complete information for each container.
func (opts ListOpts) ToContainerListParams() (bool, string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	if err != nil {
		return false, "", err
	}
	return false, q.String(), nil
}

// List is a function that retrieves containers associated with the account as
// well as account metadata. It returns a pager which can be iterated with the
// EachPage function.
func List(c *gophercloud.ServiceClient, opts os.ListOptsBuilder) pagination.Pager {
	return os.List(c, opts)
}
