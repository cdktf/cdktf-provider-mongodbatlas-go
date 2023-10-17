// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v6/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v6/cloudbackupschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudBackupScheduleCopySettingsOutputReference interface {
	cdktf.ComplexObject
	CloudProvider() *string
	SetCloudProvider(val *string)
	CloudProviderInput() *string
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
	Frequencies() *[]*string
	SetFrequencies(val *[]*string)
	FrequenciesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RegionName() *string
	SetRegionName(val *string)
	RegionNameInput() *string
	ReplicationSpecId() *string
	SetReplicationSpecId(val *string)
	ReplicationSpecIdInput() *string
	ShouldCopyOplogs() interface{}
	SetShouldCopyOplogs(val interface{})
	ShouldCopyOplogsInput() interface{}
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
	ResetCloudProvider()
	ResetFrequencies()
	ResetRegionName()
	ResetReplicationSpecId()
	ResetShouldCopyOplogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudBackupScheduleCopySettingsOutputReference
type jsiiProxy_CloudBackupScheduleCopySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) CloudProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) CloudProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) Frequencies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frequencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) FrequenciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"frequenciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) RegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) RegionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ReplicationSpecId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSpecId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ReplicationSpecIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSpecIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ShouldCopyOplogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldCopyOplogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ShouldCopyOplogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shouldCopyOplogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudBackupScheduleCopySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudBackupScheduleCopySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudBackupScheduleCopySettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudBackupScheduleCopySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleCopySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudBackupScheduleCopySettingsOutputReference_Override(c CloudBackupScheduleCopySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleCopySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetCloudProvider(val *string) {
	if err := j.validateSetCloudProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudProvider",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetFrequencies(val *[]*string) {
	if err := j.validateSetFrequenciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequencies",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetRegionName(val *string) {
	if err := j.validateSetRegionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionName",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetReplicationSpecId(val *string) {
	if err := j.validateSetReplicationSpecIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSpecId",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetShouldCopyOplogs(val interface{}) {
	if err := j.validateSetShouldCopyOplogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shouldCopyOplogs",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ResetCloudProvider() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudProvider",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ResetFrequencies() {
	_jsii_.InvokeVoid(
		c,
		"resetFrequencies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ResetRegionName() {
	_jsii_.InvokeVoid(
		c,
		"resetRegionName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ResetReplicationSpecId() {
	_jsii_.InvokeVoid(
		c,
		"resetReplicationSpecId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ResetShouldCopyOplogs() {
	_jsii_.InvokeVoid(
		c,
		"resetShouldCopyOplogs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudBackupScheduleCopySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

