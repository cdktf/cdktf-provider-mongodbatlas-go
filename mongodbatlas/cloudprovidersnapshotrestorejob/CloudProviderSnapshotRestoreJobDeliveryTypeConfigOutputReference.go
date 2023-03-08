package cloudprovidersnapshotrestorejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/cloudprovidersnapshotrestorejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference interface {
	cdktf.ComplexObject
	Automated() interface{}
	SetAutomated(val interface{})
	AutomatedInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Download() interface{}
	SetDownload(val interface{})
	DownloadInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CloudProviderSnapshotRestoreJobDeliveryTypeConfig
	SetInternalValue(val *CloudProviderSnapshotRestoreJobDeliveryTypeConfig)
	OplogInc() *float64
	SetOplogInc(val *float64)
	OplogIncInput() *float64
	OplogTs() *float64
	SetOplogTs(val *float64)
	OplogTsInput() *float64
	PointInTime() interface{}
	SetPointInTime(val interface{})
	PointInTimeInput() interface{}
	PointInTimeUtcSeconds() *float64
	SetPointInTimeUtcSeconds(val *float64)
	PointInTimeUtcSecondsInput() *float64
	TargetClusterName() *string
	SetTargetClusterName(val *string)
	TargetClusterNameInput() *string
	TargetProjectId() *string
	SetTargetProjectId(val *string)
	TargetProjectIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAutomated()
	ResetDownload()
	ResetOplogInc()
	ResetOplogTs()
	ResetPointInTime()
	ResetPointInTimeUtcSeconds()
	ResetTargetClusterName()
	ResetTargetProjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference
type jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) Automated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) AutomatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) Download() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"download",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) DownloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) InternalValue() *CloudProviderSnapshotRestoreJobDeliveryTypeConfig {
	var returns *CloudProviderSnapshotRestoreJobDeliveryTypeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) OplogInc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogInc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) OplogIncInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogIncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) OplogTs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogTs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) OplogTsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogTsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) PointInTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) PointInTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) PointInTimeUtcSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pointInTimeUtcSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) PointInTimeUtcSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pointInTimeUtcSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) TargetClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) TargetClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) TargetProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) TargetProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference_Override(c CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetAutomated(val interface{}) {
	if err := j.validateSetAutomatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automated",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetDownload(val interface{}) {
	if err := j.validateSetDownloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"download",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetInternalValue(val *CloudProviderSnapshotRestoreJobDeliveryTypeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetOplogInc(val *float64) {
	if err := j.validateSetOplogIncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplogInc",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetOplogTs(val *float64) {
	if err := j.validateSetOplogTsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplogTs",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetPointInTime(val interface{}) {
	if err := j.validateSetPointInTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTime",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetPointInTimeUtcSeconds(val *float64) {
	if err := j.validateSetPointInTimeUtcSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeUtcSeconds",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetTargetClusterName(val *string) {
	if err := j.validateSetTargetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetClusterName",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetTargetProjectId(val *string) {
	if err := j.validateSetTargetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProjectId",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetAutomated() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomated",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetDownload() {
	_jsii_.InvokeVoid(
		c,
		"resetDownload",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetOplogInc() {
	_jsii_.InvokeVoid(
		c,
		"resetOplogInc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetOplogTs() {
	_jsii_.InvokeVoid(
		c,
		"resetOplogTs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetPointInTime() {
	_jsii_.InvokeVoid(
		c,
		"resetPointInTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetPointInTimeUtcSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetPointInTimeUtcSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetTargetClusterName() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetClusterName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ResetTargetProjectId() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetProjectId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

