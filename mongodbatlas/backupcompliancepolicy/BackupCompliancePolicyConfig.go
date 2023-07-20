package backupcompliancepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupCompliancePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#authorized_email BackupCompliancePolicy#authorized_email}.
	AuthorizedEmail *string `field:"required" json:"authorizedEmail" yaml:"authorizedEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#copy_protection_enabled BackupCompliancePolicy#copy_protection_enabled}.
	CopyProtectionEnabled interface{} `field:"required" json:"copyProtectionEnabled" yaml:"copyProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#encryption_at_rest_enabled BackupCompliancePolicy#encryption_at_rest_enabled}.
	EncryptionAtRestEnabled interface{} `field:"required" json:"encryptionAtRestEnabled" yaml:"encryptionAtRestEnabled"`
	// on_demand_policy_item block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#on_demand_policy_item BackupCompliancePolicy#on_demand_policy_item}
	OnDemandPolicyItem *BackupCompliancePolicyOnDemandPolicyItem `field:"required" json:"onDemandPolicyItem" yaml:"onDemandPolicyItem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#pit_enabled BackupCompliancePolicy#pit_enabled}.
	PitEnabled interface{} `field:"required" json:"pitEnabled" yaml:"pitEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#project_id BackupCompliancePolicy#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#id BackupCompliancePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// policy_item_daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#policy_item_daily BackupCompliancePolicy#policy_item_daily}
	PolicyItemDaily *BackupCompliancePolicyPolicyItemDaily `field:"optional" json:"policyItemDaily" yaml:"policyItemDaily"`
	// policy_item_hourly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#policy_item_hourly BackupCompliancePolicy#policy_item_hourly}
	PolicyItemHourly *BackupCompliancePolicyPolicyItemHourly `field:"optional" json:"policyItemHourly" yaml:"policyItemHourly"`
	// policy_item_monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#policy_item_monthly BackupCompliancePolicy#policy_item_monthly}
	PolicyItemMonthly interface{} `field:"optional" json:"policyItemMonthly" yaml:"policyItemMonthly"`
	// policy_item_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#policy_item_weekly BackupCompliancePolicy#policy_item_weekly}
	PolicyItemWeekly interface{} `field:"optional" json:"policyItemWeekly" yaml:"policyItemWeekly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/backup_compliance_policy#restore_window_days BackupCompliancePolicy#restore_window_days}.
	RestoreWindowDays *float64 `field:"optional" json:"restoreWindowDays" yaml:"restoreWindowDays"`
}

