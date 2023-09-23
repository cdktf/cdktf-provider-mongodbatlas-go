// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalake


type DataLakeAws struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/data_lake#role_id DataLake#role_id}.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/data_lake#test_s3_bucket DataLake#test_s3_bucket}.
	TestS3Bucket *string `field:"required" json:"testS3Bucket" yaml:"testS3Bucket"`
}

