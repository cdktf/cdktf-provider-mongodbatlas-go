// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupschedule


type CloudBackupScheduleExport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/resources/cloud_backup_schedule#export_bucket_id CloudBackupSchedule#export_bucket_id}.
	ExportBucketId *string `field:"optional" json:"exportBucketId" yaml:"exportBucketId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/resources/cloud_backup_schedule#frequency_type CloudBackupSchedule#frequency_type}.
	FrequencyType *string `field:"optional" json:"frequencyType" yaml:"frequencyType"`
}

