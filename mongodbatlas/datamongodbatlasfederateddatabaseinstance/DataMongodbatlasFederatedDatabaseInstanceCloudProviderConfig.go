package datamongodbatlasfederateddatabaseinstance


type DataMongodbatlasFederatedDatabaseInstanceCloudProviderConfig struct {
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/federated_database_instance#aws DataMongodbatlasFederatedDatabaseInstance#aws}
	Aws *DataMongodbatlasFederatedDatabaseInstanceCloudProviderConfigAws `field:"optional" json:"aws" yaml:"aws"`
}
