// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupsnapshotexportbucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudBackupSnapshotExportBucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/cloud_backup_snapshot_export_bucket#bucket_name CloudBackupSnapshotExportBucket#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/cloud_backup_snapshot_export_bucket#cloud_provider CloudBackupSnapshotExportBucket#cloud_provider}.
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/cloud_backup_snapshot_export_bucket#iam_role_id CloudBackupSnapshotExportBucket#iam_role_id}.
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/cloud_backup_snapshot_export_bucket#project_id CloudBackupSnapshotExportBucket#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

