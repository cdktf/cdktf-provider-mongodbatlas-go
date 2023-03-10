package advancedcluster


type AdvancedClusterReplicationSpecsRegionConfigsAutoScaling struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#compute_enabled AdvancedCluster#compute_enabled}.
	ComputeEnabled interface{} `field:"optional" json:"computeEnabled" yaml:"computeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#compute_max_instance_size AdvancedCluster#compute_max_instance_size}.
	ComputeMaxInstanceSize *string `field:"optional" json:"computeMaxInstanceSize" yaml:"computeMaxInstanceSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#compute_min_instance_size AdvancedCluster#compute_min_instance_size}.
	ComputeMinInstanceSize *string `field:"optional" json:"computeMinInstanceSize" yaml:"computeMinInstanceSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#compute_scale_down_enabled AdvancedCluster#compute_scale_down_enabled}.
	ComputeScaleDownEnabled interface{} `field:"optional" json:"computeScaleDownEnabled" yaml:"computeScaleDownEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#disk_gb_enabled AdvancedCluster#disk_gb_enabled}.
	DiskGbEnabled interface{} `field:"optional" json:"diskGbEnabled" yaml:"diskGbEnabled"`
}

