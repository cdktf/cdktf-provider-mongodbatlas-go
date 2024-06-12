// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.0/docs/resources/cluster mongodbatlas_cluster}.
type Cluster interface {
	cdktf.TerraformResource
	AcceptDataRisksAndForceReplicaSetReconfig() *string
	SetAcceptDataRisksAndForceReplicaSetReconfig(val *string)
	AcceptDataRisksAndForceReplicaSetReconfigInput() *string
	AdvancedConfiguration() ClusterAdvancedConfigurationOutputReference
	AdvancedConfigurationInput() *ClusterAdvancedConfiguration
	AutoScalingComputeEnabled() interface{}
	SetAutoScalingComputeEnabled(val interface{})
	AutoScalingComputeEnabledInput() interface{}
	AutoScalingComputeScaleDownEnabled() interface{}
	SetAutoScalingComputeScaleDownEnabled(val interface{})
	AutoScalingComputeScaleDownEnabledInput() interface{}
	AutoScalingDiskGbEnabled() interface{}
	SetAutoScalingDiskGbEnabled(val interface{})
	AutoScalingDiskGbEnabledInput() interface{}
	BackingProviderName() *string
	SetBackingProviderName(val *string)
	BackingProviderNameInput() *string
	BackupEnabled() interface{}
	SetBackupEnabled(val interface{})
	BackupEnabledInput() interface{}
	BiConnectorConfig() ClusterBiConnectorConfigOutputReference
	BiConnectorConfigInput() *ClusterBiConnectorConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudBackup() interface{}
	SetCloudBackup(val interface{})
	CloudBackupInput() interface{}
	ClusterId() *string
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionStrings() ClusterConnectionStringsList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerId() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	EncryptionAtRestProvider() *string
	SetEncryptionAtRestProvider(val *string)
	EncryptionAtRestProviderInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() ClusterLabelsList
	LabelsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongoDbMajorVersion() *string
	SetMongoDbMajorVersion(val *string)
	MongoDbMajorVersionInput() *string
	MongoDbVersion() *string
	MongoUri() *string
	MongoUriUpdated() *string
	MongoUriWithOptions() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NumShards() *float64
	SetNumShards(val *float64)
	NumShardsInput() *float64
	Paused() interface{}
	SetPaused(val interface{})
	PausedInput() interface{}
	PitEnabled() interface{}
	SetPitEnabled(val interface{})
	PitEnabledInput() interface{}
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderAutoScalingComputeMaxInstanceSize() *string
	SetProviderAutoScalingComputeMaxInstanceSize(val *string)
	ProviderAutoScalingComputeMaxInstanceSizeInput() *string
	ProviderAutoScalingComputeMinInstanceSize() *string
	SetProviderAutoScalingComputeMinInstanceSize(val *string)
	ProviderAutoScalingComputeMinInstanceSizeInput() *string
	ProviderDiskIops() *float64
	SetProviderDiskIops(val *float64)
	ProviderDiskIopsInput() *float64
	ProviderDiskTypeName() *string
	SetProviderDiskTypeName(val *string)
	ProviderDiskTypeNameInput() *string
	ProviderEncryptEbsVolume() interface{}
	SetProviderEncryptEbsVolume(val interface{})
	ProviderEncryptEbsVolumeFlag() cdktf.IResolvable
	ProviderEncryptEbsVolumeInput() interface{}
	ProviderInstanceSizeName() *string
	SetProviderInstanceSizeName(val *string)
	ProviderInstanceSizeNameInput() *string
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	ProviderRegionName() *string
	SetProviderRegionName(val *string)
	ProviderRegionNameInput() *string
	ProviderVolumeType() *string
	SetProviderVolumeType(val *string)
	ProviderVolumeTypeInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReplicationFactor() *float64
	SetReplicationFactor(val *float64)
	ReplicationFactorInput() *float64
	ReplicationSpecs() ClusterReplicationSpecsList
	ReplicationSpecsInput() interface{}
	RetainBackupsEnabled() interface{}
	SetRetainBackupsEnabled(val interface{})
	RetainBackupsEnabledInput() interface{}
	SnapshotBackupPolicy() ClusterSnapshotBackupPolicyList
	SrvAddress() *string
	StateName() *string
	Tags() ClusterTagsList
	TagsInput() interface{}
	TerminationProtectionEnabled() interface{}
	SetTerminationProtectionEnabled(val interface{})
	TerminationProtectionEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionReleaseSystem() *string
	SetVersionReleaseSystem(val *string)
	VersionReleaseSystemInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdvancedConfiguration(value *ClusterAdvancedConfiguration)
	PutBiConnectorConfig(value *ClusterBiConnectorConfig)
	PutLabels(value interface{})
	PutReplicationSpecs(value interface{})
	PutTags(value interface{})
	PutTimeouts(value *ClusterTimeouts)
	ResetAcceptDataRisksAndForceReplicaSetReconfig()
	ResetAdvancedConfiguration()
	ResetAutoScalingComputeEnabled()
	ResetAutoScalingComputeScaleDownEnabled()
	ResetAutoScalingDiskGbEnabled()
	ResetBackingProviderName()
	ResetBackupEnabled()
	ResetBiConnectorConfig()
	ResetCloudBackup()
	ResetClusterType()
	ResetDiskSizeGb()
	ResetEncryptionAtRestProvider()
	ResetId()
	ResetLabels()
	ResetMongoDbMajorVersion()
	ResetNumShards()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPaused()
	ResetPitEnabled()
	ResetProviderAutoScalingComputeMaxInstanceSize()
	ResetProviderAutoScalingComputeMinInstanceSize()
	ResetProviderDiskIops()
	ResetProviderDiskTypeName()
	ResetProviderEncryptEbsVolume()
	ResetProviderRegionName()
	ResetProviderVolumeType()
	ResetReplicationFactor()
	ResetReplicationSpecs()
	ResetRetainBackupsEnabled()
	ResetTags()
	ResetTerminationProtectionEnabled()
	ResetTimeouts()
	ResetVersionReleaseSystem()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Cluster) AcceptDataRisksAndForceReplicaSetReconfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptDataRisksAndForceReplicaSetReconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AcceptDataRisksAndForceReplicaSetReconfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptDataRisksAndForceReplicaSetReconfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AdvancedConfiguration() ClusterAdvancedConfigurationOutputReference {
	var returns ClusterAdvancedConfigurationOutputReference
	_jsii_.Get(
		j,
		"advancedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AdvancedConfigurationInput() *ClusterAdvancedConfiguration {
	var returns *ClusterAdvancedConfiguration
	_jsii_.Get(
		j,
		"advancedConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoScalingComputeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingComputeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoScalingComputeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingComputeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoScalingComputeScaleDownEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingComputeScaleDownEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoScalingComputeScaleDownEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingComputeScaleDownEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoScalingDiskGbEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingDiskGbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) AutoScalingDiskGbEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingDiskGbEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BackingProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backingProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BackingProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backingProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BiConnectorConfig() ClusterBiConnectorConfigOutputReference {
	var returns ClusterBiConnectorConfigOutputReference
	_jsii_.Get(
		j,
		"biConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BiConnectorConfigInput() *ClusterBiConnectorConfig {
	var returns *ClusterBiConnectorConfig
	_jsii_.Get(
		j,
		"biConnectorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CloudBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CloudBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConnectionStrings() ClusterConnectionStringsList {
	var returns ClusterConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EncryptionAtRestProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) EncryptionAtRestProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Labels() ClusterLabelsList {
	var returns ClusterLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MongoDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MongoDbMajorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MongoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MongoUriUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUriUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) MongoUriWithOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUriWithOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NumShards() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numShards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) NumShardsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numShardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Paused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderAutoScalingComputeMaxInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMaxInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderAutoScalingComputeMaxInstanceSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMaxInstanceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderAutoScalingComputeMinInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMinInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderAutoScalingComputeMinInstanceSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMinInstanceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderDiskIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"providerDiskIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderDiskIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"providerDiskIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderDiskTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerDiskTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderDiskTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerDiskTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderEncryptEbsVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerEncryptEbsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderEncryptEbsVolumeFlag() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"providerEncryptEbsVolumeFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderEncryptEbsVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"providerEncryptEbsVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderInstanceSizeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInstanceSizeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderInstanceSizeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInstanceSizeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderRegionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerRegionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ProviderVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ReplicationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ReplicationSpecs() ClusterReplicationSpecsList {
	var returns ClusterReplicationSpecsList
	_jsii_.Get(
		j,
		"replicationSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ReplicationSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RetainBackupsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainBackupsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) RetainBackupsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainBackupsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SnapshotBackupPolicy() ClusterSnapshotBackupPolicyList {
	var returns ClusterSnapshotBackupPolicyList
	_jsii_.Get(
		j,
		"snapshotBackupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) SrvAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srvAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Tags() ClusterTagsList {
	var returns ClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerminationProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerminationProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Timeouts() ClusterTimeoutsOutputReference {
	var returns ClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VersionReleaseSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) VersionReleaseSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystemInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.0/docs/resources/cluster mongodbatlas_cluster} Resource.
func NewCluster(scope constructs.Construct, id *string, config *ClusterConfig) Cluster {
	_init_.Initialize()

	if err := validateNewClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.0/docs/resources/cluster mongodbatlas_cluster} Resource.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, config *ClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Cluster)SetAcceptDataRisksAndForceReplicaSetReconfig(val *string) {
	if err := j.validateSetAcceptDataRisksAndForceReplicaSetReconfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptDataRisksAndForceReplicaSetReconfig",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetAutoScalingComputeEnabled(val interface{}) {
	if err := j.validateSetAutoScalingComputeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingComputeEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetAutoScalingComputeScaleDownEnabled(val interface{}) {
	if err := j.validateSetAutoScalingComputeScaleDownEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingComputeScaleDownEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetAutoScalingDiskGbEnabled(val interface{}) {
	if err := j.validateSetAutoScalingDiskGbEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingDiskGbEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetBackingProviderName(val *string) {
	if err := j.validateSetBackingProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backingProviderName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetBackupEnabled(val interface{}) {
	if err := j.validateSetBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetCloudBackup(val interface{}) {
	if err := j.validateSetCloudBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudBackup",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetEncryptionAtRestProvider(val *string) {
	if err := j.validateSetEncryptionAtRestProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtRestProvider",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetMongoDbMajorVersion(val *string) {
	if err := j.validateSetMongoDbMajorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongoDbMajorVersion",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetNumShards(val *float64) {
	if err := j.validateSetNumShardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numShards",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetPaused(val interface{}) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetPitEnabled(val interface{}) {
	if err := j.validateSetPitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderAutoScalingComputeMaxInstanceSize(val *string) {
	if err := j.validateSetProviderAutoScalingComputeMaxInstanceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerAutoScalingComputeMaxInstanceSize",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderAutoScalingComputeMinInstanceSize(val *string) {
	if err := j.validateSetProviderAutoScalingComputeMinInstanceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerAutoScalingComputeMinInstanceSize",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderDiskIops(val *float64) {
	if err := j.validateSetProviderDiskIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerDiskIops",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderDiskTypeName(val *string) {
	if err := j.validateSetProviderDiskTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerDiskTypeName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderEncryptEbsVolume(val interface{}) {
	if err := j.validateSetProviderEncryptEbsVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerEncryptEbsVolume",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderInstanceSizeName(val *string) {
	if err := j.validateSetProviderInstanceSizeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerInstanceSizeName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderRegionName(val *string) {
	if err := j.validateSetProviderRegionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerRegionName",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProviderVolumeType(val *string) {
	if err := j.validateSetProviderVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerVolumeType",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetReplicationFactor(val *float64) {
	if err := j.validateSetReplicationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationFactor",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetRetainBackupsEnabled(val interface{}) {
	if err := j.validateSetRetainBackupsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainBackupsEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetTerminationProtectionEnabled(val interface{}) {
	if err := j.validateSetTerminationProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_Cluster)SetVersionReleaseSystem(val *string) {
	if err := j.validateSetVersionReleaseSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionReleaseSystem",
		val,
	)
}

// Generates CDKTF code for importing a Cluster resource upon running "cdktf plan <stack-name>".
func Cluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Cluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.cluster.Cluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Cluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Cluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Cluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Cluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Cluster) PutAdvancedConfiguration(value *ClusterAdvancedConfiguration) {
	if err := c.validatePutAdvancedConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdvancedConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutBiConnectorConfig(value *ClusterBiConnectorConfig) {
	if err := c.validatePutBiConnectorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBiConnectorConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutLabels(value interface{}) {
	if err := c.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLabels",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutReplicationSpecs(value interface{}) {
	if err := c.validatePutReplicationSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReplicationSpecs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutTags(value interface{}) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) PutTimeouts(value *ClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cluster) ResetAcceptDataRisksAndForceReplicaSetReconfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAcceptDataRisksAndForceReplicaSetReconfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAdvancedConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvancedConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAutoScalingComputeEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingComputeEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAutoScalingComputeScaleDownEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingComputeScaleDownEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetAutoScalingDiskGbEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScalingDiskGbEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetBackingProviderName() {
	_jsii_.InvokeVoid(
		c,
		"resetBackingProviderName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetBackupEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetBackupEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetBiConnectorConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBiConnectorConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetCloudBackup() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudBackup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetClusterType() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetEncryptionAtRestProvider() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionAtRestProvider",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetMongoDbMajorVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMongoDbMajorVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetNumShards() {
	_jsii_.InvokeVoid(
		c,
		"resetNumShards",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetPaused() {
	_jsii_.InvokeVoid(
		c,
		"resetPaused",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetPitEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPitEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderAutoScalingComputeMaxInstanceSize() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderAutoScalingComputeMaxInstanceSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderAutoScalingComputeMinInstanceSize() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderAutoScalingComputeMinInstanceSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderDiskIops() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderDiskIops",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderDiskTypeName() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderDiskTypeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderEncryptEbsVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderEncryptEbsVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderRegionName() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderRegionName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetProviderVolumeType() {
	_jsii_.InvokeVoid(
		c,
		"resetProviderVolumeType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetReplicationFactor() {
	_jsii_.InvokeVoid(
		c,
		"resetReplicationFactor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetReplicationSpecs() {
	_jsii_.InvokeVoid(
		c,
		"resetReplicationSpecs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetRetainBackupsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetRetainBackupsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTerminationProtectionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationProtectionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) ResetVersionReleaseSystem() {
	_jsii_.InvokeVoid(
		c,
		"resetVersionReleaseSystem",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

