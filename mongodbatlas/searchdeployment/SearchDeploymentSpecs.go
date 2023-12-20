// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchdeployment


type SearchDeploymentSpecs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/search_deployment#instance_size SearchDeployment#instance_size}.
	InstanceSize *string `field:"required" json:"instanceSize" yaml:"instanceSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/search_deployment#node_count SearchDeployment#node_count}.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
}

