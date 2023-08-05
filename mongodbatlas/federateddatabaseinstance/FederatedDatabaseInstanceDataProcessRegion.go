package federateddatabaseinstance


type FederatedDatabaseInstanceDataProcessRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/federated_database_instance#cloud_provider FederatedDatabaseInstance#cloud_provider}.
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/federated_database_instance#region FederatedDatabaseInstance#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

