// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasstreamconnections

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasStreamConnectionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.25.0/docs/data-sources/stream_connections#instance_name DataMongodbatlasStreamConnections#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.25.0/docs/data-sources/stream_connections#project_id DataMongodbatlasStreamConnections#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.25.0/docs/data-sources/stream_connections#items_per_page DataMongodbatlasStreamConnections#items_per_page}.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.25.0/docs/data-sources/stream_connections#page_num DataMongodbatlasStreamConnections#page_num}.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
}

