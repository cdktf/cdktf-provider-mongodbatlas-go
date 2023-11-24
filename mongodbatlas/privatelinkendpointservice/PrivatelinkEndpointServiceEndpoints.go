// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatelinkendpointservice


type PrivatelinkEndpointServiceEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/privatelink_endpoint_service#endpoint_name PrivatelinkEndpointService#endpoint_name}.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/privatelink_endpoint_service#ip_address PrivatelinkEndpointService#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/privatelink_endpoint_service#service_attachment_name PrivatelinkEndpointService#service_attachment_name}.
	ServiceAttachmentName *string `field:"optional" json:"serviceAttachmentName" yaml:"serviceAttachmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/privatelink_endpoint_service#status PrivatelinkEndpointService#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

