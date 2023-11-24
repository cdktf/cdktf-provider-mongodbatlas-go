// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalakepipeline


type DataLakePipelineTransformations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/data_lake_pipeline#field DataLakePipeline#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/data_lake_pipeline#type DataLakePipeline#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

