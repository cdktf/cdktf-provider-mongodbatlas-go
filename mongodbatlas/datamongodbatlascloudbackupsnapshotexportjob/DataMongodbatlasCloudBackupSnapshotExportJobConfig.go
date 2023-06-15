package datamongodbatlascloudbackupsnapshotexportjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasCloudBackupSnapshotExportJobConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/cloud_backup_snapshot_export_job#cluster_name DataMongodbatlasCloudBackupSnapshotExportJob#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/cloud_backup_snapshot_export_job#export_job_id DataMongodbatlasCloudBackupSnapshotExportJob#export_job_id}.
	ExportJobId *string `field:"required" json:"exportJobId" yaml:"exportJobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/cloud_backup_snapshot_export_job#id DataMongodbatlasCloudBackupSnapshotExportJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.0/docs/data-sources/cloud_backup_snapshot_export_job#project_id DataMongodbatlasCloudBackupSnapshotExportJob#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

