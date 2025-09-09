// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance


type FederatedDatabaseInstanceCloudProviderConfigAzure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/resources/federated_database_instance#role_id FederatedDatabaseInstance#role_id}.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
}

