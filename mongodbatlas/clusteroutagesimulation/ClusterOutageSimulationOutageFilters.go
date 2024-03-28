// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clusteroutagesimulation


type ClusterOutageSimulationOutageFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.3/docs/resources/cluster_outage_simulation#cloud_provider ClusterOutageSimulation#cloud_provider}.
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.3/docs/resources/cluster_outage_simulation#region_name ClusterOutageSimulation#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

