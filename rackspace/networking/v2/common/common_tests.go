package common

import (
	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}
