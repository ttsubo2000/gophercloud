package webhooks

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestListURL(t *testing.T) {
	actual := listURL(endpointClient(), "123", "456")
	expected := endpoint + "groups/123/policies/456/webhooks"
	th.CheckEquals(t, expected, actual)
}

func TestCreateURL(t *testing.T) {
	actual := createURL(endpointClient(), "123", "456")
	expected := endpoint + "groups/123/policies/456/webhooks"
	th.CheckEquals(t, expected, actual)
}

func TestGetURL(t *testing.T) {
	actual := getURL(endpointClient(), "123", "456", "789")
	expected := endpoint + "groups/123/policies/456/webhooks/789"
	th.CheckEquals(t, expected, actual)
}

func TestUpdateURL(t *testing.T) {
	actual := updateURL(endpointClient(), "123", "456", "789")
	expected := endpoint + "groups/123/policies/456/webhooks/789"
	th.CheckEquals(t, expected, actual)
}

func TestDeleteURL(t *testing.T) {
	actual := deleteURL(endpointClient(), "123", "456", "789")
	expected := endpoint + "groups/123/policies/456/webhooks/789"
	th.CheckEquals(t, expected, actual)
}
