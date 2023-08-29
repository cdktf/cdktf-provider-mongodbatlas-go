// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federateddatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v5/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v5/federateddatabaseinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FederatedDatabaseInstanceStorageStoresOutputReference interface {
	cdktf.ComplexObject
	AdditionalStorageClasses() *[]*string
	SetAdditionalStorageClasses(val *[]*string)
	AdditionalStorageClassesInput() *[]*string
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	DefaultFormat() *string
	SetDefaultFormat(val *string)
	DefaultFormatInput() *string
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	// Experimental.
	Fqn() *string
	IncludeTags() interface{}
	SetIncludeTags(val interface{})
	IncludeTagsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	Provider() *string
	SetProvider(val *string)
	ProviderInput() *string
	Public() *string
	SetPublic(val *string)
	PublicInput() *string
	ReadPreference() FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference
	ReadPreferenceInput() *FederatedDatabaseInstanceStorageStoresReadPreference
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutReadPreference(value *FederatedDatabaseInstanceStorageStoresReadPreference)
	ResetAdditionalStorageClasses()
	ResetAllowInsecure()
	ResetBucket()
	ResetClusterId()
	ResetClusterName()
	ResetDefaultFormat()
	ResetDelimiter()
	ResetIncludeTags()
	ResetName()
	ResetPrefix()
	ResetProjectId()
	ResetProvider()
	ResetPublic()
	ResetReadPreference()
	ResetRegion()
	ResetUrls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FederatedDatabaseInstanceStorageStoresOutputReference
type jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) AdditionalStorageClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalStorageClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) AdditionalStorageClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalStorageClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) DefaultFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) DefaultFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) IncludeTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) IncludeTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Provider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Public() *string {
	var returns *string
	_jsii_.Get(
		j,
		"public",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) PublicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ReadPreference() FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference {
	var returns FederatedDatabaseInstanceStorageStoresReadPreferenceOutputReference
	_jsii_.Get(
		j,
		"readPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ReadPreferenceInput() *FederatedDatabaseInstanceStorageStoresReadPreference {
	var returns *FederatedDatabaseInstanceStorageStoresReadPreference
	_jsii_.Get(
		j,
		"readPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Urls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) UrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}


func NewFederatedDatabaseInstanceStorageStoresOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FederatedDatabaseInstanceStorageStoresOutputReference {
	_init_.Initialize()

	if err := validateNewFederatedDatabaseInstanceStorageStoresOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedDatabaseInstance.FederatedDatabaseInstanceStorageStoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFederatedDatabaseInstanceStorageStoresOutputReference_Override(f FederatedDatabaseInstanceStorageStoresOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedDatabaseInstance.FederatedDatabaseInstanceStorageStoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetAdditionalStorageClasses(val *[]*string) {
	if err := j.validateSetAdditionalStorageClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalStorageClasses",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetDefaultFormat(val *string) {
	if err := j.validateSetDefaultFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultFormat",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetIncludeTags(val interface{}) {
	if err := j.validateSetIncludeTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTags",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetProvider(val *string) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetPublic(val *string) {
	if err := j.validateSetPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"public",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference)SetUrls(val *[]*string) {
	if err := j.validateSetUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urls",
		val,
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) PutReadPreference(value *FederatedDatabaseInstanceStorageStoresReadPreference) {
	if err := f.validatePutReadPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putReadPreference",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetAdditionalStorageClasses() {
	_jsii_.InvokeVoid(
		f,
		"resetAdditionalStorageClasses",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		f,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetBucket() {
	_jsii_.InvokeVoid(
		f,
		"resetBucket",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		f,
		"resetClusterId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		f,
		"resetClusterName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetDefaultFormat() {
	_jsii_.InvokeVoid(
		f,
		"resetDefaultFormat",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		f,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetIncludeTags() {
	_jsii_.InvokeVoid(
		f,
		"resetIncludeTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		f,
		"resetName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		f,
		"resetPrefix",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetProjectId() {
	_jsii_.InvokeVoid(
		f,
		"resetProjectId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetProvider() {
	_jsii_.InvokeVoid(
		f,
		"resetProvider",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetPublic() {
	_jsii_.InvokeVoid(
		f,
		"resetPublic",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetReadPreference() {
	_jsii_.InvokeVoid(
		f,
		"resetReadPreference",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		f,
		"resetRegion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ResetUrls() {
	_jsii_.InvokeVoid(
		f,
		"resetUrls",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageStoresOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

