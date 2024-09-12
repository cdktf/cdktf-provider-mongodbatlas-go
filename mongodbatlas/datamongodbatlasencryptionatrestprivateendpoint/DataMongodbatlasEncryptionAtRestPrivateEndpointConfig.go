// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasencryptionatrestprivateendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasEncryptionAtRestPrivateEndpointConfig struct {
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
	// Label that identifies the cloud provider of the private endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/data-sources/encryption_at_rest_private_endpoint#cloud_provider DataMongodbatlasEncryptionAtRestPrivateEndpoint#cloud_provider}
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/data-sources/encryption_at_rest_private_endpoint#id DataMongodbatlasEncryptionAtRestPrivateEndpoint#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/data-sources/encryption_at_rest_private_endpoint#project_id DataMongodbatlasEncryptionAtRestPrivateEndpoint#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}
