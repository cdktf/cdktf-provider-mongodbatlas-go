// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamprivatelinkendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamPrivatelinkEndpointConfig struct {
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
	// Unique 24-hexadecimal digit string that identifies your project.
	//
	// Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.
	//
	// **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group or project id remains the same. The resource and corresponding endpoints use the term groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#project_id StreamPrivatelinkEndpoint#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Provider where the Kafka cluster is deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#provider_name StreamPrivatelinkEndpoint#provider_name}
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Vendor who manages the Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#vendor StreamPrivatelinkEndpoint#vendor}
	Vendor *string `field:"required" json:"vendor" yaml:"vendor"`
	// Domain name of Privatelink connected cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#dns_domain StreamPrivatelinkEndpoint#dns_domain}
	DnsDomain *string `field:"optional" json:"dnsDomain" yaml:"dnsDomain"`
	// Sub-Domain name of Confluent cluster. These are typically your availability zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#dns_sub_domain StreamPrivatelinkEndpoint#dns_sub_domain}
	DnsSubDomain *[]*string `field:"optional" json:"dnsSubDomain" yaml:"dnsSubDomain"`
	// Domain name of Confluent cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#region StreamPrivatelinkEndpoint#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Service Endpoint ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/stream_privatelink_endpoint#service_endpoint_id StreamPrivatelinkEndpoint#service_endpoint_id}
	ServiceEndpointId *string `field:"optional" json:"serviceEndpointId" yaml:"serviceEndpointId"`
}

