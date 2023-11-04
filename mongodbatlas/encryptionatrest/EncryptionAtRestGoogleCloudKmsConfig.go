// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest


type EncryptionAtRestGoogleCloudKmsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.3/docs/resources/encryption_at_rest#enabled EncryptionAtRest#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.3/docs/resources/encryption_at_rest#key_version_resource_id EncryptionAtRest#key_version_resource_id}.
	KeyVersionResourceId *string `field:"optional" json:"keyVersionResourceId" yaml:"keyVersionResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.3/docs/resources/encryption_at_rest#service_account_key EncryptionAtRest#service_account_key}.
	ServiceAccountKey *string `field:"optional" json:"serviceAccountKey" yaml:"serviceAccountKey"`
}

