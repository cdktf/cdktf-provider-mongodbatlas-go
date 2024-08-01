// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest


type EncryptionAtRestAwsKmsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/encryption_at_rest#access_key_id EncryptionAtRest#access_key_id}.
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/encryption_at_rest#customer_master_key_id EncryptionAtRest#customer_master_key_id}.
	CustomerMasterKeyId *string `field:"optional" json:"customerMasterKeyId" yaml:"customerMasterKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/encryption_at_rest#enabled EncryptionAtRest#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/encryption_at_rest#region EncryptionAtRest#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/encryption_at_rest#role_id EncryptionAtRest#role_id}.
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/encryption_at_rest#secret_access_key EncryptionAtRest#secret_access_key}.
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

