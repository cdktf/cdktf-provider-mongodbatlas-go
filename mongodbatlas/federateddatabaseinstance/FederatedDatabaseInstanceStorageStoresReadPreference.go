package federateddatabaseinstance


type FederatedDatabaseInstanceStorageStoresReadPreference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#max_staleness_seconds FederatedDatabaseInstance#max_staleness_seconds}.
	MaxStalenessSeconds *float64 `field:"optional" json:"maxStalenessSeconds" yaml:"maxStalenessSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/federated_database_instance#mode FederatedDatabaseInstance#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

