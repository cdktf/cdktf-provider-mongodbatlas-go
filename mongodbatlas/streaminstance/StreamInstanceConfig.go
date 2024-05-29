// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streaminstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.1/docs/resources/stream_instance#data_process_region StreamInstance#data_process_region}.
	DataProcessRegion *StreamInstanceDataProcessRegion `field:"required" json:"dataProcessRegion" yaml:"dataProcessRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.1/docs/resources/stream_instance#instance_name StreamInstance#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.1/docs/resources/stream_instance#project_id StreamInstance#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.1/docs/resources/stream_instance#stream_config StreamInstance#stream_config}.
	StreamConfig *StreamInstanceStreamConfig `field:"optional" json:"streamConfig" yaml:"streamConfig"`
}

