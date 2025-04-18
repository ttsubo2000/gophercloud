// +build fixtures

package networks

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	th "github.com/ttsubo2000/gophercloud/testhelper"
	"github.com/ttsubo2000/gophercloud/testhelper/client"
)

// ListOutput is a sample response to a List call.
const ListOutput = `
{
    "networks": [
        {
            "bridge": "br100",
            "bridge_interface": "eth0",
            "broadcast": "10.0.0.7",
            "cidr": "10.0.0.0/29",
            "cidr_v6": null,
            "created_at": "2011-08-15 06:19:19.387525",
            "deleted": false,
            "deleted_at": null,
            "dhcp_start": "10.0.0.3",
            "dns1": null,
            "dns2": null,
            "gateway": "10.0.0.1",
            "gateway_v6": null,
            "host": "nsokolov-desktop",
            "id": "20c8acc0-f747-4d71-a389-46d078ebf047",
            "injected": false,
            "label": "mynet_0",
            "multi_host": false,
            "netmask": "255.255.255.248",
            "netmask_v6": null,
            "priority": null,
            "project_id": "1234",
            "rxtx_base": null,
            "updated_at": "2011-08-16 09:26:13.048257",
            "vlan": 100,
            "vpn_private_address": "10.0.0.2",
            "vpn_public_address": "127.0.0.1",
            "vpn_public_port": 1000
        },
        {
            "bridge": "br101",
            "bridge_interface": "eth0",
            "broadcast": "10.0.0.15",
            "cidr": "10.0.0.10/29",
            "cidr_v6": null,
            "created_at": "2011-08-15 06:19:19.885495",
            "deleted": false,
            "deleted_at": null,
            "dhcp_start": "10.0.0.11",
            "dns1": null,
            "dns2": null,
            "gateway": "10.0.0.9",
            "gateway_v6": null,
            "host": null,
            "id": "20c8acc0-f747-4d71-a389-46d078ebf000",
            "injected": false,
            "label": "mynet_1",
            "multi_host": false,
            "netmask": "255.255.255.248",
            "netmask_v6": null,
            "priority": null,
            "project_id": null,
            "rxtx_base": null,
            "updated_at": null,
            "vlan": 101,
            "vpn_private_address": "10.0.0.10",
            "vpn_public_address": null,
            "vpn_public_port": 1001
        }
    ]
}
`

// GetOutput is a sample response to a Get call.
const GetOutput = `
{
    "network": {
			"bridge": "br101",
			"bridge_interface": "eth0",
			"broadcast": "10.0.0.15",
			"cidr": "10.0.0.10/29",
			"cidr_v6": null,
			"created_at": "2011-08-15 06:19:19.885495",
			"deleted": false,
			"deleted_at": null,
			"dhcp_start": "10.0.0.11",
			"dns1": null,
			"dns2": null,
			"gateway": "10.0.0.9",
			"gateway_v6": null,
			"host": null,
			"id": "20c8acc0-f747-4d71-a389-46d078ebf000",
			"injected": false,
			"label": "mynet_1",
			"multi_host": false,
			"netmask": "255.255.255.248",
			"netmask_v6": null,
			"priority": null,
			"project_id": null,
			"rxtx_base": null,
			"updated_at": null,
			"vlan": 101,
			"vpn_private_address": "10.0.0.10",
			"vpn_public_address": null,
			"vpn_public_port": 1001
		}
}
`

// FirstNetwork is the first result in ListOutput.
var nilTime time.Time
var FirstNetwork = Network{
	Bridge:            "br100",
	BridgeInterface:   "eth0",
	Broadcast:         "10.0.0.7",
	CIDR:              "10.0.0.0/29",
	CIDRv6:            "",
	CreatedAt:         time.Date(2011, 8, 15, 6, 19, 19, 387525000, time.UTC),
	Deleted:           false,
	DeletedAt:         nilTime,
	DHCPStart:         "10.0.0.3",
	DNS1:              "",
	DNS2:              "",
	Gateway:           "10.0.0.1",
	Gatewayv6:         "",
	Host:              "nsokolov-desktop",
	ID:                "20c8acc0-f747-4d71-a389-46d078ebf047",
	Injected:          false,
	Label:             "mynet_0",
	MultiHost:         false,
	Netmask:           "255.255.255.248",
	Netmaskv6:         "",
	Priority:          0,
	ProjectID:         "1234",
	RXTXBase:          0,
	UpdatedAt:         time.Date(2011, 8, 16, 9, 26, 13, 48257000, time.UTC),
	VLAN:              100,
	VPNPrivateAddress: "10.0.0.2",
	VPNPublicAddress:  "127.0.0.1",
	VPNPublicPort:     1000,
}

// SecondNetwork is the second result in ListOutput.
var SecondNetwork = Network{
	Bridge:            "br101",
	BridgeInterface:   "eth0",
	Broadcast:         "10.0.0.15",
	CIDR:              "10.0.0.10/29",
	CIDRv6:            "",
	CreatedAt:         time.Date(2011, 8, 15, 6, 19, 19, 885495000, time.UTC),
	Deleted:           false,
	DeletedAt:         nilTime,
	DHCPStart:         "10.0.0.11",
	DNS1:              "",
	DNS2:              "",
	Gateway:           "10.0.0.9",
	Gatewayv6:         "",
	Host:              "",
	ID:                "20c8acc0-f747-4d71-a389-46d078ebf000",
	Injected:          false,
	Label:             "mynet_1",
	MultiHost:         false,
	Netmask:           "255.255.255.248",
	Netmaskv6:         "",
	Priority:          0,
	ProjectID:         "",
	RXTXBase:          0,
	UpdatedAt:         nilTime,
	VLAN:              101,
	VPNPrivateAddress: "10.0.0.10",
	VPNPublicAddress:  "",
	VPNPublicPort:     1001,
}

// ExpectedNetworkSlice is the slice of results that should be parsed
// from ListOutput, in the expected order.
var ExpectedNetworkSlice = []Network{FirstNetwork, SecondNetwork}

// HandleListSuccessfully configures the test server to respond to a List request.
func HandleListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleGetSuccessfully configures the test server to respond to a Get request
// for an existing network.
func HandleGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-networks/20c8acc0-f747-4d71-a389-46d078ebf000", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetOutput)
	})
}
