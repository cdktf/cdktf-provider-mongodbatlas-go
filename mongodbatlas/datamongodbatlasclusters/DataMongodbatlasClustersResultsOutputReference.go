// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasclusters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/datamongodbatlasclusters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasClustersResultsOutputReference interface {
	cdktf.ComplexObject
	AdvancedConfiguration() DataMongodbatlasClustersResultsAdvancedConfigurationList
	AutoScalingComputeEnabled() cdktf.IResolvable
	AutoScalingComputeScaleDownEnabled() cdktf.IResolvable
	AutoScalingDiskGbEnabled() cdktf.IResolvable
	BackingProviderName() *string
	BackupEnabled() cdktf.IResolvable
	BiConnectorConfig() DataMongodbatlasClustersResultsBiConnectorConfigList
	ClusterType() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectionStrings() DataMongodbatlasClustersResultsConnectionStringsList
	ContainerId() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSizeGb() *float64
	EncryptionAtRestProvider() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataMongodbatlasClustersResults
	SetInternalValue(val *DataMongodbatlasClustersResults)
	Labels() DataMongodbatlasClustersResultsLabelsList
	MongoDbMajorVersion() *string
	MongoDbVersion() *string
	MongoUri() *string
	MongoUriUpdated() *string
	MongoUriWithOptions() *string
	Name() *string
	NumShards() *float64
	Paused() cdktf.IResolvable
	PinnedFcv() DataMongodbatlasClustersResultsPinnedFcvList
	PitEnabled() cdktf.IResolvable
	ProviderAutoScalingComputeMaxInstanceSize() *string
	ProviderAutoScalingComputeMinInstanceSize() *string
	ProviderBackupEnabled() cdktf.IResolvable
	ProviderDiskIops() *float64
	ProviderDiskTypeName() *string
	ProviderEncryptEbsVolume() cdktf.IResolvable
	ProviderInstanceSizeName() *string
	ProviderName() *string
	ProviderRegionName() *string
	ProviderVolumeType() *string
	RedactClientLogData() cdktf.IResolvable
	ReplicationFactor() *float64
	ReplicationSpecs() DataMongodbatlasClustersResultsReplicationSpecsList
	SnapshotBackupPolicy() DataMongodbatlasClustersResultsSnapshotBackupPolicyList
	SrvAddress() *string
	StateName() *string
	Tags() DataMongodbatlasClustersResultsTagsList
	TerminationProtectionEnabled() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersionReleaseSystem() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataMongodbatlasClustersResultsOutputReference
type jsiiProxy_DataMongodbatlasClustersResultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) AdvancedConfiguration() DataMongodbatlasClustersResultsAdvancedConfigurationList {
	var returns DataMongodbatlasClustersResultsAdvancedConfigurationList
	_jsii_.Get(
		j,
		"advancedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) AutoScalingComputeEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingComputeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) AutoScalingComputeScaleDownEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingComputeScaleDownEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) AutoScalingDiskGbEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoScalingDiskGbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) BackingProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backingProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) BackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"backupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) BiConnectorConfig() DataMongodbatlasClustersResultsBiConnectorConfigList {
	var returns DataMongodbatlasClustersResultsBiConnectorConfigList
	_jsii_.Get(
		j,
		"biConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ConnectionStrings() DataMongodbatlasClustersResultsConnectionStringsList {
	var returns DataMongodbatlasClustersResultsConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) EncryptionAtRestProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) InternalValue() *DataMongodbatlasClustersResults {
	var returns *DataMongodbatlasClustersResults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) Labels() DataMongodbatlasClustersResultsLabelsList {
	var returns DataMongodbatlasClustersResultsLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) MongoDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) MongoUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) MongoUriUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUriUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) MongoUriWithOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoUriWithOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) NumShards() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numShards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) Paused() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) PinnedFcv() DataMongodbatlasClustersResultsPinnedFcvList {
	var returns DataMongodbatlasClustersResultsPinnedFcvList
	_jsii_.Get(
		j,
		"pinnedFcv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) PitEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderAutoScalingComputeMaxInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMaxInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderAutoScalingComputeMinInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerAutoScalingComputeMinInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderBackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"providerBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderDiskIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"providerDiskIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderDiskTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerDiskTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderEncryptEbsVolume() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"providerEncryptEbsVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderInstanceSizeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInstanceSizeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ProviderVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) RedactClientLogData() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"redactClientLogData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ReplicationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ReplicationSpecs() DataMongodbatlasClustersResultsReplicationSpecsList {
	var returns DataMongodbatlasClustersResultsReplicationSpecsList
	_jsii_.Get(
		j,
		"replicationSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) SnapshotBackupPolicy() DataMongodbatlasClustersResultsSnapshotBackupPolicyList {
	var returns DataMongodbatlasClustersResultsSnapshotBackupPolicyList
	_jsii_.Get(
		j,
		"snapshotBackupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) SrvAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srvAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) Tags() DataMongodbatlasClustersResultsTagsList {
	var returns DataMongodbatlasClustersResultsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) TerminationProtectionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) VersionReleaseSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystem",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasClustersResultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataMongodbatlasClustersResultsOutputReference {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasClustersResultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasClustersResultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasClustersResultsOutputReference_Override(d DataMongodbatlasClustersResultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference)SetInternalValue(val *DataMongodbatlasClustersResults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

