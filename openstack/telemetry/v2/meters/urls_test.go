package meters

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"
const meter = "cpu"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestListURL(t *testing.T) {
	actual := listURL(endpointClient())
	expected := endpoint + "v2/meters"
	th.CheckEquals(t, expected, actual)
}

func TestShowURL(t *testing.T) {
	actual := showURL(endpointClient(), meter)
	expected := endpoint + "v2/meters/" + meter
	th.CheckEquals(t, expected, actual)
}

func TestCreateURL(t *testing.T) {
	actual := createURL(endpointClient(), meter)
	expected := endpoint + "v2/meters/" + meter
	th.CheckEquals(t, expected, actual)
}

func TestStatisticsURL(t *testing.T) {
	actual := statisticsURL(endpointClient(), meter)
	expected := endpoint + "v2/meters/" + meter + "/statistics"
	th.CheckEquals(t, expected, actual)
}
