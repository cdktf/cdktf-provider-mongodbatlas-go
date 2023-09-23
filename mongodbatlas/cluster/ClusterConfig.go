// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#name Cluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#project_id Cluster#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_instance_size_name Cluster#provider_instance_size_name}.
	ProviderInstanceSizeName *string `field:"required" json:"providerInstanceSizeName" yaml:"providerInstanceSizeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_name Cluster#provider_name}.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// advanced_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#advanced_configuration Cluster#advanced_configuration}
	AdvancedConfiguration *ClusterAdvancedConfiguration `field:"optional" json:"advancedConfiguration" yaml:"advancedConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#auto_scaling_compute_enabled Cluster#auto_scaling_compute_enabled}.
	AutoScalingComputeEnabled interface{} `field:"optional" json:"autoScalingComputeEnabled" yaml:"autoScalingComputeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#auto_scaling_compute_scale_down_enabled Cluster#auto_scaling_compute_scale_down_enabled}.
	AutoScalingComputeScaleDownEnabled interface{} `field:"optional" json:"autoScalingComputeScaleDownEnabled" yaml:"autoScalingComputeScaleDownEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#auto_scaling_disk_gb_enabled Cluster#auto_scaling_disk_gb_enabled}.
	AutoScalingDiskGbEnabled interface{} `field:"optional" json:"autoScalingDiskGbEnabled" yaml:"autoScalingDiskGbEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#backing_provider_name Cluster#backing_provider_name}.
	BackingProviderName *string `field:"optional" json:"backingProviderName" yaml:"backingProviderName"`
	// Clusters running MongoDB FCV 4.2 or later and any new Atlas clusters of any type do not support this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#backup_enabled Cluster#backup_enabled}
	BackupEnabled interface{} `field:"optional" json:"backupEnabled" yaml:"backupEnabled"`
	// bi_connector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#bi_connector_config Cluster#bi_connector_config}
	BiConnectorConfig *ClusterBiConnectorConfig `field:"optional" json:"biConnectorConfig" yaml:"biConnectorConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#cloud_backup Cluster#cloud_backup}.
	CloudBackup interface{} `field:"optional" json:"cloudBackup" yaml:"cloudBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#cluster_type Cluster#cluster_type}.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#disk_size_gb Cluster#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#encryption_at_rest_provider Cluster#encryption_at_rest_provider}.
	EncryptionAtRestProvider *string `field:"optional" json:"encryptionAtRestProvider" yaml:"encryptionAtRestProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#id Cluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#labels Cluster#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#mongo_db_major_version Cluster#mongo_db_major_version}.
	MongoDbMajorVersion *string `field:"optional" json:"mongoDbMajorVersion" yaml:"mongoDbMajorVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#num_shards Cluster#num_shards}.
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#paused Cluster#paused}.
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#pit_enabled Cluster#pit_enabled}.
	PitEnabled interface{} `field:"optional" json:"pitEnabled" yaml:"pitEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_auto_scaling_compute_max_instance_size Cluster#provider_auto_scaling_compute_max_instance_size}.
	ProviderAutoScalingComputeMaxInstanceSize *string `field:"optional" json:"providerAutoScalingComputeMaxInstanceSize" yaml:"providerAutoScalingComputeMaxInstanceSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_auto_scaling_compute_min_instance_size Cluster#provider_auto_scaling_compute_min_instance_size}.
	ProviderAutoScalingComputeMinInstanceSize *string `field:"optional" json:"providerAutoScalingComputeMinInstanceSize" yaml:"providerAutoScalingComputeMinInstanceSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_disk_iops Cluster#provider_disk_iops}.
	ProviderDiskIops *float64 `field:"optional" json:"providerDiskIops" yaml:"providerDiskIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_disk_type_name Cluster#provider_disk_type_name}.
	ProviderDiskTypeName *string `field:"optional" json:"providerDiskTypeName" yaml:"providerDiskTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_encrypt_ebs_volume Cluster#provider_encrypt_ebs_volume}.
	ProviderEncryptEbsVolume interface{} `field:"optional" json:"providerEncryptEbsVolume" yaml:"providerEncryptEbsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_region_name Cluster#provider_region_name}.
	ProviderRegionName *string `field:"optional" json:"providerRegionName" yaml:"providerRegionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#provider_volume_type Cluster#provider_volume_type}.
	ProviderVolumeType *string `field:"optional" json:"providerVolumeType" yaml:"providerVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#replication_factor Cluster#replication_factor}.
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// replication_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#replication_specs Cluster#replication_specs}
	ReplicationSpecs interface{} `field:"optional" json:"replicationSpecs" yaml:"replicationSpecs"`
	// Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#retain_backups_enabled Cluster#retain_backups_enabled}
	RetainBackupsEnabled interface{} `field:"optional" json:"retainBackupsEnabled" yaml:"retainBackupsEnabled"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#tags Cluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#termination_protection_enabled Cluster#termination_protection_enabled}.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#timeouts Cluster#timeouts}
	Timeouts *ClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.12.1/docs/resources/cluster#version_release_system Cluster#version_release_system}.
	VersionReleaseSystem *string `field:"optional" json:"versionReleaseSystem" yaml:"versionReleaseSystem"`
}

