package globalclusterconfig


type GlobalClusterConfigCustomZoneMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#location GlobalClusterConfig#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#zone GlobalClusterConfig#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

