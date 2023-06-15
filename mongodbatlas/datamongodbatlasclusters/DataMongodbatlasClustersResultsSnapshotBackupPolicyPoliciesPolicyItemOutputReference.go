package datamongodbatlasclusters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/datamongodbatlasclusters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference interface {
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
	FrequencyInterval() *float64
	FrequencyType() *string
	Id() *string
	InternalValue() *DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItem
	SetInternalValue(val *DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItem)
	RetentionUnit() *string
	RetentionValue() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference
type jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) FrequencyInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) FrequencyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) InternalValue() *DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItem {
	var returns *DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItem
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) RetentionUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) RetentionValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference_Override(d DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetInternalValue(val *DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItem) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

