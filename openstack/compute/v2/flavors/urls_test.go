package flavors

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient(), "foo")
	expected := endpoint + "flavors/foo"
	th.CheckEquals(t, expected, actual)
}

func TestListURL(t *testing.T) {
	actual := listURL(endpointClient())
	expected := endpoint + "flavors/detail"
	th.CheckEquals(t, expected, actual)
}
