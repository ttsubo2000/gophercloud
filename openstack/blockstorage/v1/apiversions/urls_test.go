package apiversions

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"
const endpoint2 = "http://localhost:57909/v1/3a02ee0b5cf14816b41b17e851d29a94"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func endpointClient2() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint2}
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient(), "v1")
	expected := endpoint + "v1/"
	th.AssertEquals(t, expected, actual)
}

func TestListURL(t *testing.T) {
	actual := listURL(endpointClient2())
	expected := endpoint
	th.AssertEquals(t, expected, actual)
}
