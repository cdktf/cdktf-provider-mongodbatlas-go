// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudBackupScheduleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#cluster_name CloudBackupSchedule#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#project_id CloudBackupSchedule#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#auto_export_enabled CloudBackupSchedule#auto_export_enabled}.
	AutoExportEnabled interface{} `field:"optional" json:"autoExportEnabled" yaml:"autoExportEnabled"`
	// copy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#copy_settings CloudBackupSchedule#copy_settings}
	CopySettings interface{} `field:"optional" json:"copySettings" yaml:"copySettings"`
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#export CloudBackupSchedule#export}
	Export *CloudBackupScheduleExport `field:"optional" json:"export" yaml:"export"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#id CloudBackupSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// policy_item_daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#policy_item_daily CloudBackupSchedule#policy_item_daily}
	PolicyItemDaily *CloudBackupSchedulePolicyItemDaily `field:"optional" json:"policyItemDaily" yaml:"policyItemDaily"`
	// policy_item_hourly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#policy_item_hourly CloudBackupSchedule#policy_item_hourly}
	PolicyItemHourly *CloudBackupSchedulePolicyItemHourly `field:"optional" json:"policyItemHourly" yaml:"policyItemHourly"`
	// policy_item_monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#policy_item_monthly CloudBackupSchedule#policy_item_monthly}
	PolicyItemMonthly interface{} `field:"optional" json:"policyItemMonthly" yaml:"policyItemMonthly"`
	// policy_item_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#policy_item_weekly CloudBackupSchedule#policy_item_weekly}
	PolicyItemWeekly interface{} `field:"optional" json:"policyItemWeekly" yaml:"policyItemWeekly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#reference_hour_of_day CloudBackupSchedule#reference_hour_of_day}.
	ReferenceHourOfDay *float64 `field:"optional" json:"referenceHourOfDay" yaml:"referenceHourOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#reference_minute_of_hour CloudBackupSchedule#reference_minute_of_hour}.
	ReferenceMinuteOfHour *float64 `field:"optional" json:"referenceMinuteOfHour" yaml:"referenceMinuteOfHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#restore_window_days CloudBackupSchedule#restore_window_days}.
	RestoreWindowDays *float64 `field:"optional" json:"restoreWindowDays" yaml:"restoreWindowDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#update_snapshots CloudBackupSchedule#update_snapshots}.
	UpdateSnapshots interface{} `field:"optional" json:"updateSnapshots" yaml:"updateSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.2/docs/resources/cloud_backup_schedule#use_org_and_group_names_in_export_prefix CloudBackupSchedule#use_org_and_group_names_in_export_prefix}.
	UseOrgAndGroupNamesInExportPrefix interface{} `field:"optional" json:"useOrgAndGroupNamesInExportPrefix" yaml:"useOrgAndGroupNamesInExportPrefix"`
}

