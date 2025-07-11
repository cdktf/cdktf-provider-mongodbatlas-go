// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flexcluster


type FlexClusterProviderSettings struct {
	// Cloud service provider on which MongoDB Cloud provisioned the flex cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#backing_provider_name FlexCluster#backing_provider_name}
	BackingProviderName *string `field:"required" json:"backingProviderName" yaml:"backingProviderName"`
	// Human-readable label that identifies the geographic location of your MongoDB flex cluster.
	//
	// The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#region_name FlexCluster#region_name}
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

