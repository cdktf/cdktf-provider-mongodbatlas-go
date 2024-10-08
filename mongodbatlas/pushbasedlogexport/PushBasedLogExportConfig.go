// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pushbasedlogexport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PushBasedLogExportConfig struct {
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
	// The name of the bucket to which the agent sends the logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/push_based_log_export#bucket_name PushBasedLogExport#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// ID of the AWS IAM role that is used to write to the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/push_based_log_export#iam_role_id PushBasedLogExport#iam_role_id}
	IamRoleId *string `field:"required" json:"iamRoleId" yaml:"iamRoleId"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/push_based_log_export#project_id PushBasedLogExport#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// S3 directory in which vector writes in order to store the logs. An empty string denotes the root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/push_based_log_export#prefix_path PushBasedLogExport#prefix_path}
	PrefixPath *string `field:"optional" json:"prefixPath" yaml:"prefixPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/push_based_log_export#timeouts PushBasedLogExport#timeouts}.
	Timeouts *PushBasedLogExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

