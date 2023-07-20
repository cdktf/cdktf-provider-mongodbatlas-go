package globalclusterconfig


type GlobalClusterConfigCustomZoneMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/global_cluster_config#location GlobalClusterConfig#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/global_cluster_config#zone GlobalClusterConfig#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

