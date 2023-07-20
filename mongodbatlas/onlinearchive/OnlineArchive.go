package onlinearchive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/onlinearchive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/online_archive mongodbatlas_online_archive}.
type OnlineArchive interface {
	cdktf.TerraformResource
	ArchiveId() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	CollectionType() *string
	SetCollectionType(val *string)
	CollectionTypeInput() *string
	CollName() *string
	SetCollName(val *string)
	CollNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Criteria() OnlineArchiveCriteriaOutputReference
	CriteriaInput() *OnlineArchiveCriteria
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PartitionFields() OnlineArchivePartitionFieldsList
	PartitionFieldsInput() interface{}
	Paused() interface{}
	SetPaused(val interface{})
	PausedInput() interface{}
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
	Schedule() OnlineArchiveScheduleOutputReference
	ScheduleInput() *OnlineArchiveSchedule
	State() *string
	SyncCreation() interface{}
	SetSyncCreation(val interface{})
	SyncCreationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutCriteria(value *OnlineArchiveCriteria)
	PutPartitionFields(value interface{})
	PutSchedule(value *OnlineArchiveSchedule)
	ResetCollectionType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionFields()
	ResetPaused()
	ResetSchedule()
	ResetSyncCreation()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OnlineArchive
type jsiiProxy_OnlineArchive struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OnlineArchive) ArchiveId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) CollectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) CollectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) CollName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) CollNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Criteria() OnlineArchiveCriteriaOutputReference {
	var returns OnlineArchiveCriteriaOutputReference
	_jsii_.Get(
		j,
		"criteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) CriteriaInput() *OnlineArchiveCriteria {
	var returns *OnlineArchiveCriteria
	_jsii_.Get(
		j,
		"criteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) PartitionFields() OnlineArchivePartitionFieldsList {
	var returns OnlineArchivePartitionFieldsList
	_jsii_.Get(
		j,
		"partitionFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) PartitionFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Paused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) PausedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) Schedule() OnlineArchiveScheduleOutputReference {
	var returns OnlineArchiveScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) ScheduleInput() *OnlineArchiveSchedule {
	var returns *OnlineArchiveSchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) SyncCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) SyncCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchive) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/online_archive mongodbatlas_online_archive} Resource.
func NewOnlineArchive(scope constructs.Construct, id *string, config *OnlineArchiveConfig) OnlineArchive {
	_init_.Initialize()

	if err := validateNewOnlineArchiveParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnlineArchive{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchive",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/online_archive mongodbatlas_online_archive} Resource.
func NewOnlineArchive_Override(o OnlineArchive, scope constructs.Construct, id *string, config *OnlineArchiveConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchive",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OnlineArchive)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetCollectionType(val *string) {
	if err := j.validateSetCollectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionType",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetCollName(val *string) {
	if err := j.validateSetCollNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collName",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetPaused(val interface{}) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OnlineArchive)SetSyncCreation(val interface{}) {
	if err := j.validateSetSyncCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncCreation",
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
func OnlineArchive_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOnlineArchive_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchive",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OnlineArchive_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOnlineArchive_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchive",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OnlineArchive_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOnlineArchive_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchive",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OnlineArchive_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchive",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OnlineArchive) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OnlineArchive) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OnlineArchive) PutCriteria(value *OnlineArchiveCriteria) {
	if err := o.validatePutCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCriteria",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnlineArchive) PutPartitionFields(value interface{}) {
	if err := o.validatePutPartitionFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPartitionFields",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnlineArchive) PutSchedule(value *OnlineArchiveSchedule) {
	if err := o.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSchedule",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OnlineArchive) ResetCollectionType() {
	_jsii_.InvokeVoid(
		o,
		"resetCollectionType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) ResetPartitionFields() {
	_jsii_.InvokeVoid(
		o,
		"resetPartitionFields",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) ResetPaused() {
	_jsii_.InvokeVoid(
		o,
		"resetPaused",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) ResetSchedule() {
	_jsii_.InvokeVoid(
		o,
		"resetSchedule",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) ResetSyncCreation() {
	_jsii_.InvokeVoid(
		o,
		"resetSyncCreation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchive) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchive) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

