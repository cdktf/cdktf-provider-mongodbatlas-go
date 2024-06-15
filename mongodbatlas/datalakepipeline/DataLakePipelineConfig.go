// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalakepipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLakePipelineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/data_lake_pipeline#name DataLakePipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/data_lake_pipeline#project_id DataLakePipeline#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// sink block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/data_lake_pipeline#sink DataLakePipeline#sink}
	Sink *DataLakePipelineSink `field:"optional" json:"sink" yaml:"sink"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/data_lake_pipeline#source DataLakePipeline#source}
	Source *DataLakePipelineSource `field:"optional" json:"source" yaml:"source"`
	// transformations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.1/docs/resources/data_lake_pipeline#transformations DataLakePipeline#transformations}
	Transformations interface{} `field:"optional" json:"transformations" yaml:"transformations"`
}

