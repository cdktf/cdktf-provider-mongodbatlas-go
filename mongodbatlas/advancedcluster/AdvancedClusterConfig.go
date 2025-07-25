// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdvancedClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#cluster_type AdvancedCluster#cluster_type}.
	ClusterType *string `field:"required" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#name AdvancedCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#project_id AdvancedCluster#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// replication_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#replication_specs AdvancedCluster#replication_specs}
	ReplicationSpecs interface{} `field:"required" json:"replicationSpecs" yaml:"replicationSpecs"`
	// Submit this field alongside your topology reconfiguration to request a new regional outage resistant topology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#accept_data_risks_and_force_replica_set_reconfig AdvancedCluster#accept_data_risks_and_force_replica_set_reconfig}
	AcceptDataRisksAndForceReplicaSetReconfig *string `field:"optional" json:"acceptDataRisksAndForceReplicaSetReconfig" yaml:"acceptDataRisksAndForceReplicaSetReconfig"`
	// advanced_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#advanced_configuration AdvancedCluster#advanced_configuration}
	AdvancedConfiguration *AdvancedClusterAdvancedConfiguration `field:"optional" json:"advancedConfiguration" yaml:"advancedConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#backup_enabled AdvancedCluster#backup_enabled}.
	BackupEnabled interface{} `field:"optional" json:"backupEnabled" yaml:"backupEnabled"`
	// bi_connector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#bi_connector_config AdvancedCluster#bi_connector_config}
	BiConnectorConfig *AdvancedClusterBiConnectorConfig `field:"optional" json:"biConnectorConfig" yaml:"biConnectorConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#config_server_management_mode AdvancedCluster#config_server_management_mode}.
	ConfigServerManagementMode *string `field:"optional" json:"configServerManagementMode" yaml:"configServerManagementMode"`
	// Flag that indicates whether to delete the cluster if the cluster creation times out. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#delete_on_create_timeout AdvancedCluster#delete_on_create_timeout}
	DeleteOnCreateTimeout interface{} `field:"optional" json:"deleteOnCreateTimeout" yaml:"deleteOnCreateTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#disk_size_gb AdvancedCluster#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#encryption_at_rest_provider AdvancedCluster#encryption_at_rest_provider}.
	EncryptionAtRestProvider *string `field:"optional" json:"encryptionAtRestProvider" yaml:"encryptionAtRestProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#global_cluster_self_managed_sharding AdvancedCluster#global_cluster_self_managed_sharding}.
	GlobalClusterSelfManagedSharding interface{} `field:"optional" json:"globalClusterSelfManagedSharding" yaml:"globalClusterSelfManagedSharding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#id AdvancedCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#labels AdvancedCluster#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#mongo_db_major_version AdvancedCluster#mongo_db_major_version}.
	MongoDbMajorVersion *string `field:"optional" json:"mongoDbMajorVersion" yaml:"mongoDbMajorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#paused AdvancedCluster#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// pinned_fcv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#pinned_fcv AdvancedCluster#pinned_fcv}
	PinnedFcv *AdvancedClusterPinnedFcv `field:"optional" json:"pinnedFcv" yaml:"pinnedFcv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#pit_enabled AdvancedCluster#pit_enabled}.
	PitEnabled interface{} `field:"optional" json:"pitEnabled" yaml:"pitEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#redact_client_log_data AdvancedCluster#redact_client_log_data}.
	RedactClientLogData interface{} `field:"optional" json:"redactClientLogData" yaml:"redactClientLogData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#replica_set_scaling_strategy AdvancedCluster#replica_set_scaling_strategy}.
	ReplicaSetScalingStrategy *string `field:"optional" json:"replicaSetScalingStrategy" yaml:"replicaSetScalingStrategy"`
	// Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#retain_backups_enabled AdvancedCluster#retain_backups_enabled}
	RetainBackupsEnabled interface{} `field:"optional" json:"retainBackupsEnabled" yaml:"retainBackupsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#root_cert_type AdvancedCluster#root_cert_type}.
	RootCertType *string `field:"optional" json:"rootCertType" yaml:"rootCertType"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#tags AdvancedCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#termination_protection_enabled AdvancedCluster#termination_protection_enabled}.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#timeouts AdvancedCluster#timeouts}
	Timeouts *AdvancedClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/advanced_cluster#version_release_system AdvancedCluster#version_release_system}.
	VersionReleaseSystem *string `field:"optional" json:"versionReleaseSystem" yaml:"versionReleaseSystem"`
}

