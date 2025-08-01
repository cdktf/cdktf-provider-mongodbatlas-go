// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FederatedDatabaseInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#name FederatedDatabaseInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#project_id FederatedDatabaseInstance#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// cloud_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#cloud_provider_config FederatedDatabaseInstance#cloud_provider_config}
	CloudProviderConfig *FederatedDatabaseInstanceCloudProviderConfig `field:"optional" json:"cloudProviderConfig" yaml:"cloudProviderConfig"`
	// data_process_region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#data_process_region FederatedDatabaseInstance#data_process_region}
	DataProcessRegion *FederatedDatabaseInstanceDataProcessRegion `field:"optional" json:"dataProcessRegion" yaml:"dataProcessRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#id FederatedDatabaseInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// storage_databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#storage_databases FederatedDatabaseInstance#storage_databases}
	StorageDatabases interface{} `field:"optional" json:"storageDatabases" yaml:"storageDatabases"`
	// storage_stores block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/federated_database_instance#storage_stores FederatedDatabaseInstance#storage_stores}
	StorageStores interface{} `field:"optional" json:"storageStores" yaml:"storageStores"`
}

