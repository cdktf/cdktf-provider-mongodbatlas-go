// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupcompliancepolicy


type BackupCompliancePolicyOnDemandPolicyItem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.26.0/docs/resources/backup_compliance_policy#frequency_interval BackupCompliancePolicy#frequency_interval}.
	FrequencyInterval *float64 `field:"required" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.26.0/docs/resources/backup_compliance_policy#retention_unit BackupCompliancePolicy#retention_unit}.
	RetentionUnit *string `field:"required" json:"retentionUnit" yaml:"retentionUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.26.0/docs/resources/backup_compliance_policy#retention_value BackupCompliancePolicy#retention_value}.
	RetentionValue *float64 `field:"required" json:"retentionValue" yaml:"retentionValue"`
}

