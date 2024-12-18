// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasclusters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/datamongodbatlasclusters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() *float64
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
	DefaultMaxTimeMs() *float64
	DefaultReadConcern() *string
	DefaultWriteConcern() *string
	FailIndexKeyTooLong() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataMongodbatlasClustersResultsAdvancedConfiguration
	SetInternalValue(val *DataMongodbatlasClustersResultsAdvancedConfiguration)
	JavascriptEnabled() cdktf.IResolvable
	MinimumEnabledTlsProtocol() *string
	NoTableScan() cdktf.IResolvable
	OplogMinRetentionHours() *float64
	OplogSizeMb() *float64
	SampleRefreshIntervalBiConnector() *float64
	SampleSizeBiConnector() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransactionLifetimeLimitSeconds() *float64
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

// The jsii proxy struct for DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference
type jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeStreamOptionsPreAndPostImagesExpireAfterSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) DefaultMaxTimeMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxTimeMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) DefaultReadConcern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReadConcern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) DefaultWriteConcern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWriteConcern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) FailIndexKeyTooLong() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"failIndexKeyTooLong",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) InternalValue() *DataMongodbatlasClustersResultsAdvancedConfiguration {
	var returns *DataMongodbatlasClustersResultsAdvancedConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) JavascriptEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"javascriptEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) MinimumEnabledTlsProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumEnabledTlsProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) NoTableScan() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"noTableScan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) OplogMinRetentionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogMinRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) OplogSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) SampleRefreshIntervalBiConnector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRefreshIntervalBiConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) SampleSizeBiConnector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeBiConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) TransactionLifetimeLimitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionLifetimeLimitSeconds",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasClustersResultsAdvancedConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasClustersResultsAdvancedConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasClustersResultsAdvancedConfigurationOutputReference_Override(d DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference)SetInternalValue(val *DataMongodbatlasClustersResultsAdvancedConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsAdvancedConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

