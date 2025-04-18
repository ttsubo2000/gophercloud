package portsbinding

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint, ResourceBase: endpoint + "v2.0/"}
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient(), "foo")
	expected := endpoint + "v2.0/ports/foo"
	th.AssertEquals(t, expected, actual)
}

func TestCreateURL(t *testing.T) {
	actual := createURL(endpointClient())
	expected := endpoint + "v2.0/ports"
	th.AssertEquals(t, expected, actual)
}

func TestUpdateURL(t *testing.T) {
	actual := updateURL(endpointClient(), "foo")
	expected := endpoint + "v2.0/ports/foo"
	th.AssertEquals(t, expected, actual)
}
