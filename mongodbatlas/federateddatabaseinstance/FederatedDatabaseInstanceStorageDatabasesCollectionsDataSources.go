// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance


type FederatedDatabaseInstanceStorageDatabasesCollectionsDataSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#allow_insecure FederatedDatabaseInstance#allow_insecure}.
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#collection FederatedDatabaseInstance#collection}.
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#collection_regex FederatedDatabaseInstance#collection_regex}.
	CollectionRegex *string `field:"optional" json:"collectionRegex" yaml:"collectionRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#database FederatedDatabaseInstance#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#database_regex FederatedDatabaseInstance#database_regex}.
	DatabaseRegex *string `field:"optional" json:"databaseRegex" yaml:"databaseRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#dataset_name FederatedDatabaseInstance#dataset_name}.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#default_format FederatedDatabaseInstance#default_format}.
	DefaultFormat *string `field:"optional" json:"defaultFormat" yaml:"defaultFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#path FederatedDatabaseInstance#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#provenance_field_name FederatedDatabaseInstance#provenance_field_name}.
	ProvenanceFieldName *string `field:"optional" json:"provenanceFieldName" yaml:"provenanceFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#store_name FederatedDatabaseInstance#store_name}.
	StoreName *string `field:"optional" json:"storeName" yaml:"storeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.40.0/docs/resources/federated_database_instance#urls FederatedDatabaseInstance#urls}.
	Urls *[]*string `field:"optional" json:"urls" yaml:"urls"`
}

