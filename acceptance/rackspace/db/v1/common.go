// +build acceptance db ttsubo2000

package v1

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/db/v1/instances"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func newClient(t *testing.T) *gophercloud.ServiceClient {
	opts, err := ttsubo2000.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)
	opts = tools.OnlyRS(opts)

	client, err := ttsubo2000.AuthenticatedClient(opts)
	th.AssertNoErr(t, err)

	c, err := ttsubo2000.NewDBV1(client, gophercloud.EndpointOpts{
		Region: "IAD",
	})
	th.AssertNoErr(t, err)

	return c
}

type context struct {
	test          *testing.T
	client        *gophercloud.ServiceClient
	instanceID    string
	DBIDs         []string
	replicaID     string
	backupID      string
	configGroupID string
	users         []string
}

func newContext(t *testing.T) context {
	return context{
		test:   t,
		client: newClient(t),
	}
}

func (c context) Logf(msg string, args ...interface{}) {
	if len(args) > 0 {
		c.test.Logf(msg, args...)
	} else {
		c.test.Log(msg)
	}
}

func (c context) AssertNoErr(err error) {
	th.AssertNoErr(c.test, err)
}

func (c context) WaitUntilActive(id string) {
	err := gophercloud.WaitFor(60, func() (bool, error) {
		inst, err := instances.Get(c.client, id).Extract()
		if err != nil {
			return false, err
		}
		if inst.Status == "ACTIVE" {
			return true, nil
		}
		return false, nil
	})
	c.AssertNoErr(err)
}
