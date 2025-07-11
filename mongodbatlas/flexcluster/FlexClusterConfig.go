// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flexcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FlexClusterConfig struct {
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
	// Human-readable label that identifies the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#name FlexCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Unique 24-hexadecimal character string that identifies the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#project_id FlexCluster#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Group of cloud provider settings that configure the provisioned MongoDB flex cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#provider_settings FlexCluster#provider_settings}
	ProviderSettings *FlexClusterProviderSettings `field:"required" json:"providerSettings" yaml:"providerSettings"`
	// Map that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#tags FlexCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Flag that indicates whether termination protection is enabled on the cluster.
	//
	// If set to `true`, MongoDB Cloud won't delete the cluster. If set to `false`, MongoDB Cloud will delete the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/flex_cluster#termination_protection_enabled FlexCluster#termination_protection_enabled}
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

