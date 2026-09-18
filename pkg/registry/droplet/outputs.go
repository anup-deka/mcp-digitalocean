package droplet

import (
	"github.com/digitalocean/godo"

	"mcp-digitalocean/pkg/registry/common"
)

// Summary shapes for the three list tools that return a curated field subset
// rather than the full godo type.
//
// These replace per-handler map[string]any literals. The fields deliberately
// carry no omitempty: a map always serialises every key it holds, so dropping
// zero-valued fields here would change what these tools have always returned.
type (
	// dropletSummary is godo.Droplet without GPUPartitionMode, which the DO
	// API only populates on the create response.
	dropletSummary struct {
		ID               int                `json:"id"`
		Name             string             `json:"name"`
		Memory           int                `json:"memory"`
		Vcpus            int                `json:"vcpus"`
		Disk             int                `json:"disk"`
		Region           *godo.Region       `json:"region"`
		Image            *godo.Image        `json:"image"`
		Size             *godo.Size         `json:"size"`
		SizeSlug         string             `json:"size_slug"`
		BackupIDs        []int              `json:"backup_ids"`
		NextBackupWindow *godo.BackupWindow `json:"next_backup_window"`
		SnapshotIDs      []int              `json:"snapshot_ids"`
		Features         []string           `json:"features"`
		Locked           bool               `json:"locked"`
		Status           string             `json:"status"`
		Networks         *godo.Networks     `json:"networks"`
		CreatedAt        string             `json:"created_at"`
		Kernel           *godo.Kernel       `json:"kernel"`
		Tags             []string           `json:"tags"`
		VolumeIDs        []string           `json:"volume_ids"`
		VPCUUID          string             `json:"vpc_uuid"`
	}

	imageSummary struct {
		ID           int      `json:"id"`
		Name         string   `json:"name"`
		Slug         string   `json:"slug"`
		Distribution string   `json:"distribution"`
		Type         string   `json:"type"`
		Public       bool     `json:"public"`
		Regions      []string `json:"regions"`
		CreatedAt    string   `json:"created_at"`
		MinDiskSize  int      `json:"min_disk_size"`
	}

	sizeSummary struct {
		Slug         string   `json:"slug"`
		Available    bool     `json:"available"`
		PriceMonthly float64  `json:"price_monthly"`
		PriceHourly  float64  `json:"price_hourly"`
		Memory       int      `json:"memory"`
		Vcpus        int      `json:"vcpus"`
		Disk         int      `json:"disk"`
		Transfer     float64  `json:"transfer"`
		Regions      []string `json:"regions"`
	}
)

// newDropletSummary projects a godo.Droplet onto the shape droplet-list
// returns.
func newDropletSummary(d godo.Droplet) dropletSummary {
	return dropletSummary{
		ID:               d.ID,
		Name:             d.Name,
		Memory:           d.Memory,
		Vcpus:            d.Vcpus,
		Disk:             d.Disk,
		Region:           d.Region,
		Image:            d.Image,
		Size:             d.Size,
		SizeSlug:         d.SizeSlug,
		BackupIDs:        d.BackupIDs,
		NextBackupWindow: d.NextBackupWindow,
		SnapshotIDs:      d.SnapshotIDs,
		Features:         d.Features,
		Locked:           d.Locked,
		Status:           d.Status,
		Networks:         d.Networks,
		CreatedAt:        d.Created,
		Kernel:           d.Kernel,
		Tags:             d.Tags,
		VolumeIDs:        d.VolumeIDs,
		VPCUUID:          d.VPCUUID,
	}
}

func newImageSummary(i godo.Image) imageSummary {
	return imageSummary{
		ID:           i.ID,
		Name:         i.Name,
		Slug:         i.Slug,
		Distribution: i.Distribution,
		Type:         i.Type,
		Public:       i.Public,
		Regions:      i.Regions,
		CreatedAt:    i.Created,
		MinDiskSize:  i.MinDiskSize,
	}
}

func newSizeSummary(s godo.Size) sizeSummary {
	return sizeSummary{
		Slug:         s.Slug,
		Available:    s.Available,
		PriceMonthly: s.PriceMonthly,
		PriceHourly:  s.PriceHourly,
		Memory:       s.Memory,
		Vcpus:        s.Vcpus,
		Disk:         s.Disk,
		Transfer:     s.Transfer,
		Regions:      s.Regions,
	}
}

// Structured output contracts for this package's tools. See the account
// package for the convention.
//
// Most of this service is droplet and image actions, which all resolve to an
// action the caller polls: the 16 single-droplet actions and image actions
// share actionOut, and the 9 by-tag variants share actionsOut.
//
// Tools that stay text-only: droplet-delete and image-delete return a
// confirmation message, not a payload.
var (
	dropletOut          = common.NewOutput[*godo.Droplet]("droplet")
	dropletSummariesOut = common.NewOutput[[]dropletSummary]("droplets")
	neighborsOut        = common.NewOutput[[]godo.Droplet]("droplets")
	kernelsOut          = common.NewOutput[[]godo.Kernel]("kernels")
	backupPolicyOut     = common.NewOutput[*godo.DropletBackupPolicy]("backup_policy")

	imageOut          = common.NewOutput[*godo.Image]("image")
	imageSummariesOut = common.NewOutput[[]imageSummary]("images")
	sizeSummariesOut  = common.NewOutput[[]sizeSummary]("sizes")

	actionOut  = common.NewOutput[*godo.Action]("action")
	actionsOut = common.NewOutput[[]godo.Action]("actions")
)
