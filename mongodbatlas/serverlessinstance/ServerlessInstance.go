// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serverlessinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/serverlessinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/resources/serverless_instance mongodbatlas_serverless_instance}.
type ServerlessInstance interface {
	cdktf.TerraformResource
	AutoIndexing() interface{}
	SetAutoIndexing(val interface{})
	AutoIndexingInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionStringsPrivateEndpointSrv() *[]*string
	ConnectionStringsStandardSrv() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContinuousBackupEnabled() interface{}
	SetContinuousBackupEnabled(val interface{})
	ContinuousBackupEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateDate() *string
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
	Links() ServerlessInstanceLinksList
	LinksInput() interface{}
	MongoDbVersion() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderSettingsBackingProviderName() *string
	SetProviderSettingsBackingProviderName(val *string)
	ProviderSettingsBackingProviderNameInput() *string
	ProviderSettingsProviderName() *string
	SetProviderSettingsProviderName(val *string)
	ProviderSettingsProviderNameInput() *string
	ProviderSettingsRegionName() *string
	SetProviderSettingsRegionName(val *string)
	ProviderSettingsRegionNameInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	StateName() *string
	SetStateName(val *string)
	StateNameInput() *string
	Tags() ServerlessInstanceTagsList
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
	PutLinks(value interface{})
	PutTags(value interface{})
	ResetAutoIndexing()
	ResetContinuousBackupEnabled()
	ResetId()
	ResetLinks()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStateName()
	ResetTags()
	ResetTerminationProtectionEnabled()
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

// The jsii proxy struct for ServerlessInstance
type jsiiProxy_ServerlessInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServerlessInstance) AutoIndexing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoIndexing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) AutoIndexingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoIndexingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ConnectionStringsPrivateEndpointSrv() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionStringsPrivateEndpointSrv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ConnectionStringsStandardSrv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringsStandardSrv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ContinuousBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ContinuousBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Links() ServerlessInstanceLinksList {
	var returns ServerlessInstanceLinksList
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) LinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProviderSettingsBackingProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSettingsBackingProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProviderSettingsBackingProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSettingsBackingProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProviderSettingsProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSettingsProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProviderSettingsProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSettingsProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProviderSettingsRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSettingsRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) ProviderSettingsRegionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerSettingsRegionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) StateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) Tags() ServerlessInstanceTagsList {
	var returns ServerlessInstanceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) TerminationProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) TerminationProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/resources/serverless_instance mongodbatlas_serverless_instance} Resource.
func NewServerlessInstance(scope constructs.Construct, id *string, config *ServerlessInstanceConfig) ServerlessInstance {
	_init_.Initialize()

	if err := validateNewServerlessInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServerlessInstance{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.32.0/docs/resources/serverless_instance mongodbatlas_serverless_instance} Resource.
func NewServerlessInstance_Override(s ServerlessInstance, scope constructs.Construct, id *string, config *ServerlessInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetAutoIndexing(val interface{}) {
	if err := j.validateSetAutoIndexingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoIndexing",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetContinuousBackupEnabled(val interface{}) {
	if err := j.validateSetContinuousBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continuousBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetProviderSettingsBackingProviderName(val *string) {
	if err := j.validateSetProviderSettingsBackingProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerSettingsBackingProviderName",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetProviderSettingsProviderName(val *string) {
	if err := j.validateSetProviderSettingsProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerSettingsProviderName",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetProviderSettingsRegionName(val *string) {
	if err := j.validateSetProviderSettingsRegionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerSettingsRegionName",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetStateName(val *string) {
	if err := j.validateSetStateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateName",
		val,
	)
}

func (j *jsiiProxy_ServerlessInstance)SetTerminationProtectionEnabled(val interface{}) {
	if err := j.validateSetTerminationProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtectionEnabled",
		val,
	)
}

// Generates CDKTF code for importing a ServerlessInstance resource upon running "cdktf plan <stack-name>".
func ServerlessInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServerlessInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
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
func ServerlessInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServerlessInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServerlessInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServerlessInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServerlessInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServerlessInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServerlessInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.serverlessInstance.ServerlessInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServerlessInstance) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServerlessInstance) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServerlessInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServerlessInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServerlessInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServerlessInstance) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServerlessInstance) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServerlessInstance) PutLinks(value interface{}) {
	if err := s.validatePutLinksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLinks",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServerlessInstance) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetAutoIndexing() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoIndexing",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetContinuousBackupEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetContinuousBackupEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetLinks() {
	_jsii_.InvokeVoid(
		s,
		"resetLinks",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetStateName() {
	_jsii_.InvokeVoid(
		s,
		"resetStateName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) ResetTerminationProtectionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetTerminationProtectionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

