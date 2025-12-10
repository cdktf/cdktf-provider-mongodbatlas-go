// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/advancedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdvancedClusterAdvancedConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() *float64
	SetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds(val *float64)
	ChangeStreamOptionsPreAndPostImagesExpireAfterSecondsInput() *float64
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
	CustomOpensslCipherConfigTls12() *[]*string
	SetCustomOpensslCipherConfigTls12(val *[]*string)
	CustomOpensslCipherConfigTls12Input() *[]*string
	DefaultMaxTimeMs() *float64
	SetDefaultMaxTimeMs(val *float64)
	DefaultMaxTimeMsInput() *float64
	DefaultReadConcern() *string
	SetDefaultReadConcern(val *string)
	DefaultReadConcernInput() *string
	DefaultWriteConcern() *string
	SetDefaultWriteConcern(val *string)
	DefaultWriteConcernInput() *string
	FailIndexKeyTooLong() interface{}
	SetFailIndexKeyTooLong(val interface{})
	FailIndexKeyTooLongInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AdvancedClusterAdvancedConfiguration
	SetInternalValue(val *AdvancedClusterAdvancedConfiguration)
	JavascriptEnabled() interface{}
	SetJavascriptEnabled(val interface{})
	JavascriptEnabledInput() interface{}
	MinimumEnabledTlsProtocol() *string
	SetMinimumEnabledTlsProtocol(val *string)
	MinimumEnabledTlsProtocolInput() *string
	NoTableScan() interface{}
	SetNoTableScan(val interface{})
	NoTableScanInput() interface{}
	OplogMinRetentionHours() *float64
	SetOplogMinRetentionHours(val *float64)
	OplogMinRetentionHoursInput() *float64
	OplogSizeMb() *float64
	SetOplogSizeMb(val *float64)
	OplogSizeMbInput() *float64
	SampleRefreshIntervalBiConnector() *float64
	SetSampleRefreshIntervalBiConnector(val *float64)
	SampleRefreshIntervalBiConnectorInput() *float64
	SampleSizeBiConnector() *float64
	SetSampleSizeBiConnector(val *float64)
	SampleSizeBiConnectorInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsCipherConfigMode() *string
	SetTlsCipherConfigMode(val *string)
	TlsCipherConfigModeInput() *string
	TransactionLifetimeLimitSeconds() *float64
	SetTransactionLifetimeLimitSeconds(val *float64)
	TransactionLifetimeLimitSecondsInput() *float64
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
	ResetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds()
	ResetCustomOpensslCipherConfigTls12()
	ResetDefaultMaxTimeMs()
	ResetDefaultReadConcern()
	ResetDefaultWriteConcern()
	ResetFailIndexKeyTooLong()
	ResetJavascriptEnabled()
	ResetMinimumEnabledTlsProtocol()
	ResetNoTableScan()
	ResetOplogMinRetentionHours()
	ResetOplogSizeMb()
	ResetSampleRefreshIntervalBiConnector()
	ResetSampleSizeBiConnector()
	ResetTlsCipherConfigMode()
	ResetTransactionLifetimeLimitSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AdvancedClusterAdvancedConfigurationOutputReference
type jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeStreamOptionsPreAndPostImagesExpireAfterSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ChangeStreamOptionsPreAndPostImagesExpireAfterSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"changeStreamOptionsPreAndPostImagesExpireAfterSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) CustomOpensslCipherConfigTls12() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customOpensslCipherConfigTls12",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) CustomOpensslCipherConfigTls12Input() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customOpensslCipherConfigTls12Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) DefaultMaxTimeMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxTimeMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) DefaultMaxTimeMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxTimeMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) DefaultReadConcern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReadConcern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) DefaultReadConcernInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReadConcernInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) DefaultWriteConcern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWriteConcern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) DefaultWriteConcernInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWriteConcernInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) FailIndexKeyTooLong() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failIndexKeyTooLong",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) FailIndexKeyTooLongInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failIndexKeyTooLongInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) InternalValue() *AdvancedClusterAdvancedConfiguration {
	var returns *AdvancedClusterAdvancedConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) JavascriptEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"javascriptEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) JavascriptEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"javascriptEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) MinimumEnabledTlsProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumEnabledTlsProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) MinimumEnabledTlsProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumEnabledTlsProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) NoTableScan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTableScan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) NoTableScanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTableScanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) OplogMinRetentionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogMinRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) OplogMinRetentionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogMinRetentionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) OplogSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) OplogSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) SampleRefreshIntervalBiConnector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRefreshIntervalBiConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) SampleRefreshIntervalBiConnectorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRefreshIntervalBiConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) SampleSizeBiConnector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeBiConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) SampleSizeBiConnectorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeBiConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) TlsCipherConfigMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCipherConfigMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) TlsCipherConfigModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCipherConfigModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) TransactionLifetimeLimitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionLifetimeLimitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) TransactionLifetimeLimitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionLifetimeLimitSecondsInput",
		&returns,
	)
	return returns
}


func NewAdvancedClusterAdvancedConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AdvancedClusterAdvancedConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAdvancedClusterAdvancedConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAdvancedClusterAdvancedConfigurationOutputReference_Override(a AdvancedClusterAdvancedConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds(val *float64) {
	if err := j.validateSetChangeStreamOptionsPreAndPostImagesExpireAfterSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeStreamOptionsPreAndPostImagesExpireAfterSeconds",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetCustomOpensslCipherConfigTls12(val *[]*string) {
	if err := j.validateSetCustomOpensslCipherConfigTls12Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customOpensslCipherConfigTls12",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetDefaultMaxTimeMs(val *float64) {
	if err := j.validateSetDefaultMaxTimeMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMaxTimeMs",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetDefaultReadConcern(val *string) {
	if err := j.validateSetDefaultReadConcernParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultReadConcern",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetDefaultWriteConcern(val *string) {
	if err := j.validateSetDefaultWriteConcernParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultWriteConcern",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetFailIndexKeyTooLong(val interface{}) {
	if err := j.validateSetFailIndexKeyTooLongParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failIndexKeyTooLong",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetInternalValue(val *AdvancedClusterAdvancedConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetJavascriptEnabled(val interface{}) {
	if err := j.validateSetJavascriptEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javascriptEnabled",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetMinimumEnabledTlsProtocol(val *string) {
	if err := j.validateSetMinimumEnabledTlsProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumEnabledTlsProtocol",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetNoTableScan(val interface{}) {
	if err := j.validateSetNoTableScanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTableScan",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetOplogMinRetentionHours(val *float64) {
	if err := j.validateSetOplogMinRetentionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplogMinRetentionHours",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetOplogSizeMb(val *float64) {
	if err := j.validateSetOplogSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplogSizeMb",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetSampleRefreshIntervalBiConnector(val *float64) {
	if err := j.validateSetSampleRefreshIntervalBiConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRefreshIntervalBiConnector",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetSampleSizeBiConnector(val *float64) {
	if err := j.validateSetSampleSizeBiConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSizeBiConnector",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetTlsCipherConfigMode(val *string) {
	if err := j.validateSetTlsCipherConfigModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCipherConfigMode",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference)SetTransactionLifetimeLimitSeconds(val *float64) {
	if err := j.validateSetTransactionLifetimeLimitSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionLifetimeLimitSeconds",
		val,
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetCustomOpensslCipherConfigTls12() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomOpensslCipherConfigTls12",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetDefaultMaxTimeMs() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultMaxTimeMs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetDefaultReadConcern() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultReadConcern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetDefaultWriteConcern() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultWriteConcern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetFailIndexKeyTooLong() {
	_jsii_.InvokeVoid(
		a,
		"resetFailIndexKeyTooLong",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetJavascriptEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetJavascriptEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetMinimumEnabledTlsProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumEnabledTlsProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetNoTableScan() {
	_jsii_.InvokeVoid(
		a,
		"resetNoTableScan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetOplogMinRetentionHours() {
	_jsii_.InvokeVoid(
		a,
		"resetOplogMinRetentionHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetOplogSizeMb() {
	_jsii_.InvokeVoid(
		a,
		"resetOplogSizeMb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetSampleRefreshIntervalBiConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleRefreshIntervalBiConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetSampleSizeBiConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleSizeBiConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetTlsCipherConfigMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsCipherConfigMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ResetTransactionLifetimeLimitSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTransactionLifetimeLimitSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

