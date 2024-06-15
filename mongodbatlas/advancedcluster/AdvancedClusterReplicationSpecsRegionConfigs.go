// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterReplicationSpecsRegionConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#priority AdvancedCluster#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#provider_name AdvancedCluster#provider_name}.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#region_name AdvancedCluster#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// analytics_auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#analytics_auto_scaling AdvancedCluster#analytics_auto_scaling}
	AnalyticsAutoScaling *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling `field:"optional" json:"analyticsAutoScaling" yaml:"analyticsAutoScaling"`
	// analytics_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#analytics_specs AdvancedCluster#analytics_specs}
	AnalyticsSpecs *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs `field:"optional" json:"analyticsSpecs" yaml:"analyticsSpecs"`
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#auto_scaling AdvancedCluster#auto_scaling}
	AutoScaling *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#backing_provider_name AdvancedCluster#backing_provider_name}.
	BackingProviderName *string `field:"optional" json:"backingProviderName" yaml:"backingProviderName"`
	// electable_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#electable_specs AdvancedCluster#electable_specs}
	ElectableSpecs *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs `field:"optional" json:"electableSpecs" yaml:"electableSpecs"`
	// read_only_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/advanced_cluster#read_only_specs AdvancedCluster#read_only_specs}
	ReadOnlySpecs *AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs `field:"optional" json:"readOnlySpecs" yaml:"readOnlySpecs"`
}

