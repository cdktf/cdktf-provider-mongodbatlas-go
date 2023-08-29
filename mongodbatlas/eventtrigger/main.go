// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventtrigger

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTrigger",
		reflect.TypeOf((*EventTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appId", GoGetter: "AppId"},
			_jsii_.MemberProperty{JsiiProperty: "appIdInput", GoGetter: "AppIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "configCollection", GoGetter: "ConfigCollection"},
			_jsii_.MemberProperty{JsiiProperty: "configCollectionInput", GoGetter: "ConfigCollectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "configDatabase", GoGetter: "ConfigDatabase"},
			_jsii_.MemberProperty{JsiiProperty: "configDatabaseInput", GoGetter: "ConfigDatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "configFullDocument", GoGetter: "ConfigFullDocument"},
			_jsii_.MemberProperty{JsiiProperty: "configFullDocumentBefore", GoGetter: "ConfigFullDocumentBefore"},
			_jsii_.MemberProperty{JsiiProperty: "configFullDocumentBeforeInput", GoGetter: "ConfigFullDocumentBeforeInput"},
			_jsii_.MemberProperty{JsiiProperty: "configFullDocumentInput", GoGetter: "ConfigFullDocumentInput"},
			_jsii_.MemberProperty{JsiiProperty: "configMatch", GoGetter: "ConfigMatch"},
			_jsii_.MemberProperty{JsiiProperty: "configMatchInput", GoGetter: "ConfigMatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "configOperationType", GoGetter: "ConfigOperationType"},
			_jsii_.MemberProperty{JsiiProperty: "configOperationTypeInput", GoGetter: "ConfigOperationTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "configOperationTypes", GoGetter: "ConfigOperationTypes"},
			_jsii_.MemberProperty{JsiiProperty: "configOperationTypesInput", GoGetter: "ConfigOperationTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "configProject", GoGetter: "ConfigProject"},
			_jsii_.MemberProperty{JsiiProperty: "configProjectInput", GoGetter: "ConfigProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "configProviders", GoGetter: "ConfigProviders"},
			_jsii_.MemberProperty{JsiiProperty: "configProvidersInput", GoGetter: "ConfigProvidersInput"},
			_jsii_.MemberProperty{JsiiProperty: "configSchedule", GoGetter: "ConfigSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "configScheduleInput", GoGetter: "ConfigScheduleInput"},
			_jsii_.MemberProperty{JsiiProperty: "configScheduleType", GoGetter: "ConfigScheduleType"},
			_jsii_.MemberProperty{JsiiProperty: "configServiceId", GoGetter: "ConfigServiceId"},
			_jsii_.MemberProperty{JsiiProperty: "configServiceIdInput", GoGetter: "ConfigServiceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disabled", GoGetter: "Disabled"},
			_jsii_.MemberProperty{JsiiProperty: "disabledInput", GoGetter: "DisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventProcessors", GoGetter: "EventProcessors"},
			_jsii_.MemberProperty{JsiiProperty: "eventProcessorsInput", GoGetter: "EventProcessorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "functionId", GoGetter: "FunctionId"},
			_jsii_.MemberProperty{JsiiProperty: "functionIdInput", GoGetter: "FunctionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "functionName", GoGetter: "FunctionName"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putEventProcessors", GoMethod: "PutEventProcessors"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigCollection", GoMethod: "ResetConfigCollection"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigDatabase", GoMethod: "ResetConfigDatabase"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigFullDocument", GoMethod: "ResetConfigFullDocument"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigFullDocumentBefore", GoMethod: "ResetConfigFullDocumentBefore"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigMatch", GoMethod: "ResetConfigMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigOperationType", GoMethod: "ResetConfigOperationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigOperationTypes", GoMethod: "ResetConfigOperationTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigProject", GoMethod: "ResetConfigProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigProviders", GoMethod: "ResetConfigProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigSchedule", GoMethod: "ResetConfigSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigServiceId", GoMethod: "ResetConfigServiceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisabled", GoMethod: "ResetDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventProcessors", GoMethod: "ResetEventProcessors"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionId", GoMethod: "ResetFunctionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnordered", GoMethod: "ResetUnordered"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "triggerId", GoGetter: "TriggerId"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "unordered", GoGetter: "Unordered"},
			_jsii_.MemberProperty{JsiiProperty: "unorderedInput", GoGetter: "UnorderedInput"},
		},
		func() interface{} {
			j := jsiiProxy_EventTrigger{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerConfig",
		reflect.TypeOf((*EventTriggerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerEventProcessors",
		reflect.TypeOf((*EventTriggerEventProcessors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerEventProcessorsAwsEventbridge",
		reflect.TypeOf((*EventTriggerEventProcessorsAwsEventbridge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerEventProcessorsAwsEventbridgeOutputReference",
		reflect.TypeOf((*EventTriggerEventProcessorsAwsEventbridgeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "configAccountId", GoGetter: "ConfigAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "configAccountIdInput", GoGetter: "ConfigAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "configRegion", GoGetter: "ConfigRegion"},
			_jsii_.MemberProperty{JsiiProperty: "configRegionInput", GoGetter: "ConfigRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigAccountId", GoMethod: "ResetConfigAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigRegion", GoMethod: "ResetConfigRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventTriggerEventProcessorsAwsEventbridgeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.eventTrigger.EventTriggerEventProcessorsOutputReference",
		reflect.TypeOf((*EventTriggerEventProcessorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsEventbridge", GoGetter: "AwsEventbridge"},
			_jsii_.MemberProperty{JsiiProperty: "awsEventbridgeInput", GoGetter: "AwsEventbridgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putAwsEventbridge", GoMethod: "PutAwsEventbridge"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsEventbridge", GoMethod: "ResetAwsEventbridge"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EventTriggerEventProcessorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
