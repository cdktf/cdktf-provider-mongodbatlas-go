package federateddatabaseinstance


type FederatedDatabaseInstanceCloudProviderConfig struct {
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/federated_database_instance#aws FederatedDatabaseInstance#aws}
	Aws *FederatedDatabaseInstanceCloudProviderConfigAws `field:"required" json:"aws" yaml:"aws"`
}

