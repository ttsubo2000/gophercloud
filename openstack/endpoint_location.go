package openstack

import (
	"fmt"

	"github.com/ttsubo2000/gophercloud"
	tokens2 "github.com/ttsubo2000/gophercloud/openstack/identity/v2/tokens"
	tokens3 "github.com/ttsubo2000/gophercloud/openstack/identity/v3/tokens"
)

// V2EndpointURL discovers the endpoint URL for a specific service from a ServiceCatalog acquired
// during the v2 identity service. The specified EndpointOpts are used to identify a unique,
// unambiguous endpoint to return. It's an error both when multiple endpoints match the provided
// criteria and when none do. The minimum that can be specified is a Type, but you will also often
// need to specify a Name and/or a Region depending on what's available on your OpenStack
// deployment.
func V2EndpointURL(catalog *tokens2.ServiceCatalog, opts gophercloud.EndpointOpts) (string, error) {
	// Extract Endpoints from the catalog entries that match the requested Type, Name if provided, and Region if provided.
	var endpoints = make([]tokens2.Endpoint, 0, 1)
	for _, entry := range catalog.Entries {
		if (entry.Type == opts.Type) && (opts.Name == "" || entry.Name == opts.Name) {
			for _, endpoint := range entry.Endpoints {
				if opts.Region == "" || endpoint.Region == opts.Region {
					endpoints = append(endpoints, endpoint)
				}
			}
		}
	}

	// Report an error if the options were ambiguous.
	if len(endpoints) > 1 {
		return "", fmt.Errorf("Discovered %d matching endpoints: %#v", len(endpoints), endpoints)
	}

	// Extract the appropriate URL from the matching Endpoint.
	for _, endpoint := range endpoints {
		switch opts.Availability {
		case gophercloud.AvailabilityPublic:
			return gophercloud.NormalizeURL(endpoint.PublicURL), nil
		case gophercloud.AvailabilityInternal:
			return gophercloud.NormalizeURL(endpoint.InternalURL), nil
		case gophercloud.AvailabilityAdmin:
			return gophercloud.NormalizeURL(endpoint.AdminURL), nil
		default:
			return "", fmt.Errorf("Unexpected availability in endpoint query: %s", opts.Availability)
		}
	}

	// Report an error if there were no matching endpoints.
	return "", gophercloud.ErrEndpointNotFound
}

// V3EndpointURL discovers the endpoint URL for a specific service from a Catalog acquired
// during the v3 identity service. The specified EndpointOpts are used to identify a unique,
// unambiguous endpoint to return. It's an error both when multiple endpoints match the provided
// criteria and when none do. The minimum that can be specified is a Type, but you will also often
// need to specify a Name and/or a Region depending on what's available on your OpenStack
// deployment.
func V3EndpointURL(catalog *tokens3.ServiceCatalog, opts gophercloud.EndpointOpts) (string, error) {
	// Extract Endpoints from the catalog entries that match the requested Type, Interface,
	// Name if provided, and Region if provided.
	var endpoints = make([]tokens3.Endpoint, 0, 1)
	for _, entry := range catalog.Entries {
		if (entry.Type == opts.Type) && (opts.Name == "" || entry.Name == opts.Name) {
			for _, endpoint := range entry.Endpoints {
				if opts.Availability != gophercloud.AvailabilityAdmin &&
					opts.Availability != gophercloud.AvailabilityPublic &&
					opts.Availability != gophercloud.AvailabilityInternal {
					return "", fmt.Errorf("Unexpected availability in endpoint query: %s", opts.Availability)
				}
				if (opts.Availability == gophercloud.Availability(endpoint.Interface)) &&
					(opts.Region == "" || endpoint.Region == opts.Region) {
					endpoints = append(endpoints, endpoint)
				}
			}
		}
	}

	// Report an error if the options were ambiguous.
	if len(endpoints) > 1 {
		return "", fmt.Errorf("Discovered %d matching endpoints: %#v", len(endpoints), endpoints)
	}

	// Extract the URL from the matching Endpoint.
	for _, endpoint := range endpoints {
		return gophercloud.NormalizeURL(endpoint.URL), nil
	}

	// Report an error if there were no matching endpoints.
	return "", gophercloud.ErrEndpointNotFound
}
