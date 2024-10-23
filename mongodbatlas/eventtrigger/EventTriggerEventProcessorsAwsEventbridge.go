// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventtrigger


type EventTriggerEventProcessorsAwsEventbridge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.2/docs/resources/event_trigger#config_account_id EventTrigger#config_account_id}.
	ConfigAccountId *string `field:"optional" json:"configAccountId" yaml:"configAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.2/docs/resources/event_trigger#config_region EventTrigger#config_region}.
	ConfigRegion *string `field:"optional" json:"configRegion" yaml:"configRegion"`
}

