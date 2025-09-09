// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streaminstance


type StreamInstanceStreamConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/stream_instance#tier StreamInstance#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

