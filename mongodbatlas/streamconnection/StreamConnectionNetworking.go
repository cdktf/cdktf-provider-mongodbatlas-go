// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamconnection


type StreamConnectionNetworking struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.25.0/docs/resources/stream_connection#access StreamConnection#access}.
	Access *StreamConnectionNetworkingAccess `field:"required" json:"access" yaml:"access"`
}

