package diskconfig

import (
	"errors"

	"github.com/ttsubo2000/gophercloud/openstack/compute/v2/servers"
)

// DiskConfig represents one of the two possible settings for the DiskConfig option when creating,
// rebuilding, or resizing servers: Auto or Manual.
type DiskConfig string

const (
	// Auto builds a server with a single partition the size of the target flavor disk and
	// automatically adjusts the filesystem to fit the entire partition. Auto may only be used with
	// images and servers that use a single EXT3 partition.
	Auto DiskConfig = "AUTO"

	// Manual builds a server using whatever partition scheme and filesystem are present in the source
	// image. If the target flavor disk is larger, the remaining space is left unpartitioned. This
	// enables images to have non-EXT3 filesystems, multiple partitions, and so on, and enables you
	// to manage the disk configuration. It also results in slightly shorter boot times.
	Manual DiskConfig = "MANUAL"
)

// ErrInvalidDiskConfig is returned if an invalid string is specified for a DiskConfig option.
var ErrInvalidDiskConfig = errors.New("DiskConfig must be either diskconfig.Auto or diskconfig.Manual.")

// Validate ensures that a DiskConfig contains an appropriate value.
func (config DiskConfig) validate() error {
	switch config {
	case Auto, Manual:
		return nil
	default:
		return ErrInvalidDiskConfig
	}
}

// CreateOptsExt adds a DiskConfig option to the base CreateOpts.
type CreateOptsExt struct {
	servers.CreateOptsBuilder

	// DiskConfig [optional] controls how the created server's disk is partitioned.
	DiskConfig DiskConfig `json:"OS-DCF:diskConfig,omitempty"`
}

// ToServerCreateMap adds the diskconfig option to the base server creation options.
func (opts CreateOptsExt) ToServerCreateMap() (map[string]interface{}, error) {
	base, err := opts.CreateOptsBuilder.ToServerCreateMap()
	if err != nil {
		return nil, err
	}

	if string(opts.DiskConfig) == "" {
		return base, nil
	}

	serverMap := base["server"].(map[string]interface{})
	serverMap["OS-DCF:diskConfig"] = string(opts.DiskConfig)

	return base, nil
}

// RebuildOptsExt adds a DiskConfig option to the base RebuildOpts.
type RebuildOptsExt struct {
	servers.RebuildOptsBuilder

	// DiskConfig [optional] controls how the rebuilt server's disk is partitioned.
	DiskConfig DiskConfig
}

// ToServerRebuildMap adds the diskconfig option to the base server rebuild options.
func (opts RebuildOptsExt) ToServerRebuildMap() (map[string]interface{}, error) {
	err := opts.DiskConfig.validate()
	if err != nil {
		return nil, err
	}

	base, err := opts.RebuildOptsBuilder.ToServerRebuildMap()
	if err != nil {
		return nil, err
	}

	serverMap := base["rebuild"].(map[string]interface{})
	serverMap["OS-DCF:diskConfig"] = string(opts.DiskConfig)

	return base, nil
}

// ResizeOptsExt adds a DiskConfig option to the base server resize options.
type ResizeOptsExt struct {
	servers.ResizeOptsBuilder

	// DiskConfig [optional] controls how the resized server's disk is partitioned.
	DiskConfig DiskConfig
}

// ToServerResizeMap adds the diskconfig option to the base server creation options.
func (opts ResizeOptsExt) ToServerResizeMap() (map[string]interface{}, error) {
	err := opts.DiskConfig.validate()
	if err != nil {
		return nil, err
	}

	base, err := opts.ResizeOptsBuilder.ToServerResizeMap()
	if err != nil {
		return nil, err
	}

	serverMap := base["resize"].(map[string]interface{})
	serverMap["OS-DCF:diskConfig"] = string(opts.DiskConfig)

	return base, nil
}
