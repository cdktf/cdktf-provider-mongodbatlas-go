package customdbrole


type CustomDbRoleInheritedRoles struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#database_name CustomDbRole#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/custom_db_role#role_name CustomDbRole#role_name}.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

