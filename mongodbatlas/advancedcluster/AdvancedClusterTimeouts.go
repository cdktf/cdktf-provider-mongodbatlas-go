package advancedcluster


type AdvancedClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#create AdvancedCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#delete AdvancedCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#update AdvancedCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

