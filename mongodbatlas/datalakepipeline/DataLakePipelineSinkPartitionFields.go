// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalakepipeline


type DataLakePipelineSinkPartitionFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/data_lake_pipeline#field_name DataLakePipeline#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/data_lake_pipeline#order DataLakePipeline#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

