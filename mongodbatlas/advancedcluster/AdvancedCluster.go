// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v5/advancedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/advanced_cluster mongodbatlas_advanced_cluster}.
type AdvancedCluster interface {
	cdktf.TerraformResource
	AdvancedConfiguration() AdvancedClusterAdvancedConfigurationOutputReference
	AdvancedConfigurationInput() *AdvancedClusterAdvancedConfiguration
	BackupEnabled() interface{}
	SetBackupEnabled(val interface{})
	BackupEnabledInput() interface{}
	BiConnector() AdvancedClusterBiConnectorOutputReference
	BiConnectorConfig() AdvancedClusterBiConnectorConfigOutputReference
	BiConnectorConfigInput() *AdvancedClusterBiConnectorConfig
	BiConnectorInput() *AdvancedClusterBiConnector
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionStrings() AdvancedClusterConnectionStringsList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateDate() *string
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
	Labels() AdvancedClusterLabelsList
	LabelsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongoDbMajorVersion() *string
	SetMongoDbMajorVersion(val *string)
	MongoDbMajorVersionInput() *string
	MongoDbVersion() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReplicationSpecs() AdvancedClusterReplicationSpecsList
	ReplicationSpecsInput() interface{}
	RetainBackupsEnabled() interface{}
	SetRetainBackupsEnabled(val interface{})
	RetainBackupsEnabledInput() interface{}
	RootCertType() *string
	SetRootCertType(val *string)
	RootCertTypeInput() *string
	StateName() *string
	TerminationProtectionEnabled() interface{}
	SetTerminationProtectionEnabled(val interface{})
	TerminationProtectionEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AdvancedClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionReleaseSystem() *string
	SetVersionReleaseSystem(val *string)
	VersionReleaseSystemInput() *string
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
	PutAdvancedConfiguration(value *AdvancedClusterAdvancedConfiguration)
	PutBiConnector(value *AdvancedClusterBiConnector)
	PutBiConnectorConfig(value *AdvancedClusterBiConnectorConfig)
	PutLabels(value interface{})
	PutReplicationSpecs(value interface{})
	PutTimeouts(value *AdvancedClusterTimeouts)
	ResetAdvancedConfiguration()
	ResetBackupEnabled()
	ResetBiConnector()
	ResetBiConnectorConfig()
	ResetDiskSizeGb()
	ResetEncryptionAtRestProvider()
	ResetId()
	ResetLabels()
	ResetMongoDbMajorVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPaused()
	ResetPitEnabled()
	ResetRetainBackupsEnabled()
	ResetRootCertType()
	ResetTerminationProtectionEnabled()
	ResetTimeouts()
	ResetVersionReleaseSystem()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AdvancedCluster
type jsiiProxy_AdvancedCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AdvancedCluster) AdvancedConfiguration() AdvancedClusterAdvancedConfigurationOutputReference {
	var returns AdvancedClusterAdvancedConfigurationOutputReference
	_jsii_.Get(
		j,
		"advancedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) AdvancedConfigurationInput() *AdvancedClusterAdvancedConfiguration {
	var returns *AdvancedClusterAdvancedConfiguration
	_jsii_.Get(
		j,
		"advancedConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) BackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) BackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) BiConnector() AdvancedClusterBiConnectorOutputReference {
	var returns AdvancedClusterBiConnectorOutputReference
	_jsii_.Get(
		j,
		"biConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) BiConnectorConfig() AdvancedClusterBiConnectorConfigOutputReference {
	var returns AdvancedClusterBiConnectorConfigOutputReference
	_jsii_.Get(
		j,
		"biConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) BiConnectorConfigInput() *AdvancedClusterBiConnectorConfig {
	var returns *AdvancedClusterBiConnectorConfig
	_jsii_.Get(
		j,
		"biConnectorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) BiConnectorInput() *AdvancedClusterBiConnector {
	var returns *AdvancedClusterBiConnector
	_jsii_.Get(
		j,
		"biConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ConnectionStrings() AdvancedClusterConnectionStringsList {
	var returns AdvancedClusterConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) EncryptionAtRestProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) EncryptionAtRestProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Labels() AdvancedClusterLabelsList {
	var returns AdvancedClusterLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) MongoDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) MongoDbMajorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Paused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) PausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) PitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) PitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ReplicationSpecs() AdvancedClusterReplicationSpecsList {
	var returns AdvancedClusterReplicationSpecsList
	_jsii_.Get(
		j,
		"replicationSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) ReplicationSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) RetainBackupsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainBackupsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) RetainBackupsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainBackupsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) RootCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) RootCertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) TerminationProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) TerminationProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) Timeouts() AdvancedClusterTimeoutsOutputReference {
	var returns AdvancedClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) VersionReleaseSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedCluster) VersionReleaseSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystemInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/advanced_cluster mongodbatlas_advanced_cluster} Resource.
func NewAdvancedCluster(scope constructs.Construct, id *string, config *AdvancedClusterConfig) AdvancedCluster {
	_init_.Initialize()

	if err := validateNewAdvancedClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdvancedCluster{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/resources/advanced_cluster mongodbatlas_advanced_cluster} Resource.
func NewAdvancedCluster_Override(a AdvancedCluster, scope constructs.Construct, id *string, config *AdvancedClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetBackupEnabled(val interface{}) {
	if err := j.validateSetBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetEncryptionAtRestProvider(val *string) {
	if err := j.validateSetEncryptionAtRestProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtRestProvider",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetMongoDbMajorVersion(val *string) {
	if err := j.validateSetMongoDbMajorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongoDbMajorVersion",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetPaused(val interface{}) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetPitEnabled(val interface{}) {
	if err := j.validateSetPitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetRetainBackupsEnabled(val interface{}) {
	if err := j.validateSetRetainBackupsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retainBackupsEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetRootCertType(val *string) {
	if err := j.validateSetRootCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootCertType",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetTerminationProtectionEnabled(val interface{}) {
	if err := j.validateSetTerminationProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedCluster)SetVersionReleaseSystem(val *string) {
	if err := j.validateSetVersionReleaseSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionReleaseSystem",
		val,
	)
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
func AdvancedCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdvancedCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdvancedCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdvancedCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdvancedCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdvancedCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AdvancedCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdvancedCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AdvancedCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AdvancedCluster) PutAdvancedConfiguration(value *AdvancedClusterAdvancedConfiguration) {
	if err := a.validatePutAdvancedConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedCluster) PutBiConnector(value *AdvancedClusterBiConnector) {
	if err := a.validatePutBiConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBiConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedCluster) PutBiConnectorConfig(value *AdvancedClusterBiConnectorConfig) {
	if err := a.validatePutBiConnectorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBiConnectorConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedCluster) PutLabels(value interface{}) {
	if err := a.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabels",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedCluster) PutReplicationSpecs(value interface{}) {
	if err := a.validatePutReplicationSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplicationSpecs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedCluster) PutTimeouts(value *AdvancedClusterTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetAdvancedConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetBackupEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetBiConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetBiConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetBiConnectorConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetBiConnectorConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		a,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetEncryptionAtRestProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionAtRestProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetMongoDbMajorVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMongoDbMajorVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetPaused() {
	_jsii_.InvokeVoid(
		a,
		"resetPaused",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetPitEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPitEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetRetainBackupsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRetainBackupsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetRootCertType() {
	_jsii_.InvokeVoid(
		a,
		"resetRootCertType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetTerminationProtectionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationProtectionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) ResetVersionReleaseSystem() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionReleaseSystem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

