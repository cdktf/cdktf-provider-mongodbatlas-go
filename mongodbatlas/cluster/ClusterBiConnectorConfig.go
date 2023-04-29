package cluster


type ClusterBiConnectorConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/cluster#enabled Cluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/cluster#read_preference Cluster#read_preference}.
	ReadPreference *string `field:"optional" json:"readPreference" yaml:"readPreference"`
}
