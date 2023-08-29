// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalake


type DataLakeDataProcessRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/data_lake#cloud_provider DataLake#cloud_provider}.
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/data_lake#region DataLake#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

