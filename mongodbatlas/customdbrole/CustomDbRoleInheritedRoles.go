// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customdbrole


type CustomDbRoleInheritedRoles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/custom_db_role#database_name CustomDbRole#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.22.0/docs/resources/custom_db_role#role_name CustomDbRole#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

