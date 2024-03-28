// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupsnapshotexportjob


type CloudBackupSnapshotExportJobCustomData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.3/docs/resources/cloud_backup_snapshot_export_job#key CloudBackupSnapshotExportJob#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.3/docs/resources/cloud_backup_snapshot_export_job#value CloudBackupSnapshotExportJob#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

