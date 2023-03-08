package cloudbackupschedule


type CloudBackupSchedulePolicyItemHourly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_backup_schedule#frequency_interval CloudBackupSchedule#frequency_interval}.
	FrequencyInterval *float64 `field:"required" json:"frequencyInterval" yaml:"frequencyInterval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_backup_schedule#retention_unit CloudBackupSchedule#retention_unit}.
	RetentionUnit *string `field:"required" json:"retentionUnit" yaml:"retentionUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_backup_schedule#retention_value CloudBackupSchedule#retention_value}.
	RetentionValue *float64 `field:"required" json:"retentionValue" yaml:"retentionValue"`
}

