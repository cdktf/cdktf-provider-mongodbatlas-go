// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertconfiguration


type AlertConfigurationMatcher struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#field_name AlertConfiguration#field_name}.
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#operator AlertConfiguration#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#value AlertConfiguration#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

