// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasdatalakepipelinerun

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasDataLakePipelineRunConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/data-sources/data_lake_pipeline_run#pipeline_name DataMongodbatlasDataLakePipelineRun#pipeline_name}.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/data-sources/data_lake_pipeline_run#pipeline_run_id DataMongodbatlasDataLakePipelineRun#pipeline_run_id}.
	PipelineRunId *string `field:"required" json:"pipelineRunId" yaml:"pipelineRunId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.0/docs/data-sources/data_lake_pipeline_run#project_id DataMongodbatlasDataLakePipelineRun#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

