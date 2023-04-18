package advancedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/advancedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdvancedClusterReplicationSpecsRegionConfigsOutputReference interface {
	cdktf.ComplexObject
	AnalyticsAutoScaling() AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScalingOutputReference
	AnalyticsAutoScalingInput() *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling
	AnalyticsSpecs() AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecsOutputReference
	AnalyticsSpecsInput() *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs
	AutoScaling() AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference
	AutoScalingInput() *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling
	BackingProviderName() *string
	SetBackingProviderName(val *string)
	BackingProviderNameInput() *string
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
	ElectableSpecs() AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference
	ElectableSpecsInput() *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	ReadOnlySpecs() AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecsOutputReference
	ReadOnlySpecsInput() *AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs
	RegionName() *string
	SetRegionName(val *string)
	RegionNameInput() *string
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
	PutAnalyticsAutoScaling(value *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling)
	PutAnalyticsSpecs(value *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs)
	PutAutoScaling(value *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling)
	PutElectableSpecs(value *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs)
	PutReadOnlySpecs(value *AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs)
	ResetAnalyticsAutoScaling()
	ResetAnalyticsSpecs()
	ResetAutoScaling()
	ResetBackingProviderName()
	ResetElectableSpecs()
	ResetReadOnlySpecs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AdvancedClusterReplicationSpecsRegionConfigsOutputReference
type jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) AnalyticsAutoScaling() AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScalingOutputReference {
	var returns AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScalingOutputReference
	_jsii_.Get(
		j,
		"analyticsAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) AnalyticsAutoScalingInput() *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling
	_jsii_.Get(
		j,
		"analyticsAutoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) AnalyticsSpecs() AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecsOutputReference {
	var returns AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecsOutputReference
	_jsii_.Get(
		j,
		"analyticsSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) AnalyticsSpecsInput() *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs
	_jsii_.Get(
		j,
		"analyticsSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) AutoScaling() AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference {
	var returns AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference
	_jsii_.Get(
		j,
		"autoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) AutoScalingInput() *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling
	_jsii_.Get(
		j,
		"autoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) BackingProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backingProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) BackingProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backingProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ElectableSpecs() AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference {
	var returns AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference
	_jsii_.Get(
		j,
		"electableSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ElectableSpecsInput() *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs
	_jsii_.Get(
		j,
		"electableSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ReadOnlySpecs() AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecsOutputReference {
	var returns AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecsOutputReference
	_jsii_.Get(
		j,
		"readOnlySpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ReadOnlySpecsInput() *AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs
	_jsii_.Get(
		j,
		"readOnlySpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) RegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) RegionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAdvancedClusterReplicationSpecsRegionConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AdvancedClusterReplicationSpecsRegionConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewAdvancedClusterReplicationSpecsRegionConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAdvancedClusterReplicationSpecsRegionConfigsOutputReference_Override(a AdvancedClusterReplicationSpecsRegionConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetBackingProviderName(val *string) {
	if err := j.validateSetBackingProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backingProviderName",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetRegionName(val *string) {
	if err := j.validateSetRegionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionName",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) PutAnalyticsAutoScaling(value *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling) {
	if err := a.validatePutAnalyticsAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnalyticsAutoScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) PutAnalyticsSpecs(value *AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs) {
	if err := a.validatePutAnalyticsSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnalyticsSpecs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) PutAutoScaling(value *AdvancedClusterReplicationSpecsRegionConfigsAutoScaling) {
	if err := a.validatePutAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) PutElectableSpecs(value *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs) {
	if err := a.validatePutElectableSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElectableSpecs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) PutReadOnlySpecs(value *AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs) {
	if err := a.validatePutReadOnlySpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReadOnlySpecs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ResetAnalyticsAutoScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalyticsAutoScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ResetAnalyticsSpecs() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalyticsSpecs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ResetAutoScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ResetBackingProviderName() {
	_jsii_.InvokeVoid(
		a,
		"resetBackingProviderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ResetElectableSpecs() {
	_jsii_.InvokeVoid(
		a,
		"resetElectableSpecs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ResetReadOnlySpecs() {
	_jsii_.InvokeVoid(
		a,
		"resetReadOnlySpecs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

