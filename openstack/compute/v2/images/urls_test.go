package images

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
	expected := endpoint + "images/foo"
	th.CheckEquals(t, expected, actual)
}

func TestListDetailURL(t *testing.T) {
	actual := listDetailURL(endpointClient())
	expected := endpoint + "images/detail"
	th.CheckEquals(t, expected, actual)
}
