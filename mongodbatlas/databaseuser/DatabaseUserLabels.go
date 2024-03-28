// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseuser


type DatabaseUserLabels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.3/docs/resources/database_user#key DatabaseUser#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.3/docs/resources/database_user#value DatabaseUser#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

