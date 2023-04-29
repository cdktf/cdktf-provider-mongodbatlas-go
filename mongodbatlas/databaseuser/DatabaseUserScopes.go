package databaseuser


type DatabaseUserScopes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/database_user#name DatabaseUser#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/database_user#type DatabaseUser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

