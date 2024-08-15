// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#event_type AlertConfiguration#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#project_id AlertConfiguration#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#enabled AlertConfiguration#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#matcher AlertConfiguration#matcher}
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// metric_threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#metric_threshold_config AlertConfiguration#metric_threshold_config}
	MetricThresholdConfig interface{} `field:"optional" json:"metricThresholdConfig" yaml:"metricThresholdConfig"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#notification AlertConfiguration#notification}
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/alert_configuration#threshold_config AlertConfiguration#threshold_config}
	ThresholdConfig interface{} `field:"optional" json:"thresholdConfig" yaml:"thresholdConfig"`
}

