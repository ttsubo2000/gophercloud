package users

import (
	"errors"

	"github.com/ttsubo2000/gophercloud"
	db "github.com/ttsubo2000/gophercloud/openstack/db/v1/databases"
	"github.com/ttsubo2000/gophercloud/pagination"
)

// CreateOptsBuilder is the top-level interface for creating JSON maps.
type CreateOptsBuilder interface {
	ToUserCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the struct responsible for configuring a new user; often in the
// context of an instance.
type CreateOpts struct {
	// [REQUIRED] Specifies a name for the user. Valid names can be composed
	// of the following characters: letters (either case); numbers; these
	// characters '@', '?', '#', ' ' but NEVER beginning a name string; '_' is
	// permitted anywhere. Prohibited characters that are forbidden include:
	// single quotes, double quotes, back quotes, semicolons, commas, backslashes,
	// and forward slashes. Spaces at the front or end of a user name are also
	// not permitted.
	Name string

	// [REQUIRED] Specifies a password for the user.
	Password string

	// [OPTIONAL] An array of databases that this user will connect to. The
	// "name" field is the only requirement for each option.
	Databases db.BatchCreateOpts

	// [OPTIONAL] Specifies the host from which a user is allowed to connect to
	// the database. Possible values are a string containing an IPv4 address or
	// "%" to allow connecting from any host. Optional; the default is "%".
	Host string
}

// ToMap is a convenience function for creating sub-maps for individual users.
func (opts CreateOpts) ToMap() (map[string]interface{}, error) {

	if opts.Name == "root" {
		return nil, errors.New("root is a reserved user name and cannot be used")
	}
	if opts.Name == "" {
		return nil, errors.New("Name is a required field")
	}
	if opts.Password == "" {
		return nil, errors.New("Password is a required field")
	}

	user := map[string]interface{}{
		"name":     opts.Name,
		"password": opts.Password,
	}

	if opts.Host != "" {
		user["host"] = opts.Host
	}

	dbs := make([]map[string]string, len(opts.Databases))
	for i, db := range opts.Databases {
		dbs[i] = map[string]string{"name": db.Name}
	}

	if len(dbs) > 0 {
		user["databases"] = dbs
	}

	return user, nil
}

// BatchCreateOpts allows multiple users to be created at once.
type BatchCreateOpts []CreateOpts

// ToUserCreateMap will generate a JSON map.
func (opts BatchCreateOpts) ToUserCreateMap() (map[string]interface{}, error) {
	users := make([]map[string]interface{}, len(opts))
	for i, opt := range opts {
		user, err := opt.ToMap()
		if err != nil {
			return nil, err
		}
		users[i] = user
	}
	return map[string]interface{}{"users": users}, nil
}

// Create asynchronously provisions a new user for the specified database
// instance based on the configuration defined in CreateOpts. If databases are
// assigned for a particular user, the user will be granted all privileges
// for those specified databases. "root" is a reserved name and cannot be used.
func Create(client *gophercloud.ServiceClient, instanceID string, opts CreateOptsBuilder) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToUserCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Request("POST", baseURL(client, instanceID), gophercloud.RequestOpts{
		JSONBody: &reqBody,
		OkCodes:  []int{202},
	})

	return res
}

// List will list all the users associated with a specified database instance,
// along with their associated databases. This operation will not return any
// system users or administrators for a database.
func List(client *gophercloud.ServiceClient, instanceID string) pagination.Pager {
	createPageFn := func(r pagination.PageResult) pagination.Page {
		return UserPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, baseURL(client, instanceID), createPageFn)
}

// Delete will permanently delete a user from a specified database instance.
func Delete(client *gophercloud.ServiceClient, instanceID, userName string) DeleteResult {
	var res DeleteResult

	_, res.Err = client.Request("DELETE", userURL(client, instanceID, userName), gophercloud.RequestOpts{
		OkCodes: []int{202},
	})

	return res
}
