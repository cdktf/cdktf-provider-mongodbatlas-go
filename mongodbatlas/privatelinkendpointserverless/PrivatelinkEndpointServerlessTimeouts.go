// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatelinkendpointserverless


type PrivatelinkEndpointServerlessTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/resources/privatelink_endpoint_serverless#create PrivatelinkEndpointServerless#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/resources/privatelink_endpoint_serverless#delete PrivatelinkEndpointServerless#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

