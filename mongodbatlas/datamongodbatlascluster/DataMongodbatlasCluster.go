// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlascluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/datamongodbatlascluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/data-sources/cluster mongodbatlas_cluster}.
type DataMongodbatlasCluster interface {
	cdktf.TerraformDataSource
	AdvancedConfiguration() DataMongodbatlasClusterAdvancedConfigurationList
	AutoScalingComputeEnabled() cdktf.IResolvable
	AutoScalingComputeScaleDownEnabled() cdktf.IResolvable
	AutoScalingDiskGbEnabled() cdktf.IResolvable
	BackingProviderName() *string
	BackupEnabled() cdktf.IResolvable
	BiConnectorConfig() DataMongodbatlasClusterBiConnectorConfigList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterType() *string
	ConnectionStrings() DataMongodbatlasClusterConnectionStringsList
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
	EncryptionAtRestProvider() *string
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
	Labels() DataMongodbatlasClusterLabelsList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongoDbMajorVersion() *string
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
	Paused() cdktf.IResolvable
	PitEnabled() cdktf.IResolvable
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderAutoScalingComputeMaxInstanceSize() *string
	ProviderAutoScalingComputeMinInstanceSize() *string
	ProviderBackupEnabled() cdktf.IResolvable
	ProviderDiskIops() *float64
	ProviderDiskTypeName() *string
	ProviderEncryptEbsVolume() cdktf.IResolvable
	ProviderEncryptEbsVolumeFlag() cdktf.IResolvable
	ProviderInstanceSizeName() *string
	ProviderName() *string
	ProviderRegionName() *string
	ProviderVolumeType() *string
	// Experimental.
	RawOverrides() interface{}
	RedactClientLogData() cdktf.IResolvable
	ReplicationFactor() *float64
	ReplicationSpecs() DataMongodbatlasClusterReplicationSpecsList
	SnapshotBackupPolicy() DataMongodbatlasClusterSnapshotBackupPolicyList
	SrvAddress() *string
	StateName() *string
	Tags() DataMongodbatlasClusterTagsList
	TerminationProtectionEnabled() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersionReleaseSystem() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataMongodbatlasCluster
type jsiiProxy_DataMongodbatlasCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataMongodbatlasCluster) AdvancedConfiguration() DataMongodbatlasClusterAdvancedConfigurationList {
	var returns DataMongodbatlasClusterAdvancedConfigurationList
	_jsii_.Get(
		j,
		"advancedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) AutoScalingComputeEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingComputeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) AutoScalingComputeScaleDownEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingComputeScaleDownEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) AutoScalingDiskGbEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingDiskGbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) BackingProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backingProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) BackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"backupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) BiConnectorConfig() DataMongodbatlasClusterBiConnectorConfigList {
	var returns DataMongodbatlasClusterBiConnectorConfigList
	_jsii_.Get(
		j,
		"biConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ConnectionStrings() DataMongodbatlasClusterConnectionStringsList {
	var returns DataMongodbatlasClusterConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) EncryptionAtRestProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Labels() DataMongodbatlasClusterLabelsList {
	var returns DataMongodbatlasClusterLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) MongoDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) MongoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) MongoUriUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUriUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) MongoUriWithOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUriWithOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) NumShards() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numShards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Paused() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) PitEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderAutoScalingComputeMaxInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMaxInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderAutoScalingComputeMinInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMinInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderBackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"providerBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderDiskIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"providerDiskIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderDiskTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerDiskTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderEncryptEbsVolume() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"providerEncryptEbsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderEncryptEbsVolumeFlag() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"providerEncryptEbsVolumeFlag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderInstanceSizeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInstanceSizeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ProviderVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) RedactClientLogData() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"redactClientLogData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) ReplicationSpecs() DataMongodbatlasClusterReplicationSpecsList {
	var returns DataMongodbatlasClusterReplicationSpecsList
	_jsii_.Get(
		j,
		"replicationSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) SnapshotBackupPolicy() DataMongodbatlasClusterSnapshotBackupPolicyList {
	var returns DataMongodbatlasClusterSnapshotBackupPolicyList
	_jsii_.Get(
		j,
		"snapshotBackupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) SrvAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srvAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) Tags() DataMongodbatlasClusterTagsList {
	var returns DataMongodbatlasClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) TerminationProtectionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCluster) VersionReleaseSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystem",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/data-sources/cluster mongodbatlas_cluster} Data Source.
func NewDataMongodbatlasCluster(scope constructs.Construct, id *string, config *DataMongodbatlasClusterConfig) DataMongodbatlasCluster {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasCluster{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/data-sources/cluster mongodbatlas_cluster} Data Source.
func NewDataMongodbatlasCluster_Override(d DataMongodbatlasCluster, scope constructs.Construct, id *string, config *DataMongodbatlasClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataMongodbatlasCluster resource upon running "cdktf plan <stack-name>".
func DataMongodbatlasCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataMongodbatlasCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
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
func DataMongodbatlasCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataMongodbatlasCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCluster.DataMongodbatlasCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataMongodbatlasCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

