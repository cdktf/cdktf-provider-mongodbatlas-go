package alertconfiguration


type AlertConfigurationThresholdConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/alert_configuration#operator AlertConfiguration#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/alert_configuration#threshold AlertConfiguration#threshold}.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/alert_configuration#units AlertConfiguration#units}.
	Units *string `field:"optional" json:"units" yaml:"units"`
}

