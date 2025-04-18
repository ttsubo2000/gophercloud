package configurations

import (
	"errors"

	"github.com/ttsubo2000/gophercloud"
	"github.com/ttsubo2000/gophercloud/openstack/db/v1/instances"
	"github.com/ttsubo2000/gophercloud/pagination"
)

// List will list all of the available configurations.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	pageFn := func(r pagination.PageResult) pagination.Page {
		return ConfigPage{pagination.SinglePageBase(r)}
	}

	return pagination.NewPager(client, baseURL(client), pageFn)
}

// CreateOptsBuilder is a top-level interface which renders a JSON map.
type CreateOptsBuilder interface {
	ToConfigCreateMap() (map[string]interface{}, error)
}

// DatastoreOpts is the primary options struct for creating and modifying
// how configuration resources are associated with datastores.
type DatastoreOpts struct {
	// [OPTIONAL] The type of datastore. Defaults to "MySQL".
	Type string

	// [OPTIONAL] The specific version of a datastore. Defaults to "5.6".
	Version string
}

// ToMap renders a JSON map for a datastore setting.
func (opts DatastoreOpts) ToMap() (map[string]string, error) {
	datastore := map[string]string{}

	if opts.Type != "" {
		datastore["type"] = opts.Type
	}

	if opts.Version != "" {
		datastore["version"] = opts.Version
	}

	return datastore, nil
}

// CreateOpts is the struct responsible for configuring new configurations.
type CreateOpts struct {
	// [REQUIRED] The configuration group name
	Name string

	// [REQUIRED] A map of user-defined configuration settings that will define
	// how each associated datastore works. Each key/value pair is specific to a
	// datastore type.
	Values map[string]interface{}

	// [OPTIONAL] Associates the configuration group with a particular datastore.
	Datastore *DatastoreOpts

	// [OPTIONAL] A human-readable explanation for the group.
	Description string
}

// ToConfigCreateMap casts a CreateOpts struct into a JSON map.
func (opts CreateOpts) ToConfigCreateMap() (map[string]interface{}, error) {
	if opts.Name == "" {
		return nil, errors.New("Name is a required field")
	}
	if len(opts.Values) == 0 {
		return nil, errors.New("Values must be a populated map")
	}

	config := map[string]interface{}{
		"name":   opts.Name,
		"values": opts.Values,
	}

	if opts.Datastore != nil {
		ds, err := opts.Datastore.ToMap()
		if err != nil {
			return config, err
		}
		config["datastore"] = ds
	}

	if opts.Description != "" {
		config["description"] = opts.Description
	}

	return map[string]interface{}{"configuration": config}, nil
}

// Create will create a new configuration group.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToConfigCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Request("POST", baseURL(client), gophercloud.RequestOpts{
		OkCodes:      []int{200},
		JSONBody:     &reqBody,
		JSONResponse: &res.Body,
	})

	return res
}

// Get will retrieve the details for a specified configuration group.
func Get(client *gophercloud.ServiceClient, configID string) GetResult {
	var res GetResult

	_, res.Err = client.Request("GET", resourceURL(client, configID), gophercloud.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: &res.Body,
	})

	return res
}

