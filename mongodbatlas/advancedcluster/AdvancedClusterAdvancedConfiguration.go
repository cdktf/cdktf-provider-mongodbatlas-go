// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterAdvancedConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#change_stream_options_pre_and_post_images_expire_after_seconds AdvancedCluster#change_stream_options_pre_and_post_images_expire_after_seconds}.
	ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds *float64 `field:"optional" json:"changeStreamOptionsPreAndPostImagesExpireAfterSeconds" yaml:"changeStreamOptionsPreAndPostImagesExpireAfterSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#default_read_concern AdvancedCluster#default_read_concern}.
	DefaultReadConcern *string `field:"optional" json:"defaultReadConcern" yaml:"defaultReadConcern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#default_write_concern AdvancedCluster#default_write_concern}.
	DefaultWriteConcern *string `field:"optional" json:"defaultWriteConcern" yaml:"defaultWriteConcern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#fail_index_key_too_long AdvancedCluster#fail_index_key_too_long}.
	FailIndexKeyTooLong interface{} `field:"optional" json:"failIndexKeyTooLong" yaml:"failIndexKeyTooLong"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#javascript_enabled AdvancedCluster#javascript_enabled}.
	JavascriptEnabled interface{} `field:"optional" json:"javascriptEnabled" yaml:"javascriptEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#minimum_enabled_tls_protocol AdvancedCluster#minimum_enabled_tls_protocol}.
	MinimumEnabledTlsProtocol *string `field:"optional" json:"minimumEnabledTlsProtocol" yaml:"minimumEnabledTlsProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#no_table_scan AdvancedCluster#no_table_scan}.
	NoTableScan interface{} `field:"optional" json:"noTableScan" yaml:"noTableScan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#oplog_min_retention_hours AdvancedCluster#oplog_min_retention_hours}.
	OplogMinRetentionHours *float64 `field:"optional" json:"oplogMinRetentionHours" yaml:"oplogMinRetentionHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#oplog_size_mb AdvancedCluster#oplog_size_mb}.
	OplogSizeMb *float64 `field:"optional" json:"oplogSizeMb" yaml:"oplogSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#sample_refresh_interval_bi_connector AdvancedCluster#sample_refresh_interval_bi_connector}.
	SampleRefreshIntervalBiConnector *float64 `field:"optional" json:"sampleRefreshIntervalBiConnector" yaml:"sampleRefreshIntervalBiConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#sample_size_bi_connector AdvancedCluster#sample_size_bi_connector}.
	SampleSizeBiConnector *float64 `field:"optional" json:"sampleSizeBiConnector" yaml:"sampleSizeBiConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.3/docs/resources/advanced_cluster#transaction_lifetime_limit_seconds AdvancedCluster#transaction_lifetime_limit_seconds}.
	TransactionLifetimeLimitSeconds *float64 `field:"optional" json:"transactionLifetimeLimitSeconds" yaml:"transactionLifetimeLimitSeconds"`
}

