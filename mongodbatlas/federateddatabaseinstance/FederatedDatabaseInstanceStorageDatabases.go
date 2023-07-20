package federateddatabaseinstance


type FederatedDatabaseInstanceStorageDatabases struct {
	// collections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#collections FederatedDatabaseInstance#collections}
	Collections interface{} `field:"optional" json:"collections" yaml:"collections"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#name FederatedDatabaseInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// views block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#views FederatedDatabaseInstance#views}
	Views interface{} `field:"optional" json:"views" yaml:"views"`
}
