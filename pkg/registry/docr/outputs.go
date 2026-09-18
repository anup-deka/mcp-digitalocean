package docr

import (
	"github.com/digitalocean/godo"

	"mcp-digitalocean/pkg/registry/common"
)

// Paginated result shapes for the list tools. These were previously anonymous
// structs declared inline in each handler; they are named here so the same
// type drives both the declared output schema and the emitted payload.
type (
	repositoryList struct {
		Repositories []*godo.RepositoryV2 `json:"repositories"`
		Meta         *godo.Meta           `json:"meta,omitempty"`
	}

	repositoryTagList struct {
		Tags []*godo.RepositoryTag `json:"tags"`
		Meta *godo.Meta            `json:"meta,omitempty"`
	}

	repositoryManifestList struct {
		Manifests []*godo.RepositoryManifest `json:"manifests"`
		Meta      *godo.Meta                 `json:"meta,omitempty"`
	}

	garbageCollectionList struct {
		GarbageCollections []*godo.GarbageCollection `json:"garbage_collections"`
		Meta               *godo.Meta                `json:"meta,omitempty"`
	}
)

// Structured output contracts for this package's tools. See the account
// package for the convention.
//
// Tools that stay text-only:
//   - docr-delete, docr-repository-tag-delete, docr-repository-manifest-delete
//     and docr-validate-name return a confirmation message, not a payload.
//   - docr-docker-credentials returns a Docker config blob whose shape is
//     defined by Docker rather than by the DigitalOcean API, so there is no
//     godo type to describe it with.
var (
	registryOut     = common.NewOutput[*godo.Registry]("registry")
	registriesOut   = common.NewOutput[[]*godo.Registry]("registries")
	registryOptsOut = common.NewOutput[*godo.RegistryOptions]("options")
	subscriptionOut = common.NewOutput[*godo.RegistrySubscription]("subscription")
	gcOut           = common.NewOutput[*godo.GarbageCollection]("garbage_collection")

	// These payloads are already objects that pair a collection with its
	// pagination metadata, so they are published without an extra envelope.
	repositoryListOut         = common.NewObjectOutput[repositoryList]()
	repositoryTagListOut      = common.NewObjectOutput[repositoryTagList]()
	repositoryManifestListOut = common.NewObjectOutput[repositoryManifestList]()
	garbageCollectionListOut  = common.NewObjectOutput[garbageCollectionList]()
)
