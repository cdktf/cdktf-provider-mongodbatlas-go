// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenancewindow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/maintenancewindow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/maintenance_window mongodbatlas_maintenance_window}.
type MaintenanceWindow interface {
	cdktf.TerraformResource
	AutoDefer() interface{}
	SetAutoDefer(val interface{})
	AutoDeferInput() interface{}
	AutoDeferOnceEnabled() interface{}
	SetAutoDeferOnceEnabled(val interface{})
	AutoDeferOnceEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DayOfWeek() *float64
	SetDayOfWeek(val *float64)
	DayOfWeekInput() *float64
	Defer() interface{}
	SetDefer(val interface{})
	DeferInput() interface{}
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
	HourOfDay() *float64
	SetHourOfDay(val *float64)
	HourOfDayInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NumberOfDeferrals() *float64
	SetNumberOfDeferrals(val *float64)
	NumberOfDeferralsInput() *float64
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
	StartAsap() cdktf.IResolvable
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
	ResetAutoDefer()
	ResetAutoDeferOnceEnabled()
	ResetDayOfWeek()
	ResetDefer()
	ResetHourOfDay()
	ResetId()
	ResetNumberOfDeferrals()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for MaintenanceWindow
type jsiiProxy_MaintenanceWindow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MaintenanceWindow) AutoDefer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDefer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) AutoDeferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) AutoDeferOnceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeferOnceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) AutoDeferOnceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeferOnceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) DayOfWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) DayOfWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Defer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) DeferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) HourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) HourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) NumberOfDeferrals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfDeferrals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) NumberOfDeferralsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfDeferralsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) StartAsap() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"startAsap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceWindow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/maintenance_window mongodbatlas_maintenance_window} Resource.
func NewMaintenanceWindow(scope constructs.Construct, id *string, config *MaintenanceWindowConfig) MaintenanceWindow {
	_init_.Initialize()

	if err := validateNewMaintenanceWindowParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MaintenanceWindow{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/maintenance_window mongodbatlas_maintenance_window} Resource.
func NewMaintenanceWindow_Override(m MaintenanceWindow, scope constructs.Construct, id *string, config *MaintenanceWindowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetAutoDefer(val interface{}) {
	if err := j.validateSetAutoDeferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDefer",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetAutoDeferOnceEnabled(val interface{}) {
	if err := j.validateSetAutoDeferOnceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeferOnceEnabled",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetDayOfWeek(val *float64) {
	if err := j.validateSetDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetDefer(val interface{}) {
	if err := j.validateSetDeferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defer",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetHourOfDay(val *float64) {
	if err := j.validateSetHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hourOfDay",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetNumberOfDeferrals(val *float64) {
	if err := j.validateSetNumberOfDeferralsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfDeferrals",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MaintenanceWindow)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a MaintenanceWindow resource upon running "cdktf plan <stack-name>".
func MaintenanceWindow_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMaintenanceWindow_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
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
func MaintenanceWindow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMaintenanceWindow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MaintenanceWindow_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMaintenanceWindow_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MaintenanceWindow_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMaintenanceWindow_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MaintenanceWindow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.maintenanceWindow.MaintenanceWindow",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MaintenanceWindow) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MaintenanceWindow) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MaintenanceWindow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MaintenanceWindow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MaintenanceWindow) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MaintenanceWindow) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MaintenanceWindow) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetAutoDefer() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoDefer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetAutoDeferOnceEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoDeferOnceEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetDayOfWeek() {
	_jsii_.InvokeVoid(
		m,
		"resetDayOfWeek",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetDefer() {
	_jsii_.InvokeVoid(
		m,
		"resetDefer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetHourOfDay() {
	_jsii_.InvokeVoid(
		m,
		"resetHourOfDay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetNumberOfDeferrals() {
	_jsii_.InvokeVoid(
		m,
		"resetNumberOfDeferrals",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceWindow) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceWindow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

