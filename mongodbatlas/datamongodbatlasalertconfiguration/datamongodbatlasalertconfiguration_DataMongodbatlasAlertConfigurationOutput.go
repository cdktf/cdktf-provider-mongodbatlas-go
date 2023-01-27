package datamongodbatlasalertconfiguration


type DataMongodbatlasAlertConfigurationOutput struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/alert_configuration#type DataMongodbatlasAlertConfiguration#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/alert_configuration#label DataMongodbatlasAlertConfiguration#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

