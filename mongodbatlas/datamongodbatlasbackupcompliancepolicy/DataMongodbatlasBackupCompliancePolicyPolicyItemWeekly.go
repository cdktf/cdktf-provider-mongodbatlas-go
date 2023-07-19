package datamongodbatlasbackupcompliancepolicy


type DataMongodbatlasBackupCompliancePolicyPolicyItemWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/data-sources/backup_compliance_policy#frequency_interval DataMongodbatlasBackupCompliancePolicy#frequency_interval}.
	FrequencyInterval *float64 `field:"required" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/data-sources/backup_compliance_policy#retention_unit DataMongodbatlasBackupCompliancePolicy#retention_unit}.
	RetentionUnit *string `field:"required" json:"retentionUnit" yaml:"retentionUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/data-sources/backup_compliance_policy#retention_value DataMongodbatlasBackupCompliancePolicy#retention_value}.
	RetentionValue *float64 `field:"required" json:"retentionValue" yaml:"retentionValue"`
}

