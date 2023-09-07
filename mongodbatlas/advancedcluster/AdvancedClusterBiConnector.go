// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterBiConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/advanced_cluster#enabled AdvancedCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/advanced_cluster#read_preference AdvancedCluster#read_preference}.
	ReadPreference *string `field:"optional" json:"readPreference" yaml:"readPreference"`
}

