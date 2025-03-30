// +build acceptance

package v1

import (
	"fmt"
	"os"
	"testing"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

var template = fmt.Sprintf(`
{
		"heat_template_version": "2013-05-23",
		"description": "Simple template to test heat commands",
		"parameters": {},
		"resources": {
				"hello_world": {
						"type":"OS::Nova::Server",
						"properties": {
								"flavor": "%s",
								"image": "%s",
								"user_data": "#!/bin/bash -xv\necho \"hello world\" &gt; /root/hello-world.txt\n"
						}
				}
		}
}
`, os.Getenv("RS_FLAVOR_ID"), os.Getenv("RS_IMAGE_ID"))

func newClient(t *testing.T) *gophercloud.ServiceClient {
	ao, err := ttsubo2000.AuthOptionsFromEnv()
	th.AssertNoErr(t, err)

	client, err := ttsubo2000.AuthenticatedClient(ao)
	th.AssertNoErr(t, err)

	c, err := ttsubo2000.NewOrchestrationV1(client, gophercloud.EndpointOpts{
		Region: os.Getenv("RS_REGION_NAME"),
	})
	th.AssertNoErr(t, err)
	return c
}
