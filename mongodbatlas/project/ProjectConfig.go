// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#name Project#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#org_id Project#org_id}.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_collect_database_specifics_statistics_enabled Project#is_collect_database_specifics_statistics_enabled}.
	IsCollectDatabaseSpecificsStatisticsEnabled interface{} `field:"optional" json:"isCollectDatabaseSpecificsStatisticsEnabled" yaml:"isCollectDatabaseSpecificsStatisticsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_data_explorer_enabled Project#is_data_explorer_enabled}.
	IsDataExplorerEnabled interface{} `field:"optional" json:"isDataExplorerEnabled" yaml:"isDataExplorerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_extended_storage_sizes_enabled Project#is_extended_storage_sizes_enabled}.
	IsExtendedStorageSizesEnabled interface{} `field:"optional" json:"isExtendedStorageSizesEnabled" yaml:"isExtendedStorageSizesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_performance_advisor_enabled Project#is_performance_advisor_enabled}.
	IsPerformanceAdvisorEnabled interface{} `field:"optional" json:"isPerformanceAdvisorEnabled" yaml:"isPerformanceAdvisorEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_realtime_performance_panel_enabled Project#is_realtime_performance_panel_enabled}.
	IsRealtimePerformancePanelEnabled interface{} `field:"optional" json:"isRealtimePerformancePanelEnabled" yaml:"isRealtimePerformancePanelEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_schema_advisor_enabled Project#is_schema_advisor_enabled}.
	IsSchemaAdvisorEnabled interface{} `field:"optional" json:"isSchemaAdvisorEnabled" yaml:"isSchemaAdvisorEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#is_slow_operation_thresholding_enabled Project#is_slow_operation_thresholding_enabled}.
	IsSlowOperationThresholdingEnabled interface{} `field:"optional" json:"isSlowOperationThresholdingEnabled" yaml:"isSlowOperationThresholdingEnabled"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#limits Project#limits}
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#project_owner_id Project#project_owner_id}.
	ProjectOwnerId *string `field:"optional" json:"projectOwnerId" yaml:"projectOwnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#region_usage_restrictions Project#region_usage_restrictions}.
	RegionUsageRestrictions *string `field:"optional" json:"regionUsageRestrictions" yaml:"regionUsageRestrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#tags Project#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// teams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#teams Project#teams}
	Teams interface{} `field:"optional" json:"teams" yaml:"teams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/project#with_default_alerts_settings Project#with_default_alerts_settings}.
	WithDefaultAlertsSettings interface{} `field:"optional" json:"withDefaultAlertsSettings" yaml:"withDefaultAlertsSettings"`
}

