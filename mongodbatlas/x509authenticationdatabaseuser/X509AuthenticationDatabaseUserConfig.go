// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package x509authenticationdatabaseuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type X509AuthenticationDatabaseUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/x509_authentication_database_user#project_id X509AuthenticationDatabaseUser#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/x509_authentication_database_user#customer_x509_cas X509AuthenticationDatabaseUser#customer_x509_cas}.
	CustomerX509Cas *string `field:"optional" json:"customerX509Cas" yaml:"customerX509Cas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/x509_authentication_database_user#id X509AuthenticationDatabaseUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/x509_authentication_database_user#months_until_expiration X509AuthenticationDatabaseUser#months_until_expiration}.
	MonthsUntilExpiration *float64 `field:"optional" json:"monthsUntilExpiration" yaml:"monthsUntilExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/x509_authentication_database_user#username X509AuthenticationDatabaseUser#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

