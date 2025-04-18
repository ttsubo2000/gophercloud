// +build acceptance

package v1

import (
	"testing"

	"github.com/ttsubo2000/gophercloud"
	osStackEvents "github.com/ttsubo2000/gophercloud/openstack/orchestration/v1/stackevents"
	osStacks "github.com/ttsubo2000/gophercloud/openstack/orchestration/v1/stacks"
	"github.com/ttsubo2000/gophercloud/pagination"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/orchestration/v1/stackevents"
	"github.com/ttsubo2000/gophercloud/ttsubo2000/orchestration/v1/stacks"
	th "github.com/ttsubo2000/gophercloud/testhelper"
)

func TestStackEvents(t *testing.T) {
	// Create a provider client for making the HTTP requests.
	// See common.go in this directory for more information.
	client := newClient(t)

	stackName := "postman_stack_2"
	resourceName := "hello_world"
	var eventID string

	createOpts := osStacks.CreateOpts{
		Name:     stackName,
		Template: template,
		Timeout:  5,
	}
	stack, err := stacks.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created stack: %+v\n", stack)
	defer func() {
		err := stacks.Delete(client, stackName, stack.ID).ExtractErr()
		th.AssertNoErr(t, err)
		t.Logf("Deleted stack (%s)", stackName)
	}()
	err = gophercloud.WaitFor(60, func() (bool, error) {
		getStack, err := stacks.Get(client, stackName, stack.ID).Extract()
		if err != nil {
			return false, err
		}
		if getStack.Status == "CREATE_COMPLETE" {
			return true, nil
		}
		return false, nil
	})

	err = stackevents.List(client, stackName, stack.ID, nil).EachPage(func(page pagination.Page) (bool, error) {
		events, err := osStackEvents.ExtractEvents(page)
		th.AssertNoErr(t, err)
		t.Logf("listed events: %+v\n", events)
		eventID = events[0].ID
		return false, nil
	})
	th.AssertNoErr(t, err)

	err = stackevents.ListResourceEvents(client, stackName, stack.ID, resourceName, nil).EachPage(func(page pagination.Page) (bool, error) {
		resourceEvents, err := osStackEvents.ExtractResourceEvents(page)
		th.AssertNoErr(t, err)
		t.Logf("listed resource events: %+v\n", resourceEvents)
		return false, nil
	})
	th.AssertNoErr(t, err)

	event, err := stackevents.Get(client, stackName, stack.ID, resourceName, eventID).Extract()
	th.AssertNoErr(t, err)
	t.Logf("retrieved event: %+v\n", event)
}
