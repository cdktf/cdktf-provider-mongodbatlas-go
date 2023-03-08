package advancedcluster


type AdvancedClusterBiConnector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#enabled AdvancedCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#read_preference AdvancedCluster#read_preference}.
	ReadPreference *string `field:"optional" json:"readPreference" yaml:"readPreference"`
}

