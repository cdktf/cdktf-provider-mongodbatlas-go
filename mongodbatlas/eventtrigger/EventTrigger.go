// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/eventtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/event_trigger mongodbatlas_event_trigger}.
type EventTrigger interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigCollection() *string
	SetConfigCollection(val *string)
	ConfigCollectionInput() *string
	ConfigDatabase() *string
	SetConfigDatabase(val *string)
	ConfigDatabaseInput() *string
	ConfigFullDocument() interface{}
	SetConfigFullDocument(val interface{})
	ConfigFullDocumentBefore() interface{}
	SetConfigFullDocumentBefore(val interface{})
	ConfigFullDocumentBeforeInput() interface{}
	ConfigFullDocumentInput() interface{}
	ConfigMatch() *string
	SetConfigMatch(val *string)
	ConfigMatchInput() *string
	ConfigOperationType() *string
	SetConfigOperationType(val *string)
	ConfigOperationTypeInput() *string
	ConfigOperationTypes() *[]*string
	SetConfigOperationTypes(val *[]*string)
	ConfigOperationTypesInput() *[]*string
	ConfigProject() *string
	SetConfigProject(val *string)
	ConfigProjectInput() *string
	ConfigProviders() *[]*string
	SetConfigProviders(val *[]*string)
	ConfigProvidersInput() *[]*string
	ConfigSchedule() *string
	SetConfigSchedule(val *string)
	ConfigScheduleInput() *string
	ConfigScheduleType() *string
	ConfigServiceId() *string
	SetConfigServiceId(val *string)
	ConfigServiceIdInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EventProcessors() EventTriggerEventProcessorsOutputReference
	EventProcessorsInput() *EventTriggerEventProcessors
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionId() *string
	SetFunctionId(val *string)
	FunctionIdInput() *string
	FunctionName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TriggerId() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Unordered() interface{}
	SetUnordered(val interface{})
	UnorderedInput() interface{}
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
	PutEventProcessors(value *EventTriggerEventProcessors)
	ResetConfigCollection()
	ResetConfigDatabase()
	ResetConfigFullDocument()
	ResetConfigFullDocumentBefore()
	ResetConfigMatch()
	ResetConfigOperationType()
	ResetConfigOperationTypes()
	ResetConfigProject()
	ResetConfigProviders()
	ResetConfigSchedule()
	ResetConfigServiceId()
	ResetDisabled()
	ResetEventProcessors()
	ResetFunctionId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUnordered()
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

// The jsii proxy struct for EventTrigger
type jsiiProxy_EventTrigger struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EventTrigger) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigCollection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigCollectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configCollectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigFullDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configFullDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigFullDocumentBefore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configFullDocumentBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigFullDocumentBeforeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configFullDocumentBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigFullDocumentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configFullDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigOperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configOperationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigOperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configOperationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigOperationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configOperationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigOperationTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configOperationTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConfigServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) EventProcessors() EventTriggerEventProcessorsOutputReference {
	var returns EventTriggerEventProcessorsOutputReference
	_jsii_.Get(
		j,
		"eventProcessors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) EventProcessorsInput() *EventTriggerEventProcessors {
	var returns *EventTriggerEventProcessors
	_jsii_.Get(
		j,
		"eventProcessorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) FunctionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) TriggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) Unordered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unordered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTrigger) UnorderedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unorderedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/event_trigger mongodbatlas_event_trigger} Resource.
func NewEventTrigger(scope constructs.Construct, id *string, config *EventTriggerConfig) EventTrigger {
	_init_.Initialize()

	if err := validateNewEventTriggerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventTrigger{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/event_trigger mongodbatlas_event_trigger} Resource.
func NewEventTrigger_Override(e EventTrigger, scope constructs.Construct, id *string, config *EventTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventTrigger)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigCollection(val *string) {
	if err := j.validateSetConfigCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configCollection",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigDatabase(val *string) {
	if err := j.validateSetConfigDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configDatabase",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigFullDocument(val interface{}) {
	if err := j.validateSetConfigFullDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configFullDocument",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigFullDocumentBefore(val interface{}) {
	if err := j.validateSetConfigFullDocumentBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configFullDocumentBefore",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigMatch(val *string) {
	if err := j.validateSetConfigMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configMatch",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigOperationType(val *string) {
	if err := j.validateSetConfigOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configOperationType",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigOperationTypes(val *[]*string) {
	if err := j.validateSetConfigOperationTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configOperationTypes",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigProject(val *string) {
	if err := j.validateSetConfigProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configProject",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigProviders(val *[]*string) {
	if err := j.validateSetConfigProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configProviders",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigSchedule(val *string) {
	if err := j.validateSetConfigScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configSchedule",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConfigServiceId(val *string) {
	if err := j.validateSetConfigServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configServiceId",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetFunctionId(val *string) {
	if err := j.validateSetFunctionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionId",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventTrigger)SetUnordered(val interface{}) {
	if err := j.validateSetUnorderedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unordered",
		val,
	)
}

// Generates CDKTF code for importing a EventTrigger resource upon running "cdktf plan <stack-name>".
func EventTrigger_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEventTrigger_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
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
func EventTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventTrigger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventTrigger_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventTrigger_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventTrigger_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventTrigger_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EventTrigger_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventTrigger) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EventTrigger) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventTrigger) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EventTrigger) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventTrigger) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EventTrigger) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventTrigger) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventTrigger) PutEventProcessors(value *EventTriggerEventProcessors) {
	if err := e.validatePutEventProcessorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEventProcessors",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigCollection() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigCollection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigDatabase() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigDatabase",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigFullDocument() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigFullDocument",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigFullDocumentBefore() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigFullDocumentBefore",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigMatch() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigMatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigOperationType() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigOperationType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigOperationTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigOperationTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigProject() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigProject",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigProviders() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigProviders",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigSchedule() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigSchedule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetConfigServiceId() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigServiceId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetDisabled() {
	_jsii_.InvokeVoid(
		e,
		"resetDisabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetEventProcessors() {
	_jsii_.InvokeVoid(
		e,
		"resetEventProcessors",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetFunctionId() {
	_jsii_.InvokeVoid(
		e,
		"resetFunctionId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) ResetUnordered() {
	_jsii_.InvokeVoid(
		e,
		"resetUnordered",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTrigger) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTrigger) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

