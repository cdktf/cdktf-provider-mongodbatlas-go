// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privateendpointregionalmode


type PrivateEndpointRegionalModeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.20.0/docs/resources/private_endpoint_regional_mode#create PrivateEndpointRegionalMode#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.20.0/docs/resources/private_endpoint_regional_mode#delete PrivateEndpointRegionalMode#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.20.0/docs/resources/private_endpoint_regional_mode#update PrivateEndpointRegionalMode#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

