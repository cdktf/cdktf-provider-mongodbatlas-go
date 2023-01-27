package databaseuser


type DatabaseUserScopes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#name DatabaseUser#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#type DatabaseUser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

