// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamconnection


type StreamConnectionNetworkingAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.25.0/docs/resources/stream_connection#type StreamConnection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

