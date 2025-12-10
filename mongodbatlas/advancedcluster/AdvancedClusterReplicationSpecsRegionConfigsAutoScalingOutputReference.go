// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/advancedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference interface {
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
	ComputeEnabled() interface{}
	SetComputeEnabled(val interface{})
	ComputeEnabledInput() interface{}
	ComputeMaxInstanceSize() *string
	SetComputeMaxInstanceSize(val *string)
	ComputeMaxInstanceSizeInput() *string
	ComputeMinInstanceSize() *string
	SetComputeMinInstanceSize(val *string)
	ComputeMinInstanceSizeInput() *string
	ComputeScaleDownEnabled() interface{}
	SetComputeScaleDownEnabled(val interface{})
	ComputeScaleDownEnabledInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskGbEnabled() interface{}
	SetDiskGbEnabled(val interface{})
	DiskGbEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling
	SetInternalValue(val *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetComputeEnabled()
	ResetComputeMaxInstanceSize()
	ResetComputeMinInstanceSize()
	ResetComputeScaleDownEnabled()
	ResetDiskGbEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference
type jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeMaxInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeMaxInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeMaxInstanceSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeMaxInstanceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeMinInstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeMinInstanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeMinInstanceSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeMinInstanceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeScaleDownEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeScaleDownEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeScaleDownEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeScaleDownEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) DiskGbEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskGbEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) DiskGbEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskGbEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) InternalValue() *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewAdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference_Override(a AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetComputeEnabled(val interface{}) {
	if err := j.validateSetComputeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetComputeMaxInstanceSize(val *string) {
	if err := j.validateSetComputeMaxInstanceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMaxInstanceSize",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetComputeMinInstanceSize(val *string) {
	if err := j.validateSetComputeMinInstanceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeMinInstanceSize",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetComputeScaleDownEnabled(val interface{}) {
	if err := j.validateSetComputeScaleDownEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeScaleDownEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetDiskGbEnabled(val interface{}) {
	if err := j.validateSetDiskGbEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskGbEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetInternalValue(val *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ResetComputeEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ResetComputeMaxInstanceSize() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeMaxInstanceSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ResetComputeMinInstanceSize() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeMinInstanceSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ResetComputeScaleDownEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeScaleDownEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ResetDiskGbEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDiskGbEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

