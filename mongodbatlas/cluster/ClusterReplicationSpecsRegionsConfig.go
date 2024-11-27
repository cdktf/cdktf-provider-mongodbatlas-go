// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterReplicationSpecsRegionsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/cluster#region_name Cluster#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/cluster#analytics_nodes Cluster#analytics_nodes}.
	AnalyticsNodes *float64 `field:"optional" json:"analyticsNodes" yaml:"analyticsNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/cluster#electable_nodes Cluster#electable_nodes}.
	ElectableNodes *float64 `field:"optional" json:"electableNodes" yaml:"electableNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/cluster#priority Cluster#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/cluster#read_only_nodes Cluster#read_only_nodes}.
	ReadOnlyNodes *float64 `field:"optional" json:"readOnlyNodes" yaml:"readOnlyNodes"`
}

