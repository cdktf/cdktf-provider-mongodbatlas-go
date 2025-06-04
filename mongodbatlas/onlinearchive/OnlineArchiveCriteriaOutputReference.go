// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinearchive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/onlinearchive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OnlineArchiveCriteriaOutputReference interface {
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
	DateField() *string
	SetDateField(val *string)
	DateFieldInput() *string
	DateFormat() *string
	SetDateFormat(val *string)
	DateFormatInput() *string
	ExpireAfterDays() *float64
	SetExpireAfterDays(val *float64)
	ExpireAfterDaysInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *OnlineArchiveCriteria
	SetInternalValue(val *OnlineArchiveCriteria)
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetDateField()
	ResetDateFormat()
	ResetExpireAfterDays()
	ResetQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OnlineArchiveCriteriaOutputReference
type jsiiProxy_OnlineArchiveCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) DateField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) DateFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) ExpireAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) ExpireAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) InternalValue() *OnlineArchiveCriteria {
	var returns *OnlineArchiveCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewOnlineArchiveCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OnlineArchiveCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewOnlineArchiveCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OnlineArchiveCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchiveCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOnlineArchiveCriteriaOutputReference_Override(o OnlineArchiveCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.onlineArchive.OnlineArchiveCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetDateField(val *string) {
	if err := j.validateSetDateFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateField",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetExpireAfterDays(val *float64) {
	if err := j.validateSetExpireAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireAfterDays",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetInternalValue(val *OnlineArchiveCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OnlineArchiveCriteriaOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) ResetDateField() {
	_jsii_.InvokeVoid(
		o,
		"resetDateField",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) ResetDateFormat() {
	_jsii_.InvokeVoid(
		o,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) ResetExpireAfterDays() {
	_jsii_.InvokeVoid(
		o,
		"resetExpireAfterDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) ResetQuery() {
	_jsii_.InvokeVoid(
		o,
		"resetQuery",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OnlineArchiveCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

