// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federatedsettingsidentityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/federatedsettingsidentityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/resources/federated_settings_identity_provider mongodbatlas_federated_settings_identity_provider}.
type FederatedSettingsIdentityProvider interface {
	cdktf.TerraformResource
	AssociatedDomains() *[]*string
	SetAssociatedDomains(val *[]*string)
	AssociatedDomainsInput() *[]*string
	Audience() *string
	SetAudience(val *string)
	AudienceInput() *string
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	AuthorizationTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	SetDescription(val *string)
	DescriptionInput() *string
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
	SetGroupsClaim(val *string)
	GroupsClaimInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdpId() *string
	IdpType() *string
	SetIdpType(val *string)
	IdpTypeInput() *string
	IssuerUri() *string
	SetIssuerUri(val *string)
	IssuerUriInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OktaIdpId() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RequestBinding() *string
	SetRequestBinding(val *string)
	RequestBindingInput() *string
	RequestedScopes() *[]*string
	SetRequestedScopes(val *[]*string)
	RequestedScopesInput() *[]*string
	ResponseSignatureAlgorithm() *string
	SetResponseSignatureAlgorithm(val *string)
	ResponseSignatureAlgorithmInput() *string
	SsoDebugEnabled() interface{}
	SetSsoDebugEnabled(val interface{})
	SsoDebugEnabledInput() interface{}
	SsoUrl() *string
	SetSsoUrl(val *string)
	SsoUrlInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserClaim() *string
	SetUserClaim(val *string)
	UserClaimInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAssociatedDomains()
	ResetAudience()
	ResetAuthorizationType()
	ResetClientId()
	ResetDescription()
	ResetGroupsClaim()
	ResetId()
	ResetIdpType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocol()
	ResetRequestBinding()
	ResetRequestedScopes()
	ResetResponseSignatureAlgorithm()
	ResetSsoDebugEnabled()
	ResetSsoUrl()
	ResetStatus()
	ResetUserClaim()
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

// The jsii proxy struct for FederatedSettingsIdentityProvider
type jsiiProxy_FederatedSettingsIdentityProvider struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) AssociatedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) AssociatedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"associatedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) AuthorizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) FederationSettingsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationSettingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) FederationSettingsIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationSettingsIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) GroupsClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) GroupsClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) IdpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) IdpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) IdpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) IssuerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) OktaIdpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oktaIdpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) RequestBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) RequestBindingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) RequestedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) RequestedScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ResponseSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) ResponseSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) SsoDebugEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoDebugEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) SsoDebugEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssoDebugEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) SsoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) UserClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider) UserClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userClaimInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/resources/federated_settings_identity_provider mongodbatlas_federated_settings_identity_provider} Resource.
func NewFederatedSettingsIdentityProvider(scope constructs.Construct, id *string, config *FederatedSettingsIdentityProviderConfig) FederatedSettingsIdentityProvider {
	_init_.Initialize()

	if err := validateNewFederatedSettingsIdentityProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedSettingsIdentityProvider{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.1/docs/resources/federated_settings_identity_provider mongodbatlas_federated_settings_identity_provider} Resource.
func NewFederatedSettingsIdentityProvider_Override(f FederatedSettingsIdentityProvider, scope constructs.Construct, id *string, config *FederatedSettingsIdentityProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetAssociatedDomains(val *[]*string) {
	if err := j.validateSetAssociatedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedDomains",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetAudience(val *string) {
	if err := j.validateSetAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetAuthorizationType(val *string) {
	if err := j.validateSetAuthorizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetFederationSettingsId(val *string) {
	if err := j.validateSetFederationSettingsIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federationSettingsId",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetGroupsClaim(val *string) {
	if err := j.validateSetGroupsClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsClaim",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetIdpType(val *string) {
	if err := j.validateSetIdpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpType",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetIssuerUri(val *string) {
	if err := j.validateSetIssuerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUri",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetRequestBinding(val *string) {
	if err := j.validateSetRequestBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBinding",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetRequestedScopes(val *[]*string) {
	if err := j.validateSetRequestedScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedScopes",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetResponseSignatureAlgorithm(val *string) {
	if err := j.validateSetResponseSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetSsoDebugEnabled(val interface{}) {
	if err := j.validateSetSsoDebugEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoDebugEnabled",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetSsoUrl(val *string) {
	if err := j.validateSetSsoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssoUrl",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_FederatedSettingsIdentityProvider)SetUserClaim(val *string) {
	if err := j.validateSetUserClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userClaim",
		val,
	)
}

// Generates CDKTF code for importing a FederatedSettingsIdentityProvider resource upon running "cdktf plan <stack-name>".
func FederatedSettingsIdentityProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFederatedSettingsIdentityProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
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
func FederatedSettingsIdentityProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFederatedSettingsIdentityProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FederatedSettingsIdentityProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFederatedSettingsIdentityProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FederatedSettingsIdentityProvider_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFederatedSettingsIdentityProvider_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FederatedSettingsIdentityProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FederatedSettingsIdentityProvider) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetAssociatedDomains() {
	_jsii_.InvokeVoid(
		f,
		"resetAssociatedDomains",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetAudience() {
	_jsii_.InvokeVoid(
		f,
		"resetAudience",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetAuthorizationType() {
	_jsii_.InvokeVoid(
		f,
		"resetAuthorizationType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		f,
		"resetClientId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetGroupsClaim() {
	_jsii_.InvokeVoid(
		f,
		"resetGroupsClaim",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetIdpType() {
	_jsii_.InvokeVoid(
		f,
		"resetIdpType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetProtocol() {
	_jsii_.InvokeVoid(
		f,
		"resetProtocol",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetRequestBinding() {
	_jsii_.InvokeVoid(
		f,
		"resetRequestBinding",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetRequestedScopes() {
	_jsii_.InvokeVoid(
		f,
		"resetRequestedScopes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetResponseSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		f,
		"resetResponseSignatureAlgorithm",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetSsoDebugEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetSsoDebugEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetSsoUrl() {
	_jsii_.InvokeVoid(
		f,
		"resetSsoUrl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetStatus() {
	_jsii_.InvokeVoid(
		f,
		"resetStatus",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ResetUserClaim() {
	_jsii_.InvokeVoid(
		f,
		"resetUserClaim",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedSettingsIdentityProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

