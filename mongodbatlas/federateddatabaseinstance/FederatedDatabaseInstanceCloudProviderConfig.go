// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance


type FederatedDatabaseInstanceCloudProviderConfig struct {
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/resources/federated_database_instance#aws FederatedDatabaseInstance#aws}
	Aws *FederatedDatabaseInstanceCloudProviderConfigAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/resources/federated_database_instance#azure FederatedDatabaseInstance#azure}
	Azure *FederatedDatabaseInstanceCloudProviderConfigAzure `field:"optional" json:"azure" yaml:"azure"`
}

