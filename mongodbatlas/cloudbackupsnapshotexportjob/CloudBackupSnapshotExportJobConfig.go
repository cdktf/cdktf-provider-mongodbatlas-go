// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupsnapshotexportjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudBackupSnapshotExportJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/cloud_backup_snapshot_export_job#cluster_name CloudBackupSnapshotExportJob#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// custom_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/cloud_backup_snapshot_export_job#custom_data CloudBackupSnapshotExportJob#custom_data}
	CustomData interface{} `field:"required" json:"customData" yaml:"customData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/cloud_backup_snapshot_export_job#export_bucket_id CloudBackupSnapshotExportJob#export_bucket_id}.
	ExportBucketId *string `field:"required" json:"exportBucketId" yaml:"exportBucketId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/cloud_backup_snapshot_export_job#project_id CloudBackupSnapshotExportJob#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/cloud_backup_snapshot_export_job#snapshot_id CloudBackupSnapshotExportJob#snapshot_id}.
	SnapshotId *string `field:"required" json:"snapshotId" yaml:"snapshotId"`
}

