// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project


type ProjectTeams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/project#role_names Project#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/project#team_id Project#team_id}.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

