package datamongodbatlasbackupcompliancepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasBackupCompliancePolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#project_id DataMongodbatlasBackupCompliancePolicy#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#id DataMongodbatlasBackupCompliancePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// on_demand_policy_item block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#on_demand_policy_item DataMongodbatlasBackupCompliancePolicy#on_demand_policy_item}
	OnDemandPolicyItem *DataMongodbatlasBackupCompliancePolicyOnDemandPolicyItem `field:"optional" json:"onDemandPolicyItem" yaml:"onDemandPolicyItem"`
	// policy_item_daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#policy_item_daily DataMongodbatlasBackupCompliancePolicy#policy_item_daily}
	PolicyItemDaily *DataMongodbatlasBackupCompliancePolicyPolicyItemDaily `field:"optional" json:"policyItemDaily" yaml:"policyItemDaily"`
	// policy_item_hourly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#policy_item_hourly DataMongodbatlasBackupCompliancePolicy#policy_item_hourly}
	PolicyItemHourly *DataMongodbatlasBackupCompliancePolicyPolicyItemHourly `field:"optional" json:"policyItemHourly" yaml:"policyItemHourly"`
	// policy_item_monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#policy_item_monthly DataMongodbatlasBackupCompliancePolicy#policy_item_monthly}
	PolicyItemMonthly interface{} `field:"optional" json:"policyItemMonthly" yaml:"policyItemMonthly"`
	// policy_item_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/backup_compliance_policy#policy_item_weekly DataMongodbatlasBackupCompliancePolicy#policy_item_weekly}
	PolicyItemWeekly interface{} `field:"optional" json:"policyItemWeekly" yaml:"policyItemWeekly"`
}

