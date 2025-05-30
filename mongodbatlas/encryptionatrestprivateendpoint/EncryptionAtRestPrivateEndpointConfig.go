// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrestprivateendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EncryptionAtRestPrivateEndpointConfig struct {
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
	// Label that identifies the cloud provider for the Encryption At Rest private endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/encryption_at_rest_private_endpoint#cloud_provider EncryptionAtRestPrivateEndpoint#cloud_provider}
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/encryption_at_rest_private_endpoint#project_id EncryptionAtRestPrivateEndpoint#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Cloud provider region in which the Encryption At Rest private endpoint is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/encryption_at_rest_private_endpoint#region_name EncryptionAtRestPrivateEndpoint#region_name}
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

