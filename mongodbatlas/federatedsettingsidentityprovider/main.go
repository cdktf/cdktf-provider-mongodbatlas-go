// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federatedsettingsidentityprovider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProvider",
		reflect.TypeOf((*FederatedSettingsIdentityProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "associatedDomains", GoGetter: "AssociatedDomains"},
			_jsii_.MemberProperty{JsiiProperty: "associatedDomainsInput", GoGetter: "AssociatedDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "audience", GoGetter: "Audience"},
			_jsii_.MemberProperty{JsiiProperty: "audienceInput", GoGetter: "AudienceInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationType", GoGetter: "AuthorizationType"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationTypeInput", GoGetter: "AuthorizationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "federationSettingsId", GoGetter: "FederationSettingsId"},
			_jsii_.MemberProperty{JsiiProperty: "federationSettingsIdInput", GoGetter: "FederationSettingsIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupsClaim", GoGetter: "GroupsClaim"},
			_jsii_.MemberProperty{JsiiProperty: "groupsClaimInput", GoGetter: "GroupsClaimInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idpId", GoGetter: "IdpId"},
			_jsii_.MemberProperty{JsiiProperty: "idpType", GoGetter: "IdpType"},
			_jsii_.MemberProperty{JsiiProperty: "idpTypeInput", GoGetter: "IdpTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuerUri", GoGetter: "IssuerUri"},
			_jsii_.MemberProperty{JsiiProperty: "issuerUriInput", GoGetter: "IssuerUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oktaIdpId", GoGetter: "OktaIdpId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestBinding", GoGetter: "RequestBinding"},
			_jsii_.MemberProperty{JsiiProperty: "requestBindingInput", GoGetter: "RequestBindingInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestedScopes", GoGetter: "RequestedScopes"},
			_jsii_.MemberProperty{JsiiProperty: "requestedScopesInput", GoGetter: "RequestedScopesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssociatedDomains", GoMethod: "ResetAssociatedDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudience", GoMethod: "ResetAudience"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationType", GoMethod: "ResetAuthorizationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupsClaim", GoMethod: "ResetGroupsClaim"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdpType", GoMethod: "ResetIdpType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocol", GoMethod: "ResetProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestBinding", GoMethod: "ResetRequestBinding"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestedScopes", GoMethod: "ResetRequestedScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseSignatureAlgorithm", GoMethod: "ResetResponseSignatureAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsoDebugEnabled", GoMethod: "ResetSsoDebugEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsoUrl", GoMethod: "ResetSsoUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserClaim", GoMethod: "ResetUserClaim"},
			_jsii_.MemberProperty{JsiiProperty: "responseSignatureAlgorithm", GoGetter: "ResponseSignatureAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "responseSignatureAlgorithmInput", GoGetter: "ResponseSignatureAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssoDebugEnabled", GoGetter: "SsoDebugEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "ssoDebugEnabledInput", GoGetter: "SsoDebugEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssoUrl", GoGetter: "SsoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "ssoUrlInput", GoGetter: "SsoUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userClaim", GoGetter: "UserClaim"},
			_jsii_.MemberProperty{JsiiProperty: "userClaimInput", GoGetter: "UserClaimInput"},
		},
		func() interface{} {
			j := jsiiProxy_FederatedSettingsIdentityProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.federatedSettingsIdentityProvider.FederatedSettingsIdentityProviderConfig",
		reflect.TypeOf((*FederatedSettingsIdentityProviderConfig)(nil)).Elem(),
	)
}
