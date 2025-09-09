// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#auth_database_name DatabaseUser#auth_database_name}.
	AuthDatabaseName *string `field:"required" json:"authDatabaseName" yaml:"authDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#project_id DatabaseUser#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#username DatabaseUser#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#aws_iam_type DatabaseUser#aws_iam_type}.
	AwsIamType *string `field:"optional" json:"awsIamType" yaml:"awsIamType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#description DatabaseUser#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#labels DatabaseUser#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#ldap_auth_type DatabaseUser#ldap_auth_type}.
	LdapAuthType *string `field:"optional" json:"ldapAuthType" yaml:"ldapAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#oidc_auth_type DatabaseUser#oidc_auth_type}.
	OidcAuthType *string `field:"optional" json:"oidcAuthType" yaml:"oidcAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#password DatabaseUser#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#roles DatabaseUser#roles}
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// scopes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#scopes DatabaseUser#scopes}
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/database_user#x509_type DatabaseUser#x509_type}.
	X509Type *string `field:"optional" json:"x509Type" yaml:"x509Type"`
}

