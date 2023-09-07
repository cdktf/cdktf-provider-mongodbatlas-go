// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectipaccesslist


type ProjectIpAccessListTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/project_ip_access_list#delete ProjectIpAccessList#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/project_ip_access_list#read ProjectIpAccessList#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

