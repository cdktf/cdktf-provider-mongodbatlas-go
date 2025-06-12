// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterAdvancedConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#change_stream_options_pre_and_post_images_expire_after_seconds Cluster#change_stream_options_pre_and_post_images_expire_after_seconds}.
	ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds *float64 `field:"optional" json:"changeStreamOptionsPreAndPostImagesExpireAfterSeconds" yaml:"changeStreamOptionsPreAndPostImagesExpireAfterSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#custom_openssl_cipher_config_tls12 Cluster#custom_openssl_cipher_config_tls12}.
	CustomOpensslCipherConfigTls12 *[]*string `field:"optional" json:"customOpensslCipherConfigTls12" yaml:"customOpensslCipherConfigTls12"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#default_max_time_ms Cluster#default_max_time_ms}.
	DefaultMaxTimeMs *float64 `field:"optional" json:"defaultMaxTimeMs" yaml:"defaultMaxTimeMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#default_read_concern Cluster#default_read_concern}.
	DefaultReadConcern *string `field:"optional" json:"defaultReadConcern" yaml:"defaultReadConcern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#default_write_concern Cluster#default_write_concern}.
	DefaultWriteConcern *string `field:"optional" json:"defaultWriteConcern" yaml:"defaultWriteConcern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#fail_index_key_too_long Cluster#fail_index_key_too_long}.
	FailIndexKeyTooLong interface{} `field:"optional" json:"failIndexKeyTooLong" yaml:"failIndexKeyTooLong"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#javascript_enabled Cluster#javascript_enabled}.
	JavascriptEnabled interface{} `field:"optional" json:"javascriptEnabled" yaml:"javascriptEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#minimum_enabled_tls_protocol Cluster#minimum_enabled_tls_protocol}.
	MinimumEnabledTlsProtocol *string `field:"optional" json:"minimumEnabledTlsProtocol" yaml:"minimumEnabledTlsProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#no_table_scan Cluster#no_table_scan}.
	NoTableScan interface{} `field:"optional" json:"noTableScan" yaml:"noTableScan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#oplog_min_retention_hours Cluster#oplog_min_retention_hours}.
	OplogMinRetentionHours *float64 `field:"optional" json:"oplogMinRetentionHours" yaml:"oplogMinRetentionHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#oplog_size_mb Cluster#oplog_size_mb}.
	OplogSizeMb *float64 `field:"optional" json:"oplogSizeMb" yaml:"oplogSizeMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#sample_refresh_interval_bi_connector Cluster#sample_refresh_interval_bi_connector}.
	SampleRefreshIntervalBiConnector *float64 `field:"optional" json:"sampleRefreshIntervalBiConnector" yaml:"sampleRefreshIntervalBiConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#sample_size_bi_connector Cluster#sample_size_bi_connector}.
	SampleSizeBiConnector *float64 `field:"optional" json:"sampleSizeBiConnector" yaml:"sampleSizeBiConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#tls_cipher_config_mode Cluster#tls_cipher_config_mode}.
	TlsCipherConfigMode *string `field:"optional" json:"tlsCipherConfigMode" yaml:"tlsCipherConfigMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs/resources/cluster#transaction_lifetime_limit_seconds Cluster#transaction_lifetime_limit_seconds}.
	TransactionLifetimeLimitSeconds *float64 `field:"optional" json:"transactionLifetimeLimitSeconds" yaml:"transactionLifetimeLimitSeconds"`
}

