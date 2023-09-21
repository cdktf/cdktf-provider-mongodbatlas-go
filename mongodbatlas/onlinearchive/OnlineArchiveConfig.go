// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinearchive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OnlineArchiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#cluster_name OnlineArchive#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#coll_name OnlineArchive#coll_name}.
	CollName *string `field:"required" json:"collName" yaml:"collName"`
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#criteria OnlineArchive#criteria}
	Criteria *OnlineArchiveCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#db_name OnlineArchive#db_name}.
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#project_id OnlineArchive#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#collection_type OnlineArchive#collection_type}.
	CollectionType *string `field:"optional" json:"collectionType" yaml:"collectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#id OnlineArchive#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// partition_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#partition_fields OnlineArchive#partition_fields}
	PartitionFields interface{} `field:"optional" json:"partitionFields" yaml:"partitionFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#paused OnlineArchive#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#schedule OnlineArchive#schedule}
	Schedule *OnlineArchiveSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/online_archive#sync_creation OnlineArchive#sync_creation}.
	SyncCreation interface{} `field:"optional" json:"syncCreation" yaml:"syncCreation"`
}

