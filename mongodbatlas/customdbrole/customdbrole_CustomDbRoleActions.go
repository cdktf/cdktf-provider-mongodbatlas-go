package customdbrole


type CustomDbRoleActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#action CustomDbRole#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#resources CustomDbRole#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

