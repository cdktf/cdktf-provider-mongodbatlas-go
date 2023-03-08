package cluster


type ClusterBiConnectorConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#enabled Cluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#read_preference Cluster#read_preference}.
	ReadPreference *string `field:"optional" json:"readPreference" yaml:"readPreference"`
}

