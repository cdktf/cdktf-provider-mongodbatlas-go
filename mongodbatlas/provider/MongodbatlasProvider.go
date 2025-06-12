// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs mongodbatlas}.
type MongodbatlasProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AssumeRole() interface{}
	SetAssumeRole(val interface{})
	AssumeRoleInput() interface{}
	AwsAccessKeyId() *string
	SetAwsAccessKeyId(val *string)
	AwsAccessKeyIdInput() *string
	AwsSecretAccessKey() *string
	SetAwsSecretAccessKey(val *string)
	AwsSecretAccessKeyInput() *string
	AwsSessionToken() *string
	SetAwsSessionToken(val *string)
	AwsSessionTokenInput() *string
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IsMongodbgovCloud() interface{}
	SetIsMongodbgovCloud(val interface{})
	IsMongodbgovCloudInput() interface{}
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	PublicKey() *string
	SetPublicKey(val *string)
	PublicKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	RealmBaseUrl() *string
	SetRealmBaseUrl(val *string)
	RealmBaseUrlInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecretName() *string
	SetSecretName(val *string)
	SecretNameInput() *string
	StsEndpoint() *string
	SetStsEndpoint(val *string)
	StsEndpointInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAssumeRole()
	ResetAwsAccessKeyId()
	ResetAwsSecretAccessKey()
	ResetAwsSessionToken()
	ResetBaseUrl()
	ResetIsMongodbgovCloud()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetPublicKey()
	ResetRealmBaseUrl()
	ResetRegion()
	ResetSecretName()
	ResetStsEndpoint()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MongodbatlasProvider
type jsiiProxy_MongodbatlasProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_MongodbatlasProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AssumeRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AssumeRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AwsAccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AwsAccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AwsSecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AwsSecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AwsSessionToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSessionToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) AwsSessionTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSessionTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) IsMongodbgovCloud() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMongodbgovCloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) IsMongodbgovCloudInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMongodbgovCloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) PublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) RealmBaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmBaseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) RealmBaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmBaseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) SecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) StsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) StsEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongodbatlasProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs mongodbatlas} Resource.
func NewMongodbatlasProvider(scope constructs.Construct, id *string, config *MongodbatlasProviderConfig) MongodbatlasProvider {
	_init_.Initialize()

	if err := validateNewMongodbatlasProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MongodbatlasProvider{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.36.0/docs mongodbatlas} Resource.
func NewMongodbatlasProvider_Override(m MongodbatlasProvider, scope constructs.Construct, id *string, config *MongodbatlasProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetAssumeRole(val interface{}) {
	if err := j.validateSetAssumeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetAwsAccessKeyId(val *string) {
	_jsii_.Set(
		j,
		"awsAccessKeyId",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetAwsSecretAccessKey(val *string) {
	_jsii_.Set(
		j,
		"awsSecretAccessKey",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetAwsSessionToken(val *string) {
	_jsii_.Set(
		j,
		"awsSessionToken",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetIsMongodbgovCloud(val interface{}) {
	if err := j.validateSetIsMongodbgovCloudParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMongodbgovCloud",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetPublicKey(val *string) {
	_jsii_.Set(
		j,
		"publicKey",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetRealmBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"realmBaseUrl",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetSecretName(val *string) {
	_jsii_.Set(
		j,
		"secretName",
		val,
	)
}

func (j *jsiiProxy_MongodbatlasProvider)SetStsEndpoint(val *string) {
	_jsii_.Set(
		j,
		"stsEndpoint",
		val,
	)
}

// Generates CDKTF code for importing a MongodbatlasProvider resource upon running "cdktf plan <stack-name>".
func MongodbatlasProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMongodbatlasProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MongodbatlasProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongodbatlasProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MongodbatlasProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongodbatlasProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MongodbatlasProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongodbatlasProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MongodbatlasProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MongodbatlasProvider) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MongodbatlasProvider) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		m,
		"resetAlias",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		m,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetAwsAccessKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsAccessKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetAwsSecretAccessKey() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsSecretAccessKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetAwsSessionToken() {
	_jsii_.InvokeVoid(
		m,
		"resetAwsSessionToken",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetIsMongodbgovCloud() {
	_jsii_.InvokeVoid(
		m,
		"resetIsMongodbgovCloud",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		m,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetPublicKey() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetRealmBaseUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetRealmBaseUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetSecretName() {
	_jsii_.InvokeVoid(
		m,
		"resetSecretName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) ResetStsEndpoint() {
	_jsii_.InvokeVoid(
		m,
		"resetStsEndpoint",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MongodbatlasProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbatlasProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbatlasProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbatlasProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbatlasProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongodbatlasProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

