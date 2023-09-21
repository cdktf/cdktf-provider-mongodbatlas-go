// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package thirdpartyintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ThirdPartyIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#project_id ThirdPartyIntegration#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#type ThirdPartyIntegration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#api_key ThirdPartyIntegration#api_key}.
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#channel_name ThirdPartyIntegration#channel_name}.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#enabled ThirdPartyIntegration#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#id ThirdPartyIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#microsoft_teams_webhook_url ThirdPartyIntegration#microsoft_teams_webhook_url}.
	MicrosoftTeamsWebhookUrl *string `field:"optional" json:"microsoftTeamsWebhookUrl" yaml:"microsoftTeamsWebhookUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#password ThirdPartyIntegration#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#region ThirdPartyIntegration#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#routing_key ThirdPartyIntegration#routing_key}.
	RoutingKey *string `field:"optional" json:"routingKey" yaml:"routingKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#scheme ThirdPartyIntegration#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#secret ThirdPartyIntegration#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#service_discovery ThirdPartyIntegration#service_discovery}.
	ServiceDiscovery *string `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#service_key ThirdPartyIntegration#service_key}.
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#team_name ThirdPartyIntegration#team_name}.
	TeamName *string `field:"optional" json:"teamName" yaml:"teamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#url ThirdPartyIntegration#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.0/docs/resources/third_party_integration#user_name ThirdPartyIntegration#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

