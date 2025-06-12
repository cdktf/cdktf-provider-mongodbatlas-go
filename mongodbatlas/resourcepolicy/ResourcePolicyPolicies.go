// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcepolicy


type ResourcePolicyPolicies struct {
	// A string that defines the permissions for the policy. The syntax used is the Cedar Policy language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/resource_policy#body ResourcePolicy#body}
	Body *string `field:"required" json:"body" yaml:"body"`
}

