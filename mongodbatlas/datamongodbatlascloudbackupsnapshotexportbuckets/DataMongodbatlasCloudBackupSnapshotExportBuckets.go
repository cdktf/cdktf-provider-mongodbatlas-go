package datamongodbatlascloudbackupsnapshotexportbuckets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/datamongodbatlascloudbackupsnapshotexportbuckets/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets mongodbatlas_cloud_backup_snapshot_export_buckets}.
type DataMongodbatlasCloudBackupSnapshotExportBuckets interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	ItemsPerPage() *float64
	SetItemsPerPage(val *float64)
	ItemsPerPageInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PageNum() *float64
	SetPageNum(val *float64)
	PageNumInput() *float64
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Results() DataMongodbatlasCloudBackupSnapshotExportBucketsResultsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TotalCount() *float64
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
	ResetItemsPerPage()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPageNum()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataMongodbatlasCloudBackupSnapshotExportBuckets
type jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ItemsPerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"itemsPerPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ItemsPerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"itemsPerPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) PageNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) PageNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) Results() DataMongodbatlasCloudBackupSnapshotExportBucketsResultsList {
	var returns DataMongodbatlasCloudBackupSnapshotExportBucketsResultsList
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) TotalCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalCount",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets mongodbatlas_cloud_backup_snapshot_export_buckets} Data Source.
func NewDataMongodbatlasCloudBackupSnapshotExportBuckets(scope constructs.Construct, id *string, config *DataMongodbatlasCloudBackupSnapshotExportBucketsConfig) DataMongodbatlasCloudBackupSnapshotExportBuckets {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasCloudBackupSnapshotExportBucketsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotExportBuckets.DataMongodbatlasCloudBackupSnapshotExportBuckets",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/data-sources/cloud_backup_snapshot_export_buckets mongodbatlas_cloud_backup_snapshot_export_buckets} Data Source.
func NewDataMongodbatlasCloudBackupSnapshotExportBuckets_Override(d DataMongodbatlasCloudBackupSnapshotExportBuckets, scope constructs.Construct, id *string, config *DataMongodbatlasCloudBackupSnapshotExportBucketsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotExportBuckets.DataMongodbatlasCloudBackupSnapshotExportBuckets",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetItemsPerPage(val *float64) {
	if err := j.validateSetItemsPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"itemsPerPage",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetPageNum(val *float64) {
	if err := j.validateSetPageNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pageNum",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets)SetProvider(val cdktf.TerraformProvider) {
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
func DataMongodbatlasCloudBackupSnapshotExportBuckets_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasCloudBackupSnapshotExportBuckets_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotExportBuckets.DataMongodbatlasCloudBackupSnapshotExportBuckets",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasCloudBackupSnapshotExportBuckets_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasCloudBackupSnapshotExportBuckets_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotExportBuckets.DataMongodbatlasCloudBackupSnapshotExportBuckets",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasCloudBackupSnapshotExportBuckets_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasCloudBackupSnapshotExportBuckets_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotExportBuckets.DataMongodbatlasCloudBackupSnapshotExportBuckets",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataMongodbatlasCloudBackupSnapshotExportBuckets_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotExportBuckets.DataMongodbatlasCloudBackupSnapshotExportBuckets",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ResetItemsPerPage() {
	_jsii_.InvokeVoid(
		d,
		"resetItemsPerPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ResetPageNum() {
	_jsii_.InvokeVoid(
		d,
		"resetPageNum",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotExportBuckets) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

