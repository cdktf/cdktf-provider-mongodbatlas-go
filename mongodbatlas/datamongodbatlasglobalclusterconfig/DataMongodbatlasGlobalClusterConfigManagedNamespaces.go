package datamongodbatlasglobalclusterconfig


type DataMongodbatlasGlobalClusterConfigManagedNamespaces struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/global_cluster_config#collection DataMongodbatlasGlobalClusterConfig#collection}.
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/global_cluster_config#custom_shard_key DataMongodbatlasGlobalClusterConfig#custom_shard_key}.
	CustomShardKey *string `field:"required" json:"customShardKey" yaml:"customShardKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/global_cluster_config#db DataMongodbatlasGlobalClusterConfig#db}.
	Db *string `field:"required" json:"db" yaml:"db"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/global_cluster_config#is_custom_shard_key_hashed DataMongodbatlasGlobalClusterConfig#is_custom_shard_key_hashed}.
	IsCustomShardKeyHashed interface{} `field:"optional" json:"isCustomShardKeyHashed" yaml:"isCustomShardKeyHashed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/global_cluster_config#is_shard_key_unique DataMongodbatlasGlobalClusterConfig#is_shard_key_unique}.
	IsShardKeyUnique interface{} `field:"optional" json:"isShardKeyUnique" yaml:"isShardKeyUnique"`
}

