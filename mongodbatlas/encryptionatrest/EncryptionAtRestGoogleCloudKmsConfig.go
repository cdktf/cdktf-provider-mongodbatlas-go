// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest


type EncryptionAtRestGoogleCloudKmsConfig struct {
	// Flag that indicates whether someone enabled encryption at rest for the specified  project.
	//
	// To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/encryption_at_rest#enabled EncryptionAtRest#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Resource path that displays the key version resource ID for your Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/encryption_at_rest#key_version_resource_id EncryptionAtRest#key_version_resource_id}
	KeyVersionResourceId *string `field:"optional" json:"keyVersionResourceId" yaml:"keyVersionResourceId"`
	// Unique 24-hexadecimal digit string that identifies the Google Cloud Provider Access Role that MongoDB Cloud uses to access the Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/encryption_at_rest#role_id EncryptionAtRest#role_id}
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// JavaScript Object Notation (JSON) object that contains the Google Cloud Key Management Service (KMS).
	//
	// Format the JSON as a string and not as an object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/encryption_at_rest#service_account_key EncryptionAtRest#service_account_key}
	ServiceAccountKey *string `field:"optional" json:"serviceAccountKey" yaml:"serviceAccountKey"`
}

