package advancedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/advancedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference interface {
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
	DiskIops() *float64
	SetDiskIops(val *float64)
	DiskIopsInput() *float64
	EbsVolumeType() *string
	SetEbsVolumeType(val *string)
	EbsVolumeTypeInput() *string
	// Experimental.
	Fqn() *string
	InstanceSize() *string
	SetInstanceSize(val *string)
	InstanceSizeInput() *string
	InternalValue() *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs
	SetInternalValue(val *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs)
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
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
	ResetDiskIops()
	ResetEbsVolumeType()
	ResetNodeCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference
type jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) DiskIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) DiskIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) EbsVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) EbsVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ebsVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) InstanceSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) InstanceSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) InternalValue() *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs {
	var returns *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference {
	_init_.Initialize()

	if err := validateNewAdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference_Override(a AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetDiskIops(val *float64) {
	if err := j.validateSetDiskIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskIops",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetEbsVolumeType(val *string) {
	if err := j.validateSetEbsVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsVolumeType",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetInstanceSize(val *string) {
	if err := j.validateSetInstanceSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceSize",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetInternalValue(val *AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ResetDiskIops() {
	_jsii_.InvokeVoid(
		a,
		"resetDiskIops",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ResetEbsVolumeType() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsVolumeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

