package cluster


type ClusterLabels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/cluster#key Cluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/cluster#value Cluster#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

