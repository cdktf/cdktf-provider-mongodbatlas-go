// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamconnection


type StreamConnectionAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/stream_connection#mechanism StreamConnection#mechanism}.
	Mechanism *string `field:"optional" json:"mechanism" yaml:"mechanism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/stream_connection#password StreamConnection#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.4/docs/resources/stream_connection#username StreamConnection#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

