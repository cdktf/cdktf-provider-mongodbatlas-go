// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlascloudbackupsnapshotrestorejobs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/datamongodbatlascloudbackupsnapshotrestorejobs/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference interface {
	cdktf.ComplexObject
	Cancelled() cdktf.IResolvable
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
	DeliveryType() *string
	DeliveryUrl() *[]*string
	Expired() cdktf.IResolvable
	ExpiresAt() *string
	Failed() cdktf.IResolvable
	FinishedAt() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataMongodbatlasCloudBackupSnapshotRestoreJobsResults
	SetInternalValue(val *DataMongodbatlasCloudBackupSnapshotRestoreJobsResults)
	OplogInc() *float64
	OplogTs() *float64
	PointInTimeUtcSeconds() *float64
	SnapshotId() *string
	TargetClusterName() *string
	TargetProjectId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timestamp() *string
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

// The jsii proxy struct for DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference
type jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Cancelled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cancelled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) DeliveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) DeliveryUrl() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deliveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Expired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"expired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) ExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Failed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"failed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) FinishedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finishedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) InternalValue() *DataMongodbatlasCloudBackupSnapshotRestoreJobsResults {
	var returns *DataMongodbatlasCloudBackupSnapshotRestoreJobsResults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) OplogInc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogInc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) OplogTs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oplogTs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) PointInTimeUtcSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pointInTimeUtcSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) TargetClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) TargetProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Timestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestamp",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotRestoreJobs.DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference_Override(d DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasCloudBackupSnapshotRestoreJobs.DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference)SetInternalValue(val *DataMongodbatlasCloudBackupSnapshotRestoreJobsResults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataMongodbatlasCloudBackupSnapshotRestoreJobsResultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

