package globalclusterconfig


type GlobalClusterConfigManagedNamespaces struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#collection GlobalClusterConfig#collection}.
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#custom_shard_key GlobalClusterConfig#custom_shard_key}.
	CustomShardKey *string `field:"required" json:"customShardKey" yaml:"customShardKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#db GlobalClusterConfig#db}.
	Db *string `field:"required" json:"db" yaml:"db"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#is_custom_shard_key_hashed GlobalClusterConfig#is_custom_shard_key_hashed}.
	IsCustomShardKeyHashed interface{} `field:"optional" json:"isCustomShardKeyHashed" yaml:"isCustomShardKeyHashed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/global_cluster_config#is_shard_key_unique GlobalClusterConfig#is_shard_key_unique}.
	IsShardKeyUnique interface{} `field:"optional" json:"isShardKeyUnique" yaml:"isShardKeyUnique"`
}

