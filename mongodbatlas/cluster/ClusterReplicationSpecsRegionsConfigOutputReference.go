package cluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/cluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterReplicationSpecsRegionsConfigOutputReference interface {
	cdktf.ComplexObject
	AnalyticsNodes() *float64
	SetAnalyticsNodes(val *float64)
	AnalyticsNodesInput() *float64
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
	ElectableNodes() *float64
	SetElectableNodes(val *float64)
	ElectableNodesInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	ReadOnlyNodes() *float64
	SetReadOnlyNodes(val *float64)
	ReadOnlyNodesInput() *float64
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
	ResetAnalyticsNodes()
	ResetElectableNodes()
	ResetPriority()
	ResetReadOnlyNodes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClusterReplicationSpecsRegionsConfigOutputReference
type jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) AnalyticsNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analyticsNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) AnalyticsNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analyticsNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ElectableNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"electableNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ElectableNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"electableNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ReadOnlyNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readOnlyNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ReadOnlyNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readOnlyNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) RegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) RegionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClusterReplicationSpecsRegionsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClusterReplicationSpecsRegionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewClusterReplicationSpecsRegionsConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cluster.ClusterReplicationSpecsRegionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClusterReplicationSpecsRegionsConfigOutputReference_Override(c ClusterReplicationSpecsRegionsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cluster.ClusterReplicationSpecsRegionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetAnalyticsNodes(val *float64) {
	if err := j.validateSetAnalyticsNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticsNodes",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetElectableNodes(val *float64) {
	if err := j.validateSetElectableNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"electableNodes",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetReadOnlyNodes(val *float64) {
	if err := j.validateSetReadOnlyNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyNodes",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetRegionName(val *string) {
	if err := j.validateSetRegionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionName",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ResetAnalyticsNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalyticsNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ResetElectableNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetElectableNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		c,
		"resetPriority",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ResetReadOnlyNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetReadOnlyNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
