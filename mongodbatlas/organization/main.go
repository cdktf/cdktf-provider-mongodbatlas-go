// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organization

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.organization.Organization",
		reflect.TypeOf((*Organization)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiAccessListRequired", GoGetter: "ApiAccessListRequired"},
			_jsii_.MemberProperty{JsiiProperty: "apiAccessListRequiredInput", GoGetter: "ApiAccessListRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "genAiFeaturesEnabled", GoGetter: "GenAiFeaturesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "genAiFeaturesEnabledInput", GoGetter: "GenAiFeaturesEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "multiFactorAuthRequired", GoGetter: "MultiFactorAuthRequired"},
			_jsii_.MemberProperty{JsiiProperty: "multiFactorAuthRequiredInput", GoGetter: "MultiFactorAuthRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orgId", GoGetter: "OrgId"},
			_jsii_.MemberProperty{JsiiProperty: "orgOwnerId", GoGetter: "OrgOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "orgOwnerIdInput", GoGetter: "OrgOwnerIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "publicKey", GoGetter: "PublicKey"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiAccessListRequired", GoMethod: "ResetApiAccessListRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetFederationSettingsId", GoMethod: "ResetFederationSettingsId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenAiFeaturesEnabled", GoMethod: "ResetGenAiFeaturesEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMultiFactorAuthRequired", GoMethod: "ResetMultiFactorAuthRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrgOwnerId", GoMethod: "ResetOrgOwnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictEmployeeAccess", GoMethod: "ResetRestrictEmployeeAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleNames", GoMethod: "ResetRoleNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityContact", GoMethod: "ResetSecurityContact"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipDefaultAlertsSettings", GoMethod: "ResetSkipDefaultAlertsSettings"},
			_jsii_.MemberProperty{JsiiProperty: "restrictEmployeeAccess", GoGetter: "RestrictEmployeeAccess"},
			_jsii_.MemberProperty{JsiiProperty: "restrictEmployeeAccessInput", GoGetter: "RestrictEmployeeAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleNames", GoGetter: "RoleNames"},
			_jsii_.MemberProperty{JsiiProperty: "roleNamesInput", GoGetter: "RoleNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityContact", GoGetter: "SecurityContact"},
			_jsii_.MemberProperty{JsiiProperty: "securityContactInput", GoGetter: "SecurityContactInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipDefaultAlertsSettings", GoGetter: "SkipDefaultAlertsSettings"},
			_jsii_.MemberProperty{JsiiProperty: "skipDefaultAlertsSettingsInput", GoGetter: "SkipDefaultAlertsSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Organization{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.organization.OrganizationConfig",
		reflect.TypeOf((*OrganizationConfig)(nil)).Elem(),
	)
}
