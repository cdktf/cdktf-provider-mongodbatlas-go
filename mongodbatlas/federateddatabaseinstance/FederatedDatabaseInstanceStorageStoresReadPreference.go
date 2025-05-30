// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance


type FederatedDatabaseInstanceStorageStoresReadPreference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#max_staleness_seconds FederatedDatabaseInstance#max_staleness_seconds}.
	MaxStalenessSeconds *float64 `field:"optional" json:"maxStalenessSeconds" yaml:"maxStalenessSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#mode FederatedDatabaseInstance#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// tag_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#tag_sets FederatedDatabaseInstance#tag_sets}
	TagSets interface{} `field:"optional" json:"tagSets" yaml:"tagSets"`
}

