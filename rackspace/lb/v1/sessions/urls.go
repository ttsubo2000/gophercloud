package sessions

import (
	"strconv"

	"github.com/ttsubo2000/gophercloud"
)

const (
	path   = "loadbalancers"
	spPath = "sessionpersistence"
)

func rootURL(c *gophercloud.ServiceClient, id int) string {
	return c.ServiceURL(path, strconv.Itoa(id), spPath)
}
