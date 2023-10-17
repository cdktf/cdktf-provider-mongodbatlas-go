// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasclusters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v6/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v6/datamongodbatlasclusters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList
type jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList_Override(d DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasClusters.DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) Get(index *float64) DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataMongodbatlasClustersResultsSnapshotBackupPolicyPoliciesPolicyItemList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

