/*
Copyright 2024 The Crossplane Authors.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "openpanel"
	modulePath     = "github.com/hops-ops/provider-openpanel"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// All four OpenPanel /manage resources mint server-assigned IDs at create
// time (UUIDs for Client, slug-or-UUID for Organization/Project,
// composite for Reference). None accept a caller-supplied external name,
// so they all use `IdentifierFromProvider`.
var ExternalNameConfigs = map[string]ujconfig.ExternalName{
	"openpanel_organization": ujconfig.IdentifierFromProvider,
	"openpanel_project":      ujconfig.IdentifierFromProvider,
	"openpanel_client":       ujconfig.IdentifierFromProvider,
	"openpanel_reference":    ujconfig.IdentifierFromProvider,
}

// ExternalNameConfigurations + ExternalNameConfigured are defined in
// config/external_name.go so the same helpers are reused for both the
// cluster-scoped and namespaced providers.

func newProvider(rootGroup string) *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	// Organization is the top-level tenant primitive. Project, Client,
	// and Reference are all anchored to it (Client and Reference via
	// project_id; Project lives directly under an Organization in the
	// /manage API).
	pc.AddResourceConfigurator("openpanel_organization", func(r *ujconfig.Resource) {
		r.ShortGroup = "organization"
	})

	pc.AddResourceConfigurator("openpanel_project", func(r *ujconfig.Resource) {
		r.ShortGroup = "project"
		// Project belongs to the caller's root Client's Organization via the
		// /manage API; no explicit organization_id field today. Add a
		// Reference here if/when the upstream API gains per-Project
		// organization scoping.
	})

	pc.AddResourceConfigurator("openpanel_client", func(r *ujconfig.Resource) {
		r.ShortGroup = "client"
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "openpanel_project",
		}
	})

	pc.AddResourceConfigurator("openpanel_reference", func(r *ujconfig.Resource) {
		r.ShortGroup = "reference"
		r.References["project_id"] = ujconfig.Reference{
			TerraformName: "openpanel_project",
		}
	})

	pc.ConfigureResources()
	return pc
}

// GetProvider returns cluster-scoped provider configuration.
func GetProvider() *ujconfig.Provider {
	return newProvider("openpanel.crossplane.io")
}

// GetProviderNamespaced returns namespaced MR provider configuration
// (Crossplane v2).
func GetProviderNamespaced() *ujconfig.Provider {
	return newProvider("openpanel.m.crossplane.io")
}
