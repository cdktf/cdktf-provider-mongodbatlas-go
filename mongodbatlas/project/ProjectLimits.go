// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/project#name Project#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/project#value Project#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

