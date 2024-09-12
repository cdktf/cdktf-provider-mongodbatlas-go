// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertconfiguration


type AlertConfigurationMatcher struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/alert_configuration#field_name AlertConfiguration#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/alert_configuration#operator AlertConfiguration#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.19.0/docs/resources/alert_configuration#value AlertConfiguration#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

