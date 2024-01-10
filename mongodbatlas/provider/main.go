// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProvider",
		reflect.TypeOf((*MongodbatlasProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRole", GoGetter: "AssumeRole"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleInput", GoGetter: "AssumeRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyId", GoGetter: "AwsAccessKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessKeyIdInput", GoGetter: "AwsAccessKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKey", GoGetter: "AwsSecretAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretAccessKeyInput", GoGetter: "AwsSecretAccessKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSessionToken", GoGetter: "AwsSessionToken"},
			_jsii_.MemberProperty{JsiiProperty: "awsSessionTokenInput", GoGetter: "AwsSessionTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrlInput", GoGetter: "BaseUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "isMongodbgovCloud", GoGetter: "IsMongodbgovCloud"},
			_jsii_.MemberProperty{JsiiProperty: "isMongodbgovCloudInput", GoGetter: "IsMongodbgovCloudInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicKey", GoGetter: "PublicKey"},
			_jsii_.MemberProperty{JsiiProperty: "publicKeyInput", GoGetter: "PublicKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "realmBaseUrl", GoGetter: "RealmBaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "realmBaseUrlInput", GoGetter: "RealmBaseUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssumeRole", GoMethod: "ResetAssumeRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccessKeyId", GoMethod: "ResetAwsAccessKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSecretAccessKey", GoMethod: "ResetAwsSecretAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSessionToken", GoMethod: "ResetAwsSessionToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseUrl", GoMethod: "ResetBaseUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsMongodbgovCloud", GoMethod: "ResetIsMongodbgovCloud"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKey", GoMethod: "ResetPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicKey", GoMethod: "ResetPublicKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealmBaseUrl", GoMethod: "ResetRealmBaseUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretName", GoMethod: "ResetSecretName"},
			_jsii_.MemberMethod{JsiiMethod: "resetStsEndpoint", GoMethod: "ResetStsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "stsEndpoint", GoGetter: "StsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "stsEndpointInput", GoGetter: "StsEndpointInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_MongodbatlasProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProviderAssumeRole",
		reflect.TypeOf((*MongodbatlasProviderAssumeRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.provider.MongodbatlasProviderConfig",
		reflect.TypeOf((*MongodbatlasProviderConfig)(nil)).Elem(),
	)
}
