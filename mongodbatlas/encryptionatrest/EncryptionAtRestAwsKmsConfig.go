// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest


type EncryptionAtRestAwsKmsConfig struct {
	// Unique alphanumeric string that identifies an Identity and Access Management (IAM) access key with permissions required to access your Amazon Web Services (AWS) Customer Master Key (CMK).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#access_key_id EncryptionAtRest#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Unique alphanumeric string that identifies the Amazon Web Services (AWS) Customer Master Key (CMK) you used to encrypt and decrypt the MongoDB master keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#customer_master_key_id EncryptionAtRest#customer_master_key_id}
	CustomerMasterKeyId *string `field:"optional" json:"customerMasterKeyId" yaml:"customerMasterKeyId"`
	// Flag that indicates whether someone enabled encryption at rest for the specified project through Amazon Web Services (AWS) Key Management Service (KMS).
	//
	// To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#enabled EncryptionAtRest#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Physical location where MongoDB Atlas deploys your AWS-hosted MongoDB cluster nodes.
	//
	// The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Atlas creates them as part of the deployment. MongoDB Atlas assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#region EncryptionAtRest#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Enable connection to your Amazon Web Services (AWS) Key Management Service (KMS) over private networking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#require_private_networking EncryptionAtRest#require_private_networking}
	RequirePrivateNetworking interface{} `field:"optional" json:"requirePrivateNetworking" yaml:"requirePrivateNetworking"`
	// Unique 24-hexadecimal digit string that identifies an Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	//
	// This IAM role has the permissions required to manage your AWS customer master key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#role_id EncryptionAtRest#role_id}
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Human-readable label of the Identity and Access Management (IAM) secret access key with permissions required to access your Amazon Web Services (AWS) customer master key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.33.0/docs/resources/encryption_at_rest#secret_access_key EncryptionAtRest#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

