// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupsnapshot


type CloudBackupSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/cloud_backup_snapshot#create CloudBackupSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

