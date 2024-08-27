// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupsnapshotrestorejob


type CloudBackupSnapshotRestoreJobDeliveryTypeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#automated CloudBackupSnapshotRestoreJob#automated}.
	Automated interface{} `field:"optional" json:"automated" yaml:"automated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#download CloudBackupSnapshotRestoreJob#download}.
	Download interface{} `field:"optional" json:"download" yaml:"download"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#oplog_inc CloudBackupSnapshotRestoreJob#oplog_inc}.
	OplogInc *float64 `field:"optional" json:"oplogInc" yaml:"oplogInc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#oplog_ts CloudBackupSnapshotRestoreJob#oplog_ts}.
	OplogTs *float64 `field:"optional" json:"oplogTs" yaml:"oplogTs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#point_in_time CloudBackupSnapshotRestoreJob#point_in_time}.
	PointInTime interface{} `field:"optional" json:"pointInTime" yaml:"pointInTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#point_in_time_utc_seconds CloudBackupSnapshotRestoreJob#point_in_time_utc_seconds}.
	PointInTimeUtcSeconds *float64 `field:"optional" json:"pointInTimeUtcSeconds" yaml:"pointInTimeUtcSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#target_cluster_name CloudBackupSnapshotRestoreJob#target_cluster_name}.
	TargetClusterName *string `field:"optional" json:"targetClusterName" yaml:"targetClusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/cloud_backup_snapshot_restore_job#target_project_id CloudBackupSnapshotRestoreJob#target_project_id}.
	TargetProjectId *string `field:"optional" json:"targetProjectId" yaml:"targetProjectId"`
}

