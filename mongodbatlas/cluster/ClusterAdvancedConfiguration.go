package cluster


type ClusterAdvancedConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#default_read_concern Cluster#default_read_concern}.
	DefaultReadConcern *string `field:"optional" json:"defaultReadConcern" yaml:"defaultReadConcern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#default_write_concern Cluster#default_write_concern}.
	DefaultWriteConcern *string `field:"optional" json:"defaultWriteConcern" yaml:"defaultWriteConcern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#fail_index_key_too_long Cluster#fail_index_key_too_long}.
	FailIndexKeyTooLong interface{} `field:"optional" json:"failIndexKeyTooLong" yaml:"failIndexKeyTooLong"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#javascript_enabled Cluster#javascript_enabled}.
	JavascriptEnabled interface{} `field:"optional" json:"javascriptEnabled" yaml:"javascriptEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#minimum_enabled_tls_protocol Cluster#minimum_enabled_tls_protocol}.
	MinimumEnabledTlsProtocol *string `field:"optional" json:"minimumEnabledTlsProtocol" yaml:"minimumEnabledTlsProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#no_table_scan Cluster#no_table_scan}.
	NoTableScan interface{} `field:"optional" json:"noTableScan" yaml:"noTableScan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#oplog_min_retention_hours Cluster#oplog_min_retention_hours}.
	OplogMinRetentionHours *float64 `field:"optional" json:"oplogMinRetentionHours" yaml:"oplogMinRetentionHours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#oplog_size_mb Cluster#oplog_size_mb}.
	OplogSizeMb *float64 `field:"optional" json:"oplogSizeMb" yaml:"oplogSizeMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#sample_refresh_interval_bi_connector Cluster#sample_refresh_interval_bi_connector}.
	SampleRefreshIntervalBiConnector *float64 `field:"optional" json:"sampleRefreshIntervalBiConnector" yaml:"sampleRefreshIntervalBiConnector"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cluster#sample_size_bi_connector Cluster#sample_size_bi_connector}.
	SampleSizeBiConnector *float64 `field:"optional" json:"sampleSizeBiConnector" yaml:"sampleSizeBiConnector"`
}

