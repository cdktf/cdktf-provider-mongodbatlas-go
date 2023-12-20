// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v6/project/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/project mongodbatlas_project}.
type Project interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterCount() *float64
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
	Created() *string
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
	IsCollectDatabaseSpecificsStatisticsEnabled() interface{}
	SetIsCollectDatabaseSpecificsStatisticsEnabled(val interface{})
	IsCollectDatabaseSpecificsStatisticsEnabledInput() interface{}
	IsDataExplorerEnabled() interface{}
	SetIsDataExplorerEnabled(val interface{})
	IsDataExplorerEnabledInput() interface{}
	IsExtendedStorageSizesEnabled() interface{}
	SetIsExtendedStorageSizesEnabled(val interface{})
	IsExtendedStorageSizesEnabledInput() interface{}
	IsPerformanceAdvisorEnabled() interface{}
	SetIsPerformanceAdvisorEnabled(val interface{})
	IsPerformanceAdvisorEnabledInput() interface{}
	IsRealtimePerformancePanelEnabled() interface{}
	SetIsRealtimePerformancePanelEnabled(val interface{})
	IsRealtimePerformancePanelEnabledInput() interface{}
	IsSchemaAdvisorEnabled() interface{}
	SetIsSchemaAdvisorEnabled(val interface{})
	IsSchemaAdvisorEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Limits() ProjectLimitsList
	LimitsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OrgId() *string
	SetOrgId(val *string)
	OrgIdInput() *string
	ProjectOwnerId() *string
	SetProjectOwnerId(val *string)
	ProjectOwnerIdInput() *string
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
	RegionUsageRestrictions() *string
	SetRegionUsageRestrictions(val *string)
	RegionUsageRestrictionsInput() *string
	Teams() ProjectTeamsList
	TeamsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithDefaultAlertsSettings() interface{}
	SetWithDefaultAlertsSettings(val interface{})
	WithDefaultAlertsSettingsInput() interface{}
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
	PutLimits(value interface{})
	PutTeams(value interface{})
	ResetIsCollectDatabaseSpecificsStatisticsEnabled()
	ResetIsDataExplorerEnabled()
	ResetIsExtendedStorageSizesEnabled()
	ResetIsPerformanceAdvisorEnabled()
	ResetIsRealtimePerformancePanelEnabled()
	ResetIsSchemaAdvisorEnabled()
	ResetLimits()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProjectOwnerId()
	ResetRegionUsageRestrictions()
	ResetTeams()
	ResetWithDefaultAlertsSettings()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Project) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsCollectDatabaseSpecificsStatisticsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCollectDatabaseSpecificsStatisticsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsCollectDatabaseSpecificsStatisticsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCollectDatabaseSpecificsStatisticsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsDataExplorerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDataExplorerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsDataExplorerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDataExplorerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsExtendedStorageSizesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExtendedStorageSizesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsExtendedStorageSizesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isExtendedStorageSizesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsPerformanceAdvisorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPerformanceAdvisorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsPerformanceAdvisorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPerformanceAdvisorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsRealtimePerformancePanelEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRealtimePerformancePanelEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsRealtimePerformancePanelEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRealtimePerformancePanelEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsSchemaAdvisorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSchemaAdvisorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) IsSchemaAdvisorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSchemaAdvisorEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Limits() ProjectLimitsList {
	var returns ProjectLimitsList
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) LimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectOwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectOwnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RegionUsageRestrictions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionUsageRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) RegionUsageRestrictionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionUsageRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Teams() ProjectTeamsList {
	var returns ProjectTeamsList
	_jsii_.Get(
		j,
		"teams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TeamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"teamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WithDefaultAlertsSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDefaultAlertsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) WithDefaultAlertsSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDefaultAlertsSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/project mongodbatlas_project} Resource.
func NewProject(scope constructs.Construct, id *string, config *ProjectConfig) Project {
	_init_.Initialize()

	if err := validateNewProjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Project{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.project.Project",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.14.0/docs/resources/project mongodbatlas_project} Resource.
func NewProject_Override(p Project, scope constructs.Construct, id *string, config *ProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.project.Project",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_Project)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Project)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Project)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Project)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Project)SetIsCollectDatabaseSpecificsStatisticsEnabled(val interface{}) {
	if err := j.validateSetIsCollectDatabaseSpecificsStatisticsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCollectDatabaseSpecificsStatisticsEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetIsDataExplorerEnabled(val interface{}) {
	if err := j.validateSetIsDataExplorerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDataExplorerEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetIsExtendedStorageSizesEnabled(val interface{}) {
	if err := j.validateSetIsExtendedStorageSizesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isExtendedStorageSizesEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetIsPerformanceAdvisorEnabled(val interface{}) {
	if err := j.validateSetIsPerformanceAdvisorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPerformanceAdvisorEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetIsRealtimePerformancePanelEnabled(val interface{}) {
	if err := j.validateSetIsRealtimePerformancePanelEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRealtimePerformancePanelEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetIsSchemaAdvisorEnabled(val interface{}) {
	if err := j.validateSetIsSchemaAdvisorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSchemaAdvisorEnabled",
		val,
	)
}

func (j *jsiiProxy_Project)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Project)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Project)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_Project)SetProjectOwnerId(val *string) {
	if err := j.validateSetProjectOwnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectOwnerId",
		val,
	)
}

func (j *jsiiProxy_Project)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Project)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Project)SetRegionUsageRestrictions(val *string) {
	if err := j.validateSetRegionUsageRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionUsageRestrictions",
		val,
	)
}

func (j *jsiiProxy_Project)SetWithDefaultAlertsSettings(val interface{}) {
	if err := j.validateSetWithDefaultAlertsSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withDefaultAlertsSettings",
		val,
	)
}

// Generates CDKTF code for importing a Project resource upon running "cdktf plan <stack-name>".
func Project_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateProject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.project.Project",
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
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.project.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Project_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.project.Project",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Project_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.project.Project",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Project_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.project.Project",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Project) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_Project) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_Project) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_Project) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Project) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_Project) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_Project) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_Project) PutLimits(value interface{}) {
	if err := p.validatePutLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLimits",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Project) PutTeams(value interface{}) {
	if err := p.validatePutTeamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTeams",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_Project) ResetIsCollectDatabaseSpecificsStatisticsEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsCollectDatabaseSpecificsStatisticsEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIsDataExplorerEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsDataExplorerEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIsExtendedStorageSizesEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsExtendedStorageSizesEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIsPerformanceAdvisorEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsPerformanceAdvisorEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIsRealtimePerformancePanelEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsRealtimePerformancePanelEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetIsSchemaAdvisorEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetIsSchemaAdvisorEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetLimits() {
	_jsii_.InvokeVoid(
		p,
		"resetLimits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetProjectOwnerId() {
	_jsii_.InvokeVoid(
		p,
		"resetProjectOwnerId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetRegionUsageRestrictions() {
	_jsii_.InvokeVoid(
		p,
		"resetRegionUsageRestrictions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetTeams() {
	_jsii_.InvokeVoid(
		p,
		"resetTeams",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ResetWithDefaultAlertsSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetWithDefaultAlertsSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

