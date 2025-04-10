// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasalertconfiguration


type DataMongodbatlasAlertConfigurationOutput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/data-sources/alert_configuration#type DataMongodbatlasAlertConfiguration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/data-sources/alert_configuration#label DataMongodbatlasAlertConfiguration#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

