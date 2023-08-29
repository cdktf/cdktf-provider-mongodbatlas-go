// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventtrigger


type EventTriggerEventProcessors struct {
	// aws_eventbridge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/event_trigger#aws_eventbridge EventTrigger#aws_eventbridge}
	AwsEventbridge *EventTriggerEventProcessorsAwsEventbridge `field:"optional" json:"awsEventbridge" yaml:"awsEventbridge"`
}