// UpdateOptsBuilder is the top-level interface for casting update options into
// JSON maps.
type UpdateOptsBuilder interface {
	ToConfigUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the struct responsible for modifying existing configurations.
type UpdateOpts struct {
	// [OPTIONAL] The configuration group name
	Name string

	// [OPTIONAL] A map of user-defined configuration settings that will define
	// how each associated datastore works. Each key/value pair is specific to a
	// datastore type.
	Values map[string]interface{}

	// [OPTIONAL] Associates the configuration group with a particular datastore.
	Datastore *DatastoreOpts

	// [OPTIONAL] A human-readable explanation for the group.
	Description string
}

// ToConfigUpdateMap will cast an UpdateOpts struct into a JSON map.
func (opts UpdateOpts) ToConfigUpdateMap() (map[string]interface{}, error) {
	config := map[string]interface{}{}

	if opts.Name != "" {
		config["name"] = opts.Name
	}

	if opts.Description != "" {
		config["description"] = opts.Description
	}

	if opts.Datastore != nil {
		ds, err := opts.Datastore.ToMap()
		if err != nil {
			return config, err
		}
		config["datastore"] = ds
	}

	if len(opts.Values) > 0 {
		config["values"] = opts.Values
	}

	return map[string]interface{}{"configuration": config}, nil
}

// Update will modify an existing configuration group by performing a merge
// between new and existing values. If the key already exists, the new value
// will overwrite. All other keys will remain unaffected.
func Update(client *gophercloud.ServiceClient, configID string, opts UpdateOptsBuilder) UpdateResult {
	var res UpdateResult

	reqBody, err := opts.ToConfigUpdateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Request("PATCH", resourceURL(client, configID), gophercloud.RequestOpts{
		OkCodes:  []int{200},
		JSONBody: &reqBody,
	})

	return res
}

// Replace will modify an existing configuration group by overwriting the
// entire parameter group with the new values provided. Any existing keys not
// included in UpdateOptsBuilder will be deleted.
func Replace(client *gophercloud.ServiceClient, configID string, opts UpdateOptsBuilder) ReplaceResult {
	var res ReplaceResult

	reqBody, err := opts.ToConfigUpdateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Request("PUT", resourceURL(client, configID), gophercloud.RequestOpts{
		OkCodes:  []int{202},
		JSONBody: &reqBody,
	})

	return res
}

// Delete will permanently delete a configuration group. Please note that
// config groups cannot be deleted whilst still attached to running instances -
// you must detach and then delete them.
func Delete(client *gophercloud.ServiceClient, configID string) DeleteResult {
	var res DeleteResult

	_, res.Err = client.Request("DELETE", resourceURL(client, configID), gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return res
}

// ListInstances will list all the instances associated with a particular
// configuration group.
func ListInstances(client *gophercloud.ServiceClient, configID string) pagination.Pager {
	pageFn := func(r pagination.PageResult) pagination.Page {
		return instances.InstancePage{pagination.LinkedPageBase{PageResult: r}}
	}
	return pagination.NewPager(client, instancesURL(client, configID), pageFn)
}

// ListDatastoreParams will list all the available and supported parameters
// that can be used for a particular datastore ID and a particular version.
// For example, if you are wondering how you can configure a MySQL 5.6 instance,
// you can use this operation (you will need to retrieve the MySQL datastore ID
// by using the datastores API).
func ListDatastoreParams(client *gophercloud.ServiceClient, datastoreID, versionID string) pagination.Pager {
	pageFn := func(r pagination.PageResult) pagination.Page {
		return ParamPage{pagination.SinglePageBase(r)}
	}
	return pagination.NewPager(client, listDSParamsURL(client, datastoreID, versionID), pageFn)
}

// GetDatastoreParam will retrieve information about a specific configuration
// parameter. For example, you can use this operation to understand more about
// "innodb_file_per_table" configuration param for MySQL datastores. You will
// need the param's ID first, which can be attained by using the ListDatastoreParams
// operation.
func GetDatastoreParam(client *gophercloud.ServiceClient, datastoreID, versionID, paramID string) ParamResult {
	var res ParamResult

	_, res.Err = client.Request("GET", getDSParamURL(client, datastoreID, versionID, paramID), gophercloud.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: &res.Body,
	})

	return res
}

// ListGlobalParams is similar to ListDatastoreParams but does not require a
// DatastoreID.
func ListGlobalParams(client *gophercloud.ServiceClient, versionID string) pagination.Pager {
	pageFn := func(r pagination.PageResult) pagination.Page {
		return ParamPage{pagination.SinglePageBase(r)}
	}
	return pagination.NewPager(client, listGlobalParamsURL(client, versionID), pageFn)
}

// GetGlobalParam is similar to GetDatastoreParam but does not require a
// DatastoreID.
func GetGlobalParam(client *gophercloud.ServiceClient, versionID, paramID string) ParamResult {
	var res ParamResult

	_, res.Err = client.Request("GET", getGlobalParamURL(client, versionID, paramID), gophercloud.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: &res.Body,
	})

	return res
}
