// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mongodbemployeeaccessgrant

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MongodbEmployeeAccessGrantConfig struct {
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
	// Human-readable label that identifies this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/mongodb_employee_access_grant#cluster_name MongodbEmployeeAccessGrant#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Expiration date for the employee access grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/mongodb_employee_access_grant#expiration_time MongodbEmployeeAccessGrant#expiration_time}
	ExpirationTime *string `field:"required" json:"expirationTime" yaml:"expirationTime"`
	// Level of access to grant to MongoDB Employees. Possible values are CLUSTER_DATABASE_LOGS, CLUSTER_INFRASTRUCTURE or CLUSTER_INFRASTRUCTURE_AND_APP_SERVICES_SYNC_DATA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/mongodb_employee_access_grant#grant_type MongodbEmployeeAccessGrant#grant_type}
	GrantType *string `field:"required" json:"grantType" yaml:"grantType"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/mongodb_employee_access_grant#project_id MongodbEmployeeAccessGrant#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

