// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamconnection


type StreamConnectionSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.4/docs/resources/stream_connection#broker_public_certificate StreamConnection#broker_public_certificate}.
	BrokerPublicCertificate *string `field:"optional" json:"brokerPublicCertificate" yaml:"brokerPublicCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.4/docs/resources/stream_connection#protocol StreamConnection#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

