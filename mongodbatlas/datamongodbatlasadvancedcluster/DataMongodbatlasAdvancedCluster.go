package datamongodbatlasadvancedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/datamongodbatlasadvancedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/data-sources/advanced_cluster mongodbatlas_advanced_cluster}.
type DataMongodbatlasAdvancedCluster interface {
	cdktf.TerraformDataSource
	AdvancedConfiguration() DataMongodbatlasAdvancedClusterAdvancedConfigurationList
	BackupEnabled() cdktf.IResolvable
	BiConnectorConfig() DataMongodbatlasAdvancedClusterBiConnectorConfigList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterType() *string
	ConnectionStrings() DataMongodbatlasAdvancedClusterConnectionStringsList
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
	Labels() DataMongodbatlasAdvancedClusterLabelsList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongoDbMajorVersion() *string
	MongoDbVersion() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Paused() cdktf.IResolvable
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
	RawOverrides() interface{}
	ReplicationSpecs() DataMongodbatlasAdvancedClusterReplicationSpecsList
	RootCertType() *string
	StateName() *string
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
	ResetPitEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataMongodbatlasAdvancedCluster
type jsiiProxy_DataMongodbatlasAdvancedCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) AdvancedConfiguration() DataMongodbatlasAdvancedClusterAdvancedConfigurationList {
	var returns DataMongodbatlasAdvancedClusterAdvancedConfigurationList
	_jsii_.Get(
		j,
		"advancedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) BackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"backupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) BiConnectorConfig() DataMongodbatlasAdvancedClusterBiConnectorConfigList {
	var returns DataMongodbatlasAdvancedClusterBiConnectorConfigList
	_jsii_.Get(
		j,
		"biConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ConnectionStrings() DataMongodbatlasAdvancedClusterConnectionStringsList {
	var returns DataMongodbatlasAdvancedClusterConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) EncryptionAtRestProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Labels() DataMongodbatlasAdvancedClusterLabelsList {
	var returns DataMongodbatlasAdvancedClusterLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) MongoDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Paused() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) PitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) PitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) ReplicationSpecs() DataMongodbatlasAdvancedClusterReplicationSpecsList {
	var returns DataMongodbatlasAdvancedClusterReplicationSpecsList
	_jsii_.Get(
		j,
		"replicationSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) RootCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) TerminationProtectionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster) VersionReleaseSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystem",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/data-sources/advanced_cluster mongodbatlas_advanced_cluster} Data Source.
func NewDataMongodbatlasAdvancedCluster(scope constructs.Construct, id *string, config *DataMongodbatlasAdvancedClusterConfig) DataMongodbatlasAdvancedCluster {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasAdvancedClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasAdvancedCluster{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedCluster.DataMongodbatlasAdvancedCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/data-sources/advanced_cluster mongodbatlas_advanced_cluster} Data Source.
func NewDataMongodbatlasAdvancedCluster_Override(d DataMongodbatlasAdvancedCluster, scope constructs.Construct, id *string, config *DataMongodbatlasAdvancedClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedCluster.DataMongodbatlasAdvancedCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetPitEnabled(val interface{}) {
	if err := j.validateSetPitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitEnabled",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataMongodbatlasAdvancedCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasAdvancedCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedCluster.DataMongodbatlasAdvancedCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasAdvancedCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasAdvancedCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedCluster.DataMongodbatlasAdvancedCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasAdvancedCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasAdvancedCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedCluster.DataMongodbatlasAdvancedCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataMongodbatlasAdvancedCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedCluster.DataMongodbatlasAdvancedCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) ResetPitEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetPitEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

