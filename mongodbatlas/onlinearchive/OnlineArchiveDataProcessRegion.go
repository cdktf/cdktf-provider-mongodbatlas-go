// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinearchive


type OnlineArchiveDataProcessRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.23.0/docs/resources/online_archive#cloud_provider OnlineArchive#cloud_provider}.
	CloudProvider *string `field:"optional" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.23.0/docs/resources/online_archive#region OnlineArchive#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

