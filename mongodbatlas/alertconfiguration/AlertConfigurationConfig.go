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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#event_type AlertConfiguration#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#notification AlertConfiguration#notification}
	Notification interface{} `field:"required" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#project_id AlertConfiguration#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#enabled AlertConfiguration#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#id AlertConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#matcher AlertConfiguration#matcher}
	Matcher interface{} `field:"optional" json:"matcher" yaml:"matcher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#metric_threshold AlertConfiguration#metric_threshold}.
	MetricThreshold *map[string]*string `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
	// metric_threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#metric_threshold_config AlertConfiguration#metric_threshold_config}
	MetricThresholdConfig *AlertConfigurationMetricThresholdConfig `field:"optional" json:"metricThresholdConfig" yaml:"metricThresholdConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#threshold AlertConfiguration#threshold}.
	Threshold *map[string]*string `field:"optional" json:"threshold" yaml:"threshold"`
	// threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/alert_configuration#threshold_config AlertConfiguration#threshold_config}
	ThresholdConfig *AlertConfigurationThresholdConfig `field:"optional" json:"thresholdConfig" yaml:"thresholdConfig"`
}

