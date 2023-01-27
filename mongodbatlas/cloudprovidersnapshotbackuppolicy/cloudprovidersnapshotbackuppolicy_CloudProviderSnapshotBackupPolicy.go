package cloudprovidersnapshotbackuppolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/cloudprovidersnapshotbackuppolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_snapshot_backup_policy mongodbatlas_cloud_provider_snapshot_backup_policy}.
type CloudProviderSnapshotBackupPolicy interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
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
	NextSnapshot() *string
	// The tree node.
	Node() constructs.Node
	Policies() CloudProviderSnapshotBackupPolicyPoliciesList
	PoliciesInput() interface{}
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
	ReferenceHourOfDay() *float64
	SetReferenceHourOfDay(val *float64)
	ReferenceHourOfDayInput() *float64
	ReferenceMinuteOfHour() *float64
	SetReferenceMinuteOfHour(val *float64)
	ReferenceMinuteOfHourInput() *float64
	RestoreWindowDays() *float64
	SetRestoreWindowDays(val *float64)
	RestoreWindowDaysInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateSnapshots() interface{}
	SetUpdateSnapshots(val interface{})
	UpdateSnapshotsInput() interface{}
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
	PutPolicies(value interface{})
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReferenceHourOfDay()
	ResetReferenceMinuteOfHour()
	ResetRestoreWindowDays()
	ResetUpdateSnapshots()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudProviderSnapshotBackupPolicy
type jsiiProxy_CloudProviderSnapshotBackupPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) NextSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Policies() CloudProviderSnapshotBackupPolicyPoliciesList {
	var returns CloudProviderSnapshotBackupPolicyPoliciesList
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) PoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ReferenceHourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceHourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ReferenceHourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceHourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ReferenceMinuteOfHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceMinuteOfHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) ReferenceMinuteOfHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceMinuteOfHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) RestoreWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restoreWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) RestoreWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restoreWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) UpdateSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy) UpdateSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateSnapshotsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_snapshot_backup_policy mongodbatlas_cloud_provider_snapshot_backup_policy} Resource.
func NewCloudProviderSnapshotBackupPolicy(scope constructs.Construct, id *string, config *CloudProviderSnapshotBackupPolicyConfig) CloudProviderSnapshotBackupPolicy {
	_init_.Initialize()

	if err := validateNewCloudProviderSnapshotBackupPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudProviderSnapshotBackupPolicy{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_snapshot_backup_policy mongodbatlas_cloud_provider_snapshot_backup_policy} Resource.
func NewCloudProviderSnapshotBackupPolicy_Override(c CloudProviderSnapshotBackupPolicy, scope constructs.Construct, id *string, config *CloudProviderSnapshotBackupPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicy",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetReferenceHourOfDay(val *float64) {
	if err := j.validateSetReferenceHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceHourOfDay",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetReferenceMinuteOfHour(val *float64) {
	if err := j.validateSetReferenceMinuteOfHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceMinuteOfHour",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetRestoreWindowDays(val *float64) {
	if err := j.validateSetRestoreWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreWindowDays",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicy)SetUpdateSnapshots(val interface{}) {
	if err := j.validateSetUpdateSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSnapshots",
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
func CloudProviderSnapshotBackupPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudProviderSnapshotBackupPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudProviderSnapshotBackupPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudProviderSnapshotBackupPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudProviderSnapshotBackupPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudProviderSnapshotBackupPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudProviderSnapshotBackupPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) PutPolicies(value interface{}) {
	if err := c.validatePutPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPolicies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ResetReferenceHourOfDay() {
	_jsii_.InvokeVoid(
		c,
		"resetReferenceHourOfDay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ResetReferenceMinuteOfHour() {
	_jsii_.InvokeVoid(
		c,
		"resetReferenceMinuteOfHour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ResetRestoreWindowDays() {
	_jsii_.InvokeVoid(
		c,
		"resetRestoreWindowDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ResetUpdateSnapshots() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdateSnapshots",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

