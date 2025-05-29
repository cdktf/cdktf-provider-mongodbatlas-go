// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamprocessor


type StreamProcessorOptionsDlq struct {
	// Name of the collection to use for the DLQ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/stream_processor#coll StreamProcessor#coll}
	Coll *string `field:"required" json:"coll" yaml:"coll"`
	// Name of the connection to write DLQ messages to. Must be an Atlas connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/stream_processor#connection_name StreamProcessor#connection_name}
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Name of the database to use for the DLQ.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/stream_processor#db StreamProcessor#db}
	Db *string `field:"required" json:"db" yaml:"db"`
}

