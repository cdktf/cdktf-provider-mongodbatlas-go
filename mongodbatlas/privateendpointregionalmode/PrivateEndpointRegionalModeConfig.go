// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privateendpointregionalmode

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivateEndpointRegionalModeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/private_endpoint_regional_mode#project_id PrivateEndpointRegionalMode#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/private_endpoint_regional_mode#enabled PrivateEndpointRegionalMode#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/private_endpoint_regional_mode#id PrivateEndpointRegionalMode#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/private_endpoint_regional_mode#timeouts PrivateEndpointRegionalMode#timeouts}
	Timeouts *PrivateEndpointRegionalModeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

