package roles

import "github.com/ttsubo2000/gophercloud"

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}
