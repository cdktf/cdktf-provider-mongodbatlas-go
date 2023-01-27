package cloudprovidersnapshotbackuppolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/cloudprovidersnapshotbackuppolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	FrequencyInterval() *float64
	SetFrequencyInterval(val *float64)
	FrequencyIntervalInput() *float64
	FrequencyType() *string
	SetFrequencyType(val *string)
	FrequencyTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RetentionUnit() *string
	SetRetentionUnit(val *string)
	RetentionUnitInput() *string
	RetentionValue() *float64
	SetRetentionValue(val *float64)
	RetentionValueInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference
type jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) FrequencyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) FrequencyIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) FrequencyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) FrequencyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) RetentionUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) RetentionUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) RetentionValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) RetentionValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference {
	_init_.Initialize()

	if err := validateNewCloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference_Override(c CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotBackupPolicy.CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetFrequencyInterval(val *float64) {
	if err := j.validateSetFrequencyIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequencyInterval",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetFrequencyType(val *string) {
	if err := j.validateSetFrequencyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequencyType",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetRetentionUnit(val *string) {
	if err := j.validateSetRetentionUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionUnit",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetRetentionValue(val *float64) {
	if err := j.validateSetRetentionValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionValue",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudProviderSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

