// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/federateddatabaseinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference interface {
	cdktf.ComplexObject
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
	Collection() *string
	SetCollection(val *string)
	CollectionInput() *string
	CollectionRegex() *string
	SetCollectionRegex(val *string)
	CollectionRegexInput() *string
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DatabaseRegex() *string
	SetDatabaseRegex(val *string)
	DatabaseRegexInput() *string
	DatasetName() *string
	SetDatasetName(val *string)
	DatasetNameInput() *string
	DefaultFormat() *string
	SetDefaultFormat(val *string)
	DefaultFormatInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Path() *string
	SetPath(val *string)
	PathInput() *string
	ProvenanceFieldName() *string
	SetProvenanceFieldName(val *string)
	ProvenanceFieldNameInput() *string
	StoreName() *string
	SetStoreName(val *string)
	StoreNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Urls() *[]*string
	SetUrls(val *[]*string)
	UrlsInput() *[]*string
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
	ResetAllowInsecure()
	ResetCollection()
	ResetCollectionRegex()
	ResetDatabase()
	ResetDatabaseRegex()
	ResetDatasetName()
	ResetDefaultFormat()
	ResetPath()
	ResetProvenanceFieldName()
	ResetStoreName()
	ResetUrls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference
type jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) Collection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) CollectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) CollectionRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) CollectionRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DatabaseRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DatabaseRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DatasetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DefaultFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) DefaultFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ProvenanceFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provenanceFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ProvenanceFieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provenanceFieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) Urls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) UrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}


func NewFederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewFederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedDatabaseInstance.FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference_Override(f FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedDatabaseInstance.FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetCollection(val *string) {
	if err := j.validateSetCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collection",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetCollectionRegex(val *string) {
	if err := j.validateSetCollectionRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionRegex",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetDatabaseRegex(val *string) {
	if err := j.validateSetDatabaseRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseRegex",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetDatasetName(val *string) {
	if err := j.validateSetDatasetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetDefaultFormat(val *string) {
	if err := j.validateSetDefaultFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultFormat",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetProvenanceFieldName(val *string) {
	if err := j.validateSetProvenanceFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provenanceFieldName",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetStoreName(val *string) {
	if err := j.validateSetStoreNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storeName",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference)SetUrls(val *[]*string) {
	if err := j.validateSetUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urls",
		val,
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetCollection() {
	_jsii_.InvokeVoid(
		f,
		"resetCollection",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetCollectionRegex() {
	_jsii_.InvokeVoid(
		f,
		"resetCollectionRegex",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		f,
		"resetDatabase",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetDatabaseRegex() {
	_jsii_.InvokeVoid(
		f,
		"resetDatabaseRegex",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetDatasetName() {
	_jsii_.InvokeVoid(
		f,
		"resetDatasetName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetDefaultFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		f,
		"resetPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetProvenanceFieldName() {
	_jsii_.InvokeVoid(
		f,
		"resetProvenanceFieldName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetStoreName() {
	_jsii_.InvokeVoid(
		f,
		"resetStoreName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ResetUrls() {
	_jsii_.InvokeVoid(
		f,
		"resetUrls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesCollectionsDataSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

