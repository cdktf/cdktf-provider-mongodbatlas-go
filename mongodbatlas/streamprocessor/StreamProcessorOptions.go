// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamprocessor


type StreamProcessorOptions struct {
	// Dead letter queue for the stream processor. Refer to the [MongoDB Atlas Docs](https://www.mongodb.com/docs/atlas/reference/glossary/#std-term-dead-letter-queue) for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/stream_processor#dlq StreamProcessor#dlq}
	Dlq *StreamProcessorOptionsDlq `field:"required" json:"dlq" yaml:"dlq"`
}

