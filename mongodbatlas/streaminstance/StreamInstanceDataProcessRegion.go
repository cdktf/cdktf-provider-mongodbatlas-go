// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streaminstance


type StreamInstanceDataProcessRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.2/docs/resources/stream_instance#cloud_provider StreamInstance#cloud_provider}.
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.2/docs/resources/stream_instance#region StreamInstance#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

