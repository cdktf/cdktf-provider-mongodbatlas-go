package customdbrole


type CustomDbRoleActionsResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#cluster CustomDbRole#cluster}.
	Cluster interface{} `field:"optional" json:"cluster" yaml:"cluster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#collection_name CustomDbRole#collection_name}.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#database_name CustomDbRole#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}

