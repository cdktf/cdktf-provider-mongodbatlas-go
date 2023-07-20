package federateddatabaseinstance


type FederatedDatabaseInstanceStorageDatabasesCollections struct {
	// data_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#data_sources FederatedDatabaseInstance#data_sources}
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#name FederatedDatabaseInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

