// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectApiKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/project#api_key_id Project#api_key_id}.
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/project#role_names Project#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
}

