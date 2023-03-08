package cloudprovidersnapshotrestorejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/cloudprovidersnapshotrestorejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_snapshot_restore_job mongodbatlas_cloud_provider_snapshot_restore_job}.
type CloudProviderSnapshotRestoreJob interface {
	cdktf.TerraformResource
	Cancelled() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedAt() *string
	DeliveryType() *map[string]*string
	SetDeliveryType(val *map[string]*string)
	DeliveryTypeConfig() CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference
	DeliveryTypeConfigInput() *CloudProviderSnapshotRestoreJobDeliveryTypeConfig
	DeliveryTypeInput() *map[string]*string
	DeliveryUrl() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Expired() cdktf.IResolvable
	ExpiresAt() *string
	FinishedAt() *string
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
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	SnapshotRestoreJobId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timestamp() *string
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
	PutDeliveryTypeConfig(value *CloudProviderSnapshotRestoreJobDeliveryTypeConfig)
	ResetDeliveryType()
	ResetDeliveryTypeConfig()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudProviderSnapshotRestoreJob
type jsiiProxy_CloudProviderSnapshotRestoreJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Cancelled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cancelled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) DeliveryType() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deliveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) DeliveryTypeConfig() CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference {
	var returns CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference
	_jsii_.Get(
		j,
		"deliveryTypeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) DeliveryTypeConfigInput() *CloudProviderSnapshotRestoreJobDeliveryTypeConfig {
	var returns *CloudProviderSnapshotRestoreJobDeliveryTypeConfig
	_jsii_.Get(
		j,
		"deliveryTypeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) DeliveryTypeInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"deliveryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) DeliveryUrl() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deliveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Expired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"expired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) FinishedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) SnapshotRestoreJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotRestoreJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob) Timestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestamp",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_snapshot_restore_job mongodbatlas_cloud_provider_snapshot_restore_job} Resource.
func NewCloudProviderSnapshotRestoreJob(scope constructs.Construct, id *string, config *CloudProviderSnapshotRestoreJobConfig) CloudProviderSnapshotRestoreJob {
	_init_.Initialize()

	if err := validateNewCloudProviderSnapshotRestoreJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudProviderSnapshotRestoreJob{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_snapshot_restore_job mongodbatlas_cloud_provider_snapshot_restore_job} Resource.
func NewCloudProviderSnapshotRestoreJob_Override(c CloudProviderSnapshotRestoreJob, scope constructs.Construct, id *string, config *CloudProviderSnapshotRestoreJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetDeliveryType(val *map[string]*string) {
	if err := j.validateSetDeliveryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryType",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJob)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
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
func CloudProviderSnapshotRestoreJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudProviderSnapshotRestoreJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudProviderSnapshotRestoreJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudProviderSnapshotRestoreJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudProviderSnapshotRestoreJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudProviderSnapshotRestoreJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudProviderSnapshotRestoreJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) PutDeliveryTypeConfig(value *CloudProviderSnapshotRestoreJobDeliveryTypeConfig) {
	if err := c.validatePutDeliveryTypeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeliveryTypeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ResetDeliveryType() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ResetDeliveryTypeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDeliveryTypeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

