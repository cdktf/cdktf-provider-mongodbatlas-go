package thirdpartyintegration

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.thirdPartyIntegration.ThirdPartyIntegration",
		reflect.TypeOf((*ThirdPartyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyInput", GoGetter: "ApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiToken", GoGetter: "ApiToken"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenInput", GoGetter: "ApiTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "channelName", GoGetter: "ChannelName"},
			_jsii_.MemberProperty{JsiiProperty: "channelNameInput", GoGetter: "ChannelNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "flowName", GoGetter: "FlowName"},
			_jsii_.MemberProperty{JsiiProperty: "flowNameInput", GoGetter: "FlowNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "licenseKey", GoGetter: "LicenseKey"},
			_jsii_.MemberProperty{JsiiProperty: "licenseKeyInput", GoGetter: "LicenseKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftTeamsWebhookUrl", GoGetter: "MicrosoftTeamsWebhookUrl"},
			_jsii_.MemberProperty{JsiiProperty: "microsoftTeamsWebhookUrlInput", GoGetter: "MicrosoftTeamsWebhookUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orgName", GoGetter: "OrgName"},
			_jsii_.MemberProperty{JsiiProperty: "orgNameInput", GoGetter: "OrgNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readToken", GoGetter: "ReadToken"},
			_jsii_.MemberProperty{JsiiProperty: "readTokenInput", GoGetter: "ReadTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiKey", GoMethod: "ResetApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiToken", GoMethod: "ResetApiToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannelName", GoMethod: "ResetChannelName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFlowName", GoMethod: "ResetFlowName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseKey", GoMethod: "ResetLicenseKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetMicrosoftTeamsWebhookUrl", GoMethod: "ResetMicrosoftTeamsWebhookUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrgName", GoMethod: "ResetOrgName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadToken", GoMethod: "ResetReadToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoutingKey", GoMethod: "ResetRoutingKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheme", GoMethod: "ResetScheme"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecret", GoMethod: "ResetSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceDiscovery", GoMethod: "ResetServiceDiscovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceKey", GoMethod: "ResetServiceKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeamName", GoMethod: "ResetTeamName"},
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserName", GoMethod: "ResetUserName"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteToken", GoMethod: "ResetWriteToken"},
			_jsii_.MemberProperty{JsiiProperty: "routingKey", GoGetter: "RoutingKey"},
			_jsii_.MemberProperty{JsiiProperty: "routingKeyInput", GoGetter: "RoutingKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "schemeInput", GoGetter: "SchemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "secretInput", GoGetter: "SecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDiscovery", GoGetter: "ServiceDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDiscoveryInput", GoGetter: "ServiceDiscoveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceKey", GoGetter: "ServiceKey"},
			_jsii_.MemberProperty{JsiiProperty: "serviceKeyInput", GoGetter: "ServiceKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "teamName", GoGetter: "TeamName"},
			_jsii_.MemberProperty{JsiiProperty: "teamNameInput", GoGetter: "TeamNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userNameInput", GoGetter: "UserNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "writeToken", GoGetter: "WriteToken"},
			_jsii_.MemberProperty{JsiiProperty: "writeTokenInput", GoGetter: "WriteTokenInput"},
		},
		func() interface{} {
			j := jsiiProxy_ThirdPartyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.thirdPartyIntegration.ThirdPartyIntegrationConfig",
		reflect.TypeOf((*ThirdPartyIntegrationConfig)(nil)).Elem(),
	)
}