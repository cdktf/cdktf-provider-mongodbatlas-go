// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchdeployment


type SearchDeploymentSpecs struct {
	// Hardware specification for the search node instance sizes.
	//
	// The [MongoDB Atlas API](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/#tag/Atlas-Search/operation/createAtlasSearchDeployment) describes the valid values. More details can also be found in the [Search Node Documentation](https://www.mongodb.com/docs/atlas/cluster-config/multi-cloud-distribution/#search-tier).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/resources/search_deployment#instance_size SearchDeployment#instance_size}
	InstanceSize *string `field:"required" json:"instanceSize" yaml:"instanceSize"`
	// Number of search nodes in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/resources/search_deployment#node_count SearchDeployment#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
}

