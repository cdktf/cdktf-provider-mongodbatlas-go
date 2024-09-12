// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchindex


type SearchIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/search_index#create SearchIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/search_index#delete SearchIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/search_index#update SearchIndex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

