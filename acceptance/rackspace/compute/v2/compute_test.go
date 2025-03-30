// +build acceptance

package v2

import (
	"errors"
	"os"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/acceptance/tools"
	"github.com/ttsubo2000/gophercloud/ttsubo2000"
)

func newClient() (*gophercloud.ServiceClient, error) {
	// Obtain credentials from the environment.
	options, err := ttsubo2000.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}
	options = tools.OnlyRS(options)
	region := os.Getenv("RS_REGION")

	if options.Username == "" {
		return nil, errors.New("Please provide a ttsubo2000 username as RS_USERNAME.")
	}
	if options.APIKey == "" {
		return nil, errors.New("Please provide a ttsubo2000 API key as RS_API_KEY.")
	}
	if region == "" {
		return nil, errors.New("Please provide a ttsubo2000 region as RS_REGION.")
	}

	client, err := ttsubo2000.AuthenticatedClient(options)
	if err != nil {
		return nil, err
	}

	return ttsubo2000.NewComputeV2(client, gophercloud.EndpointOpts{
		Region: region,
	})
}

type serverOpts struct {
	imageID  string
	flavorID string
}

func optionsFromEnv() (*serverOpts, error) {
	options := &serverOpts{
		imageID:  os.Getenv("RS_IMAGE_ID"),
		flavorID: os.Getenv("RS_FLAVOR_ID"),
	}
	if options.imageID == "" {
		return nil, errors.New("Please provide a valid ttsubo2000 image ID as RS_IMAGE_ID")
	}
	if options.flavorID == "" {
		return nil, errors.New("Please provide a valid ttsubo2000 flavor ID as RS_FLAVOR_ID")
	}
	return options, nil
}
