// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPeeringConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#container_id NetworkPeering#container_id}.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#project_id NetworkPeering#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#provider_name NetworkPeering#provider_name}.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#accepter_region_name NetworkPeering#accepter_region_name}.
	AccepterRegionName *string `field:"optional" json:"accepterRegionName" yaml:"accepterRegionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#atlas_cidr_block NetworkPeering#atlas_cidr_block}.
	AtlasCidrBlock *string `field:"optional" json:"atlasCidrBlock" yaml:"atlasCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#atlas_gcp_project_id NetworkPeering#atlas_gcp_project_id}.
	AtlasGcpProjectId *string `field:"optional" json:"atlasGcpProjectId" yaml:"atlasGcpProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#atlas_vpc_name NetworkPeering#atlas_vpc_name}.
	AtlasVpcName *string `field:"optional" json:"atlasVpcName" yaml:"atlasVpcName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#aws_account_id NetworkPeering#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#azure_directory_id NetworkPeering#azure_directory_id}.
	AzureDirectoryId *string `field:"optional" json:"azureDirectoryId" yaml:"azureDirectoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#azure_subscription_id NetworkPeering#azure_subscription_id}.
	AzureSubscriptionId *string `field:"optional" json:"azureSubscriptionId" yaml:"azureSubscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#gcp_project_id NetworkPeering#gcp_project_id}.
	GcpProjectId *string `field:"optional" json:"gcpProjectId" yaml:"gcpProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#id NetworkPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#network_name NetworkPeering#network_name}.
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#resource_group_name NetworkPeering#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#route_table_cidr_block NetworkPeering#route_table_cidr_block}.
	RouteTableCidrBlock *string `field:"optional" json:"routeTableCidrBlock" yaml:"routeTableCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#vnet_name NetworkPeering#vnet_name}.
	VnetName *string `field:"optional" json:"vnetName" yaml:"vnetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/network_peering#vpc_id NetworkPeering#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

