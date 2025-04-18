package floatingips

import (
	"fmt"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID                string `q:"id"`
	FloatingNetworkID string `q:"floating_network_id"`
	PortID            string `q:"port_id"`
	FixedIP           string `q:"fixed_ip_address"`
	FloatingIP        string `q:"floating_ip_address"`
	TenantID          string `q:"tenant_id"`
	Limit             int    `q:"limit"`
	Marker            string `q:"marker"`
	SortKey           string `q:"sort_key"`
	SortDir           string `q:"sort_dir"`
}

// List returns a Pager which allows you to iterate over a collection of
// floating IP resources. It accepts a ListOpts struct, which allows you to
// filter and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := gophercloud.BuildQueryString(&opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}
	u := rootURL(c) + q.String()
	return pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return FloatingIPPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOpts contains all the values needed to create a new floating IP
// resource. The only required fields are FloatingNetworkID and PortID which
// refer to the external network and internal port respectively.
type CreateOpts struct {
	FloatingNetworkID string
	FloatingIP        string
	PortID            string
	FixedIP           string
	TenantID          string
}

var (
	errFloatingNetworkIDRequired = fmt.Errorf("A NetworkID is required")
)

// Create accepts a CreateOpts struct and uses the values provided to create a
// new floating IP resource. You can create floating IPs on external networks
// only. If you provide a FloatingNetworkID which refers to a network that is
// not external (i.e. its `router:external' attribute is False), the operation
// will fail and return a 400 error.
//
// If you do not specify a FloatingIP address value, the operation will
// automatically allocate an available address for the new resource. If you do
// choose to specify one, it must fall within the subnet range for the external
// network - otherwise the operation returns a 400 error. If the FloatingIP
// address is already in use, the operation returns a 409 error code.
//
// You can associate the new resource with an internal port by using the PortID
// field. If you specify a PortID that is not valid, the operation will fail and
// return 404 error code.
//
// You must also configure an IP address for the port associated with the PortID
// you have provided - this is what the FixedIP refers to: an IP fixed to a port.
// Because a port might be associated with multiple IP addresses, you can use
// the FixedIP field to associate a particular IP address rather than have the
// API assume for you. If you specify an IP address that is not valid, the
// operation will fail and return a 400 error code. If the PortID and FixedIP
// are already associated with another resource, the operation will fail and
// returns a 409 error code.
func Create(c *gophercloud.ServiceClient, opts CreateOpts) CreateResult {
	var res CreateResult

	// Validate
	if opts.FloatingNetworkID == "" {
		res.Err = errFloatingNetworkIDRequired
		return res
	}

	// Define structures
	type floatingIP struct {
		FloatingNetworkID string `json:"floating_network_id"`
		FloatingIP        string `json:"floating_ip_address,omitempty"`
		PortID            string `json:"port_id,omitempty"`
		FixedIP           string `json:"fixed_ip_address,omitempty"`
		TenantID          string `json:"tenant_id,omitempty"`
	}
	type request struct {
		FloatingIP floatingIP `json:"floatingip"`
	}

	// Populate request body
	reqBody := request{FloatingIP: floatingIP{
		FloatingNetworkID: opts.FloatingNetworkID,
		FloatingIP:        opts.FloatingIP,
		PortID:            opts.PortID,
		FixedIP:           opts.FixedIP,
		TenantID:          opts.TenantID,
	}}

	_, res.Err = c.Post(rootURL(c), reqBody, &res.Body, nil)
	return res
}

// Get retrieves a particular floating IP resource based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = c.Get(resourceURL(c, id), &res.Body, nil)
	return res
}

// UpdateOpts contains the values used when updating a floating IP resource. The
// only value that can be updated is which internal port the floating IP is
// linked to. To associate the floating IP with a new internal port, provide its
// ID. To disassociate the floating IP from all ports, provide an empty string.
type UpdateOpts struct {
	PortID string
}

// Update allows floating IP resources to be updated. Currently, the only way to
// "update" a floating IP is to associate it with a new internal port, or
// disassociated it from all ports. See UpdateOpts for instructions of how to
// do this.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) UpdateResult {
	type floatingIP struct {
		PortID *string `json:"port_id"`
	}

	type request struct {
		FloatingIP floatingIP `json:"floatingip"`
	}

	var portID *string
	if opts.PortID == "" {
		portID = nil
	} else {
		portID = &opts.PortID
	}

	reqBody := request{FloatingIP: floatingIP{PortID: portID}}

	// Send request to API
	var res UpdateResult
	_, res.Err = c.Put(resourceURL(c, id), reqBody, &res.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})

	return res
}

// Delete will permanently delete a particular floating IP resource. Please
// ensure this is what you want - you can also disassociate the IP from existing
// internal ports.
func Delete(c *gophercloud.ServiceClient, id string) DeleteResult {
	var res DeleteResult
	_, res.Err = c.Delete(resourceURL(c, id), nil)
	return res
}
