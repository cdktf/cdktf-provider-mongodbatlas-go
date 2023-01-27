package databaseuser


type DatabaseUserLabels struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#key DatabaseUser#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#value DatabaseUser#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

