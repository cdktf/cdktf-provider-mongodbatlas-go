// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamconnection


type StreamConnectionAws struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.31.0/docs/resources/stream_connection#role_arn StreamConnection#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

