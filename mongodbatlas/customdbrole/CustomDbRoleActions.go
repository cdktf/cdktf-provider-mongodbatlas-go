package customdbrole


type CustomDbRoleActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/custom_db_role#action CustomDbRole#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/custom_db_role#resources CustomDbRole#resources}
	Resources interface{} `field:"required" json:"resources" yaml:"resources"`
}

