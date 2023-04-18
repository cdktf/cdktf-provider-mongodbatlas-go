package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterAdvancedConfigurationOutputReference interface {
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
	InternalValue() *ClusterAdvancedConfiguration
	SetInternalValue(val *ClusterAdvancedConfiguration)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClusterAdvancedConfigurationOutputReference
type jsiiProxy_ClusterAdvancedConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) DefaultReadConcern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReadConcern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) DefaultReadConcernInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultReadConcernInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) DefaultWriteConcern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWriteConcern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) DefaultWriteConcernInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultWriteConcernInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) FailIndexKeyTooLong() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failIndexKeyTooLong",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) FailIndexKeyTooLongInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failIndexKeyTooLongInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) InternalValue() *ClusterAdvancedConfiguration {
	var returns *ClusterAdvancedConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) JavascriptEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"javascriptEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) JavascriptEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"javascriptEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) MinimumEnabledTlsProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumEnabledTlsProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) MinimumEnabledTlsProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumEnabledTlsProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) NoTableScan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTableScan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) NoTableScanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noTableScanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) OplogMinRetentionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogMinRetentionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) OplogMinRetentionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogMinRetentionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) OplogSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) OplogSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) SampleRefreshIntervalBiConnector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRefreshIntervalBiConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) SampleRefreshIntervalBiConnectorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRefreshIntervalBiConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) SampleSizeBiConnector() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeBiConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) SampleSizeBiConnectorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleSizeBiConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClusterAdvancedConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClusterAdvancedConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewClusterAdvancedConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterAdvancedConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cluster.ClusterAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClusterAdvancedConfigurationOutputReference_Override(c ClusterAdvancedConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cluster.ClusterAdvancedConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetDefaultReadConcern(val *string) {
	if err := j.validateSetDefaultReadConcernParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultReadConcern",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetDefaultWriteConcern(val *string) {
	if err := j.validateSetDefaultWriteConcernParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultWriteConcern",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetFailIndexKeyTooLong(val interface{}) {
	if err := j.validateSetFailIndexKeyTooLongParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failIndexKeyTooLong",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetInternalValue(val *ClusterAdvancedConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetJavascriptEnabled(val interface{}) {
	if err := j.validateSetJavascriptEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javascriptEnabled",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetMinimumEnabledTlsProtocol(val *string) {
	if err := j.validateSetMinimumEnabledTlsProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumEnabledTlsProtocol",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetNoTableScan(val interface{}) {
	if err := j.validateSetNoTableScanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noTableScan",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetOplogMinRetentionHours(val *float64) {
	if err := j.validateSetOplogMinRetentionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplogMinRetentionHours",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetOplogSizeMb(val *float64) {
	if err := j.validateSetOplogSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplogSizeMb",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetSampleRefreshIntervalBiConnector(val *float64) {
	if err := j.validateSetSampleRefreshIntervalBiConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRefreshIntervalBiConnector",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetSampleSizeBiConnector(val *float64) {
	if err := j.validateSetSampleSizeBiConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleSizeBiConnector",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterAdvancedConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetDefaultReadConcern() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultReadConcern",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetDefaultWriteConcern() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultWriteConcern",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetFailIndexKeyTooLong() {
	_jsii_.InvokeVoid(
		c,
		"resetFailIndexKeyTooLong",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetJavascriptEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetJavascriptEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetMinimumEnabledTlsProtocol() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumEnabledTlsProtocol",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetNoTableScan() {
	_jsii_.InvokeVoid(
		c,
		"resetNoTableScan",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetOplogMinRetentionHours() {
	_jsii_.InvokeVoid(
		c,
		"resetOplogMinRetentionHours",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetOplogSizeMb() {
	_jsii_.InvokeVoid(
		c,
		"resetOplogSizeMb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetSampleRefreshIntervalBiConnector() {
	_jsii_.InvokeVoid(
		c,
		"resetSampleRefreshIntervalBiConnector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ResetSampleSizeBiConnector() {
	_jsii_.InvokeVoid(
		c,
		"resetSampleSizeBiConnector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClusterAdvancedConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

