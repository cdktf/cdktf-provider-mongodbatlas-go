package cluster


type ClusterReplicationSpecsRegionsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#region_name Cluster#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#analytics_nodes Cluster#analytics_nodes}.
	AnalyticsNodes *float64 `field:"optional" json:"analyticsNodes" yaml:"analyticsNodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#electable_nodes Cluster#electable_nodes}.
	ElectableNodes *float64 `field:"optional" json:"electableNodes" yaml:"electableNodes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#priority Cluster#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#read_only_nodes Cluster#read_only_nodes}.
	ReadOnlyNodes *float64 `field:"optional" json:"readOnlyNodes" yaml:"readOnlyNodes"`
}

