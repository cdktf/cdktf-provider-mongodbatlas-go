// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasalertconfigurations


type DataMongodbatlasAlertConfigurationsResults struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/data-sources/alert_configurations#alert_configuration_id DataMongodbatlasAlertConfigurations#alert_configuration_id}.
	AlertConfigurationId *string `field:"required" json:"alertConfigurationId" yaml:"alertConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.41.1/docs/data-sources/alert_configurations#project_id DataMongodbatlasAlertConfigurations#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

