// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertconfiguration


type AlertConfigurationNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#type_name AlertConfiguration#type_name}.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#api_token AlertConfiguration#api_token}.
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#channel_name AlertConfiguration#channel_name}.
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#datadog_api_key AlertConfiguration#datadog_api_key}.
	DatadogApiKey *string `field:"optional" json:"datadogApiKey" yaml:"datadogApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#datadog_region AlertConfiguration#datadog_region}.
	DatadogRegion *string `field:"optional" json:"datadogRegion" yaml:"datadogRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#delay_min AlertConfiguration#delay_min}.
	DelayMin *float64 `field:"optional" json:"delayMin" yaml:"delayMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#email_address AlertConfiguration#email_address}.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#email_enabled AlertConfiguration#email_enabled}.
	EmailEnabled interface{} `field:"optional" json:"emailEnabled" yaml:"emailEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#interval_min AlertConfiguration#interval_min}.
	IntervalMin *float64 `field:"optional" json:"intervalMin" yaml:"intervalMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#microsoft_teams_webhook_url AlertConfiguration#microsoft_teams_webhook_url}.
	MicrosoftTeamsWebhookUrl *string `field:"optional" json:"microsoftTeamsWebhookUrl" yaml:"microsoftTeamsWebhookUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#mobile_number AlertConfiguration#mobile_number}.
	MobileNumber *string `field:"optional" json:"mobileNumber" yaml:"mobileNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#notifier_id AlertConfiguration#notifier_id}.
	NotifierId *string `field:"optional" json:"notifierId" yaml:"notifierId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#ops_genie_api_key AlertConfiguration#ops_genie_api_key}.
	OpsGenieApiKey *string `field:"optional" json:"opsGenieApiKey" yaml:"opsGenieApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#ops_genie_region AlertConfiguration#ops_genie_region}.
	OpsGenieRegion *string `field:"optional" json:"opsGenieRegion" yaml:"opsGenieRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#roles AlertConfiguration#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#service_key AlertConfiguration#service_key}.
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#sms_enabled AlertConfiguration#sms_enabled}.
	SmsEnabled interface{} `field:"optional" json:"smsEnabled" yaml:"smsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#team_id AlertConfiguration#team_id}.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#username AlertConfiguration#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#victor_ops_api_key AlertConfiguration#victor_ops_api_key}.
	VictorOpsApiKey *string `field:"optional" json:"victorOpsApiKey" yaml:"victorOpsApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#victor_ops_routing_key AlertConfiguration#victor_ops_routing_key}.
	VictorOpsRoutingKey *string `field:"optional" json:"victorOpsRoutingKey" yaml:"victorOpsRoutingKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#webhook_secret AlertConfiguration#webhook_secret}.
	WebhookSecret *string `field:"optional" json:"webhookSecret" yaml:"webhookSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/alert_configuration#webhook_url AlertConfiguration#webhook_url}.
	WebhookUrl *string `field:"optional" json:"webhookUrl" yaml:"webhookUrl"`
}

