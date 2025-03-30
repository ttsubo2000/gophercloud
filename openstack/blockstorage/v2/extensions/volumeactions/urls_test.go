package volumeactions

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestAttachURL(t *testing.T) {
	actual := attachURL(endpointClient(), "foo")
	expected := endpoint + "volumes/foo/action"
	th.AssertEquals(t, expected, actual)
}

func TestDettachURL(t *testing.T) {
	actual := detachURL(endpointClient(), "foo")
	expected := endpoint + "volumes/foo/action"
	th.AssertEquals(t, expected, actual)
}

func TestReserveURL(t *testing.T) {
	actual := reserveURL(endpointClient(), "foo")
	expected := endpoint + "volumes/foo/action"
	th.AssertEquals(t, expected, actual)
}

func TestUnreserveURL(t *testing.T) {
	actual := unreserveURL(endpointClient(), "foo")
	expected := endpoint + "volumes/foo/action"
	th.AssertEquals(t, expected, actual)
}

func TestInitializeConnectionURL(t *testing.T) {
	actual := initializeConnectionURL(endpointClient(), "foo")
	expected := endpoint + "volumes/foo/action"
	th.AssertEquals(t, expected, actual)
}

func TestTeminateConnectionURL(t *testing.T) {
	actual := teminateConnectionURL(endpointClient(), "foo")
	expected := endpoint + "volumes/foo/action"
	th.AssertEquals(t, expected, actual)
}
