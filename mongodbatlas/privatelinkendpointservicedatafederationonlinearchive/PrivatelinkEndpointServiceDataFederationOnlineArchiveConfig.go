// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatelinkendpointservicedatafederationonlinearchive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatelinkEndpointServiceDataFederationOnlineArchiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#endpoint_id PrivatelinkEndpointServiceDataFederationOnlineArchive#endpoint_id}.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#project_id PrivatelinkEndpointServiceDataFederationOnlineArchive#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#provider_name PrivatelinkEndpointServiceDataFederationOnlineArchive#provider_name}.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#comment PrivatelinkEndpointServiceDataFederationOnlineArchive#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#customer_endpoint_dns_name PrivatelinkEndpointServiceDataFederationOnlineArchive#customer_endpoint_dns_name}.
	CustomerEndpointDnsName *string `field:"optional" json:"customerEndpointDnsName" yaml:"customerEndpointDnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#id PrivatelinkEndpointServiceDataFederationOnlineArchive#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#region PrivatelinkEndpointServiceDataFederationOnlineArchive#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/privatelink_endpoint_service_data_federation_online_archive#timeouts PrivatelinkEndpointServiceDataFederationOnlineArchive#timeouts}
	Timeouts *PrivatelinkEndpointServiceDataFederationOnlineArchiveTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

