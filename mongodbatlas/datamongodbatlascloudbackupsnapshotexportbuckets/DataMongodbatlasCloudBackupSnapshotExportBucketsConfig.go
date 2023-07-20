package datamongodbatlascloudbackupsnapshotexportbuckets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasCloudBackupSnapshotExportBucketsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets#project_id DataMongodbatlasCloudBackupSnapshotExportBuckets#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets#id DataMongodbatlasCloudBackupSnapshotExportBuckets#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets#items_per_page DataMongodbatlasCloudBackupSnapshotExportBuckets#items_per_page}.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets#page_num DataMongodbatlasCloudBackupSnapshotExportBuckets#page_num}.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
}

