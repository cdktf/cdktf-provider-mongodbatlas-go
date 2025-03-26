// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SearchDeploymentConfig struct {
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
	// Label that identifies the cluster to return the search nodes for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.31.0/docs/resources/search_deployment#cluster_name SearchDeployment#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.31.0/docs/resources/search_deployment#project_id SearchDeployment#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// List of settings that configure the search nodes for your cluster.
	//
	// This list is currently limited to defining a single element.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.31.0/docs/resources/search_deployment#specs SearchDeployment#specs}
	Specs interface{} `field:"required" json:"specs" yaml:"specs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.31.0/docs/resources/search_deployment#timeouts SearchDeployment#timeouts}.
	Timeouts *SearchDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

