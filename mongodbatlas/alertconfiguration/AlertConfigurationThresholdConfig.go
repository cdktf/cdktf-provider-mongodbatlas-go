// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertconfiguration


type AlertConfigurationThresholdConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.2/docs/resources/alert_configuration#operator AlertConfiguration#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.2/docs/resources/alert_configuration#threshold AlertConfiguration#threshold}.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.2/docs/resources/alert_configuration#units AlertConfiguration#units}.
	Units *string `field:"optional" json:"units" yaml:"units"`
}

