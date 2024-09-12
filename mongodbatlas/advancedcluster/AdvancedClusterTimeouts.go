// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/advanced_cluster#create AdvancedCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/advanced_cluster#delete AdvancedCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/advanced_cluster#update AdvancedCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

