// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupschedule


type CloudBackupScheduleCopySettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/cloud_backup_schedule#cloud_provider CloudBackupSchedule#cloud_provider}.
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/cloud_backup_schedule#frequencies CloudBackupSchedule#frequencies}.
	Frequencies *[]*string `field:"optional" json:"frequencies" yaml:"frequencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/cloud_backup_schedule#region_name CloudBackupSchedule#region_name}.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/cloud_backup_schedule#replication_spec_id CloudBackupSchedule#replication_spec_id}.
	ReplicationSpecId *string `field:"optional" json:"replicationSpecId" yaml:"replicationSpecId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/cloud_backup_schedule#should_copy_oplogs CloudBackupSchedule#should_copy_oplogs}.
	ShouldCopyOplogs interface{} `field:"optional" json:"shouldCopyOplogs" yaml:"shouldCopyOplogs"`
}

