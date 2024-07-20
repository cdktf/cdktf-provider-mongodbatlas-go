// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serverlessinstance


type ServerlessInstanceLinks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.4/docs/resources/serverless_instance#href ServerlessInstance#href}.
	Href *string `field:"optional" json:"href" yaml:"href"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.4/docs/resources/serverless_instance#rel ServerlessInstance#rel}.
	Rel *string `field:"optional" json:"rel" yaml:"rel"`
}

