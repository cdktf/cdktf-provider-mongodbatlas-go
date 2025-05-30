// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance


type FederatedDatabaseInstanceStorageStores struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#additional_storage_classes FederatedDatabaseInstance#additional_storage_classes}.
	AdditionalStorageClasses *[]*string `field:"optional" json:"additionalStorageClasses" yaml:"additionalStorageClasses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#allow_insecure FederatedDatabaseInstance#allow_insecure}.
	AllowInsecure interface{} `field:"optional" json:"allowInsecure" yaml:"allowInsecure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#bucket FederatedDatabaseInstance#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#cluster_name FederatedDatabaseInstance#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#default_format FederatedDatabaseInstance#default_format}.
	DefaultFormat *string `field:"optional" json:"defaultFormat" yaml:"defaultFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#delimiter FederatedDatabaseInstance#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#include_tags FederatedDatabaseInstance#include_tags}.
	IncludeTags interface{} `field:"optional" json:"includeTags" yaml:"includeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#name FederatedDatabaseInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#prefix FederatedDatabaseInstance#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#project_id FederatedDatabaseInstance#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#provider FederatedDatabaseInstance#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#public FederatedDatabaseInstance#public}.
	Public *string `field:"optional" json:"public" yaml:"public"`
	// read_preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#read_preference FederatedDatabaseInstance#read_preference}
	ReadPreference *FederatedDatabaseInstanceStorageStoresReadPreference `field:"optional" json:"readPreference" yaml:"readPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#region FederatedDatabaseInstance#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/federated_database_instance#urls FederatedDatabaseInstance#urls}.
	Urls *[]*string `field:"optional" json:"urls" yaml:"urls"`
}

