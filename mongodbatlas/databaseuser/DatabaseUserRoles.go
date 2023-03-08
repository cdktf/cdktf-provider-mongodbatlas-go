package databaseuser


type DatabaseUserRoles struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#collection_name DatabaseUser#collection_name}.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#database_name DatabaseUser#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#role_name DatabaseUser#role_name}.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

