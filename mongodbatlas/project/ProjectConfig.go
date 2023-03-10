package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#name Project#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#org_id Project#org_id}.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// api_keys block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#api_keys Project#api_keys}
	ApiKeys interface{} `field:"optional" json:"apiKeys" yaml:"apiKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#id Project#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#is_collect_database_specifics_statistics_enabled Project#is_collect_database_specifics_statistics_enabled}.
	IsCollectDatabaseSpecificsStatisticsEnabled interface{} `field:"optional" json:"isCollectDatabaseSpecificsStatisticsEnabled" yaml:"isCollectDatabaseSpecificsStatisticsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#is_data_explorer_enabled Project#is_data_explorer_enabled}.
	IsDataExplorerEnabled interface{} `field:"optional" json:"isDataExplorerEnabled" yaml:"isDataExplorerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#is_performance_advisor_enabled Project#is_performance_advisor_enabled}.
	IsPerformanceAdvisorEnabled interface{} `field:"optional" json:"isPerformanceAdvisorEnabled" yaml:"isPerformanceAdvisorEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#is_realtime_performance_panel_enabled Project#is_realtime_performance_panel_enabled}.
	IsRealtimePerformancePanelEnabled interface{} `field:"optional" json:"isRealtimePerformancePanelEnabled" yaml:"isRealtimePerformancePanelEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#is_schema_advisor_enabled Project#is_schema_advisor_enabled}.
	IsSchemaAdvisorEnabled interface{} `field:"optional" json:"isSchemaAdvisorEnabled" yaml:"isSchemaAdvisorEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#project_owner_id Project#project_owner_id}.
	ProjectOwnerId *string `field:"optional" json:"projectOwnerId" yaml:"projectOwnerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#region_usage_restrictions Project#region_usage_restrictions}.
	RegionUsageRestrictions *string `field:"optional" json:"regionUsageRestrictions" yaml:"regionUsageRestrictions"`
	// teams block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#teams Project#teams}
	Teams interface{} `field:"optional" json:"teams" yaml:"teams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#with_default_alerts_settings Project#with_default_alerts_settings}.
	WithDefaultAlertsSettings interface{} `field:"optional" json:"withDefaultAlertsSettings" yaml:"withDefaultAlertsSettings"`
}

