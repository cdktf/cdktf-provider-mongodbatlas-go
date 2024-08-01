// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasfederatedsettingsidentityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/datamongodbatlasfederatedsettingsidentityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/data-sources/federated_settings_identity_provider mongodbatlas_federated_settings_identity_provider}.
type DataMongodbatlasFederatedSettingsIdentityProvider interface {
	cdktf.TerraformDataSource
	AcsUrl() *string
	AssociatedDomains() *[]*string
	AssociatedOrgs() DataMongodbatlasFederatedSettingsIdentityProviderAssociatedOrgsList
	Audience() *string
	AudienceUri() *string
	AuthorizationType() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DisplayName() *string
	FederationSettingsId() *string
	SetFederationSettingsId(val *string)
	FederationSettingsIdInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsClaim() *string
	Id() *string
	SetId(val *string)
	IdentityProviderId() *string
	SetIdentityProviderId(val *string)
	IdentityProviderIdInput() *string
	IdInput() *string
	IdpId() *string
	IdpType() *string
	IssuerUri() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OktaIdpId() *string
	PemFileInfo() DataMongodbatlasFederatedSettingsIdentityProviderPemFileInfoList
	Protocol() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RequestBinding() *string
	RequestedScopes() *[]*string
	ResponseSignatureAlgorithm() *string
	SsoDebugEnabled() cdktf.IResolvable
	SsoUrl() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserClaim() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataMongodbatlasFederatedSettingsIdentityProvider
type jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) AcsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) AssociatedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) AssociatedOrgs() DataMongodbatlasFederatedSettingsIdentityProviderAssociatedOrgsList {
	var returns DataMongodbatlasFederatedSettingsIdentityProviderAssociatedOrgsList
	_jsii_.Get(
		j,
		"associatedOrgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) AudienceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) FederationSettingsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationSettingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) FederationSettingsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationSettingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GroupsClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) IdentityProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) IdentityProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) IdpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) IdpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) OktaIdpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaIdpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) PemFileInfo() DataMongodbatlasFederatedSettingsIdentityProviderPemFileInfoList {
	var returns DataMongodbatlasFederatedSettingsIdentityProviderPemFileInfoList
	_jsii_.Get(
		j,
		"pemFileInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) RequestBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) RequestedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ResponseSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) SsoDebugEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ssoDebugEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) UserClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaim",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/data-sources/federated_settings_identity_provider mongodbatlas_federated_settings_identity_provider} Data Source.
func NewDataMongodbatlasFederatedSettingsIdentityProvider(scope constructs.Construct, id *string, config *DataMongodbatlasFederatedSettingsIdentityProviderConfig) DataMongodbatlasFederatedSettingsIdentityProvider {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasFederatedSettingsIdentityProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/data-sources/federated_settings_identity_provider mongodbatlas_federated_settings_identity_provider} Data Source.
func NewDataMongodbatlasFederatedSettingsIdentityProvider_Override(d DataMongodbatlasFederatedSettingsIdentityProvider, scope constructs.Construct, id *string, config *DataMongodbatlasFederatedSettingsIdentityProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetFederationSettingsId(val *string) {
	if err := j.validateSetFederationSettingsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federationSettingsId",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetIdentityProviderId(val *string) {
	if err := j.validateSetIdentityProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderId",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataMongodbatlasFederatedSettingsIdentityProvider resource upon running "cdktf plan <stack-name>".
func DataMongodbatlasFederatedSettingsIdentityProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataMongodbatlasFederatedSettingsIdentityProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
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
func DataMongodbatlasFederatedSettingsIdentityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasFederatedSettingsIdentityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasFederatedSettingsIdentityProvider_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasFederatedSettingsIdentityProvider_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataMongodbatlasFederatedSettingsIdentityProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataMongodbatlasFederatedSettingsIdentityProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataMongodbatlasFederatedSettingsIdentityProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasFederatedSettingsIdentityProvider.DataMongodbatlasFederatedSettingsIdentityProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasFederatedSettingsIdentityProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

