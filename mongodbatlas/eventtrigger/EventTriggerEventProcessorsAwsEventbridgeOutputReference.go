package eventtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v3/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v3/eventtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventTriggerEventProcessorsAwsEventbridgeOutputReference interface {
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
	ConfigAccountId() *string
	SetConfigAccountId(val *string)
	ConfigAccountIdInput() *string
	ConfigRegion() *string
	SetConfigRegion(val *string)
	ConfigRegionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventTriggerEventProcessorsAwsEventbridge
	SetInternalValue(val *EventTriggerEventProcessorsAwsEventbridge)
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
	ResetConfigAccountId()
	ResetConfigRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventTriggerEventProcessorsAwsEventbridgeOutputReference
type jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ConfigAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ConfigAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ConfigRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ConfigRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) InternalValue() *EventTriggerEventProcessorsAwsEventbridge {
	var returns *EventTriggerEventProcessorsAwsEventbridge
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventTriggerEventProcessorsAwsEventbridgeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventTriggerEventProcessorsAwsEventbridgeOutputReference {
	_init_.Initialize()

	if err := validateNewEventTriggerEventProcessorsAwsEventbridgeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerEventProcessorsAwsEventbridgeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventTriggerEventProcessorsAwsEventbridgeOutputReference_Override(e EventTriggerEventProcessorsAwsEventbridgeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerEventProcessorsAwsEventbridgeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetConfigAccountId(val *string) {
	if err := j.validateSetConfigAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configAccountId",
		val,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetConfigRegion(val *string) {
	if err := j.validateSetConfigRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configRegion",
		val,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetInternalValue(val *EventTriggerEventProcessorsAwsEventbridge) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ResetConfigAccountId() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigAccountId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ResetConfigRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

