// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/federateddatabaseinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference interface {
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
	InternalValue() *FederatedDatabaseInstanceStorageStoresReadPreference
	SetInternalValue(val *FederatedDatabaseInstanceStorageStoresReadPreference)
	MaxStalenessSeconds() *float64
	SetMaxStalenessSeconds(val *float64)
	MaxStalenessSecondsInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	TagSets() FederatedDatabaseInstanceStorageStoresReadPreferenceTagSetsList
	TagSetsInput() interface{}
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
	PutTagSets(value interface{})
	ResetMaxStalenessSeconds()
	ResetMode()
	ResetTagSets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference
type jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) InternalValue() *FederatedDatabaseInstanceStorageStoresReadPreference {
	var returns *FederatedDatabaseInstanceStorageStoresReadPreference
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) MaxStalenessSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStalenessSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) MaxStalenessSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStalenessSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) TagSets() FederatedDatabaseInstanceStorageStoresReadPreferenceTagSetsList {
	var returns FederatedDatabaseInstanceStorageStoresReadPreferenceTagSetsList
	_jsii_.Get(
		j,
		"tagSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) TagSetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference {
	_init_.Initialize()

	if err := validateNewFederatedDatabaseInstanceStorageStoresReadPreferenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedDatabaseInstance.FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference_Override(f FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedDatabaseInstance.FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetInternalValue(val *FederatedDatabaseInstanceStorageStoresReadPreference) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetMaxStalenessSeconds(val *float64) {
	if err := j.validateSetMaxStalenessSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStalenessSeconds",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) PutTagSets(value interface{}) {
	if err := f.validatePutTagSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTagSets",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ResetMaxStalenessSeconds() {
	_jsii_.InvokeVoid(
		f,
		"resetMaxStalenessSeconds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		f,
		"resetMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ResetTagSets() {
	_jsii_.InvokeVoid(
		f,
		"resetTagSets",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

