package advancedcluster


type AdvancedClusterReplicationSpecs struct {
	// region_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#region_configs AdvancedCluster#region_configs}
	RegionConfigs interface{} `field:"required" json:"regionConfigs" yaml:"regionConfigs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#num_shards AdvancedCluster#num_shards}.
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/advanced_cluster#zone_name AdvancedCluster#zone_name}.
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}

