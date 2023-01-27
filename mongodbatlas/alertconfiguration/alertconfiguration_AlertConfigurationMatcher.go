package alertconfiguration


type AlertConfigurationMatcher struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/alert_configuration#field_name AlertConfiguration#field_name}.
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/alert_configuration#operator AlertConfiguration#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/alert_configuration#value AlertConfiguration#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

