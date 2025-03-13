// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance


type FederatedDatabaseInstanceCloudProviderConfigAws struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/federated_database_instance#role_id FederatedDatabaseInstance#role_id}.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/federated_database_instance#test_s3_bucket FederatedDatabaseInstance#test_s3_bucket}.
	TestS3Bucket *string `field:"required" json:"testS3Bucket" yaml:"testS3Bucket"`
}

