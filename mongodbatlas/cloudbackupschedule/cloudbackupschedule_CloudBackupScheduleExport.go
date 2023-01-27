package cloudbackupschedule


type CloudBackupScheduleExport struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_backup_schedule#export_bucket_id CloudBackupSchedule#export_bucket_id}.
	ExportBucketId *string `field:"optional" json:"exportBucketId" yaml:"exportBucketId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_backup_schedule#frequency_type CloudBackupSchedule#frequency_type}.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
}

