// +build acceptance

package v1

import (
	"testing"

	"github.com/ttsubo2000/gophercloud/ttsubo2000/orchestration/v1/buildinfo"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestBuildInfo(t *testing.T) {
	// Create a provider client for making the HTTP requests.
	// See common.go in this directory for more information.
	client := newClient(t)

	bi, err := buildinfo.Get(client).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved build info: %+v\n", bi)
}
