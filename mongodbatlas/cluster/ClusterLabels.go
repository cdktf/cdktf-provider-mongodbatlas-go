package cluster


type ClusterLabels struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#key Cluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#value Cluster#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

