// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EncryptionAtRestConfig struct {
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
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/encryption_at_rest#project_id EncryptionAtRest#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// aws_kms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/encryption_at_rest#aws_kms_config EncryptionAtRest#aws_kms_config}
	AwsKmsConfig interface{} `field:"optional" json:"awsKmsConfig" yaml:"awsKmsConfig"`
	// azure_key_vault_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/encryption_at_rest#azure_key_vault_config EncryptionAtRest#azure_key_vault_config}
	AzureKeyVaultConfig interface{} `field:"optional" json:"azureKeyVaultConfig" yaml:"azureKeyVaultConfig"`
	// Flag that indicates whether Encryption at Rest for Dedicated Search Nodes is enabled in the specified project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/encryption_at_rest#enabled_for_search_nodes EncryptionAtRest#enabled_for_search_nodes}
	EnabledForSearchNodes interface{} `field:"optional" json:"enabledForSearchNodes" yaml:"enabledForSearchNodes"`
	// google_cloud_kms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/encryption_at_rest#google_cloud_kms_config EncryptionAtRest#google_cloud_kms_config}
	GoogleCloudKmsConfig interface{} `field:"optional" json:"googleCloudKmsConfig" yaml:"googleCloudKmsConfig"`
}

