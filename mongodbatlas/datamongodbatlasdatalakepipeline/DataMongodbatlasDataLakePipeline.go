// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasdatalakepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/datamongodbatlasdatalakepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/data-sources/data_lake_pipeline mongodbatlas_data_lake_pipeline}.
type DataMongodbatlasDataLakePipeline interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedDate() *string
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
	IngestionSchedules() DataMongodbatlasDataLakePipelineIngestionSchedulesList
	LastUpdatedDate() *string
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
	RawOverrides() interface{}
	Sink() DataMongodbatlasDataLakePipelineSinkList
	Snapshots() DataMongodbatlasDataLakePipelineSnapshotsList
	Source() DataMongodbatlasDataLakePipelineSourceList
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Transformations() DataMongodbatlasDataLakePipelineTransformationsList
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

// The jsii proxy struct for DataMongodbatlasDataLakePipeline
type jsiiProxy_DataMongodbatlasDataLakePipeline struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) IngestionSchedules() DataMongodbatlasDataLakePipelineIngestionSchedulesList {
	var returns DataMongodbatlasDataLakePipelineIngestionSchedulesList
	_jsii_.Get(
		j,
		"ingestionSchedules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Sink() DataMongodbatlasDataLakePipelineSinkList {
	var returns DataMongodbatlasDataLakePipelineSinkList
	_jsii_.Get(
		j,
		"sink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Snapshots() DataMongodbatlasDataLakePipelineSnapshotsList {
	var returns DataMongodbatlasDataLakePipelineSnapshotsList
	_jsii_.Get(
		j,
		"snapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Source() DataMongodbatlasDataLakePipelineSourceList {
	var returns DataMongodbatlasDataLakePipelineSourceList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline) Transformations() DataMongodbatlasDataLakePipelineTransformationsList {
	var returns DataMongodbatlasDataLakePipelineTransformationsList
	_jsii_.Get(
		j,
		"transformations",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/data-sources/data_lake_pipeline mongodbatlas_data_lake_pipeline} Data Source.
func NewDataMongodbatlasDataLakePipeline(scope constructs.Construct, id *string, config *DataMongodbatlasDataLakePipelineConfig) DataMongodbatlasDataLakePipeline {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasDataLakePipelineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasDataLakePipeline{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/data-sources/data_lake_pipeline mongodbatlas_data_lake_pipeline} Data Source.
func NewDataMongodbatlasDataLakePipeline_Override(d DataMongodbatlasDataLakePipeline, scope constructs.Construct, id *string, config *DataMongodbatlasDataLakePipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasDataLakePipeline)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataMongodbatlasDataLakePipeline resource upon running "cdktf plan <stack-name>".
func DataMongodbatlasDataLakePipeline_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataMongodbatlasDataLakePipeline_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
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
func DataMongodbatlasDataLakePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasDataLakePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasDataLakePipeline_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasDataLakePipeline_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasDataLakePipeline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasDataLakePipeline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataMongodbatlasDataLakePipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasDataLakePipeline.DataMongodbatlasDataLakePipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasDataLakePipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

