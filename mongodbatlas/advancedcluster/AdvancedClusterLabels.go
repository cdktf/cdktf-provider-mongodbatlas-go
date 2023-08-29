// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterLabels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/advanced_cluster#key AdvancedCluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/advanced_cluster#value AdvancedCluster#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

