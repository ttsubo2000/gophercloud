package utils

import (
	"fmt"
	"strings"

	"github.com/ttsubo2000/gophercloud"
)

// Version is a supported API version, corresponding to a vN package within the appropriate service.
type Version struct {
	ID       string
	Suffix   string
	Priority int
}

var goodStatus = map[string]bool{
	"current":   true,
	"supported": true,
	"stable":    true,
}

// ChooseVersion queries the base endpoint of an API to choose the most recent non-experimental alternative from a service's
// published versions.
// It returns the highest-Priority Version among the alternatives that are provided, as well as its corresponding endpoint.
func ChooseVersion(client *gophercloud.ProviderClient, recognized []*Version) (*Version, string, error) {
	type linkResp struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	}

	type valueResp struct {
		ID     string     `json:"id"`
		Status string     `json:"status"`
		Links  []linkResp `json:"links"`
	}

	type versionsResp struct {
		Values []valueResp `json:"values"`
	}

	type response struct {
		Versions versionsResp `json:"versions"`
	}

	normalize := func(endpoint string) string {
		if !strings.HasSuffix(endpoint, "/") {
			return endpoint + "/"
		}
		return endpoint
	}
	identityEndpoint := normalize(client.IdentityEndpoint)

	// If a full endpoint is specified, check version suffixes for a match first.
	for _, v := range recognized {
		if strings.HasSuffix(identityEndpoint, v.Suffix) {
			return v, identityEndpoint, nil
		}
	}

	var resp response
	_, err := client.Request("GET", client.IdentityBase, gophercloud.RequestOpts{
		JSONResponse: &resp,
		OkCodes:      []int{200, 300},
	})

	if err != nil {
		return nil, "", err
	}

	byID := make(map[string]*Version)
	for _, version := range recognized {
		byID[version.ID] = version
	}

	var highest *Version
	var endpoint string

	for _, value := range resp.Versions.Values {
		href := ""
		for _, link := range value.Links {
			if link.Rel == "self" {
				href = normalize(link.Href)
			}
		}

		if matching, ok := byID[value.ID]; ok {
			// Prefer a version that exactly matches the provided endpoint.
			if href == identityEndpoint {
				if href == "" {
					return nil, "", fmt.Errorf("Endpoint missing in version %s response from %s", value.ID, client.IdentityBase)
				}
				return matching, href, nil
			}

			// Otherwise, find the highest-priority version with a whitelisted status.
			if goodStatus[strings.ToLower(value.Status)] {
				if highest == nil || matching.Priority > highest.Priority {
					highest = matching
					endpoint = href
				}
			}
		}
	}

	if highest == nil {
		return nil, "", fmt.Errorf("No supported version available from endpoint %s", client.IdentityBase)
	}
	if endpoint == "" {
		return nil, "", fmt.Errorf("Endpoint missing in version %s response from %s", highest.ID, client.IdentityBase)
	}

	return highest, endpoint, nil
}
