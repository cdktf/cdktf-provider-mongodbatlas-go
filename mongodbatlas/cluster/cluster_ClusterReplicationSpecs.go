package cluster


type ClusterReplicationSpecs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#num_shards Cluster#num_shards}.
	NumShards *float64 `field:"required" json:"numShards" yaml:"numShards"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#id Cluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// regions_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#regions_config Cluster#regions_config}
	RegionsConfig interface{} `field:"optional" json:"regionsConfig" yaml:"regionsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#zone_name Cluster#zone_name}.
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}

