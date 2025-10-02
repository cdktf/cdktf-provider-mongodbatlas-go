// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatelinkendpointservice


type PrivatelinkEndpointServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/resources/privatelink_endpoint_service#create PrivatelinkEndpointService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/resources/privatelink_endpoint_service#delete PrivatelinkEndpointService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

