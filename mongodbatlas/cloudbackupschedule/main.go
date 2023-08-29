// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbackupschedule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		reflect.TypeOf((*CloudBackupSchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoExportEnabled", GoGetter: "AutoExportEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "autoExportEnabledInput", GoGetter: "AutoExportEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNameInput", GoGetter: "ClusterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "copySettings", GoGetter: "CopySettings"},
			_jsii_.MemberProperty{JsiiProperty: "copySettingsInput", GoGetter: "CopySettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "export", GoGetter: "Export"},
			_jsii_.MemberProperty{JsiiProperty: "exportInput", GoGetter: "ExportInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idPolicy", GoGetter: "IdPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "nextSnapshot", GoGetter: "NextSnapshot"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemDaily", GoGetter: "PolicyItemDaily"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemDailyInput", GoGetter: "PolicyItemDailyInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemHourly", GoGetter: "PolicyItemHourly"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemHourlyInput", GoGetter: "PolicyItemHourlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemMonthly", GoGetter: "PolicyItemMonthly"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemMonthlyInput", GoGetter: "PolicyItemMonthlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemWeekly", GoGetter: "PolicyItemWeekly"},
			_jsii_.MemberProperty{JsiiProperty: "policyItemWeeklyInput", GoGetter: "PolicyItemWeeklyInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCopySettings", GoMethod: "PutCopySettings"},
			_jsii_.MemberMethod{JsiiMethod: "putExport", GoMethod: "PutExport"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyItemDaily", GoMethod: "PutPolicyItemDaily"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyItemHourly", GoMethod: "PutPolicyItemHourly"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyItemMonthly", GoMethod: "PutPolicyItemMonthly"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyItemWeekly", GoMethod: "PutPolicyItemWeekly"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "referenceHourOfDay", GoGetter: "ReferenceHourOfDay"},
			_jsii_.MemberProperty{JsiiProperty: "referenceHourOfDayInput", GoGetter: "ReferenceHourOfDayInput"},
			_jsii_.MemberProperty{JsiiProperty: "referenceMinuteOfHour", GoGetter: "ReferenceMinuteOfHour"},
			_jsii_.MemberProperty{JsiiProperty: "referenceMinuteOfHourInput", GoGetter: "ReferenceMinuteOfHourInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoExportEnabled", GoMethod: "ResetAutoExportEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCopySettings", GoMethod: "ResetCopySettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetExport", GoMethod: "ResetExport"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyItemDaily", GoMethod: "ResetPolicyItemDaily"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyItemHourly", GoMethod: "ResetPolicyItemHourly"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyItemMonthly", GoMethod: "ResetPolicyItemMonthly"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyItemWeekly", GoMethod: "ResetPolicyItemWeekly"},
			_jsii_.MemberMethod{JsiiMethod: "resetReferenceHourOfDay", GoMethod: "ResetReferenceHourOfDay"},
			_jsii_.MemberMethod{JsiiMethod: "resetReferenceMinuteOfHour", GoMethod: "ResetReferenceMinuteOfHour"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestoreWindowDays", GoMethod: "ResetRestoreWindowDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdateSnapshots", GoMethod: "ResetUpdateSnapshots"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseOrgAndGroupNamesInExportPrefix", GoMethod: "ResetUseOrgAndGroupNamesInExportPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "restoreWindowDays", GoGetter: "RestoreWindowDays"},
			_jsii_.MemberProperty{JsiiProperty: "restoreWindowDaysInput", GoGetter: "RestoreWindowDaysInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updateSnapshots", GoGetter: "UpdateSnapshots"},
			_jsii_.MemberProperty{JsiiProperty: "updateSnapshotsInput", GoGetter: "UpdateSnapshotsInput"},
			_jsii_.MemberProperty{JsiiProperty: "useOrgAndGroupNamesInExportPrefix", GoGetter: "UseOrgAndGroupNamesInExportPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "useOrgAndGroupNamesInExportPrefixInput", GoGetter: "UseOrgAndGroupNamesInExportPrefixInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedule{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleConfig",
		reflect.TypeOf((*CloudBackupScheduleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleCopySettings",
		reflect.TypeOf((*CloudBackupScheduleCopySettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleCopySettingsList",
		reflect.TypeOf((*CloudBackupScheduleCopySettingsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupScheduleCopySettingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleCopySettingsOutputReference",
		reflect.TypeOf((*CloudBackupScheduleCopySettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudProvider", GoGetter: "CloudProvider"},
			_jsii_.MemberProperty{JsiiProperty: "cloudProviderInput", GoGetter: "CloudProviderInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencies", GoGetter: "Frequencies"},
			_jsii_.MemberProperty{JsiiProperty: "frequenciesInput", GoGetter: "FrequenciesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "regionName", GoGetter: "RegionName"},
			_jsii_.MemberProperty{JsiiProperty: "regionNameInput", GoGetter: "RegionNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSpecId", GoGetter: "ReplicationSpecId"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSpecIdInput", GoGetter: "ReplicationSpecIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudProvider", GoMethod: "ResetCloudProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrequencies", GoMethod: "ResetFrequencies"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegionName", GoMethod: "ResetRegionName"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplicationSpecId", GoMethod: "ResetReplicationSpecId"},
			_jsii_.MemberMethod{JsiiMethod: "resetShouldCopyOplogs", GoMethod: "ResetShouldCopyOplogs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shouldCopyOplogs", GoGetter: "ShouldCopyOplogs"},
			_jsii_.MemberProperty{JsiiProperty: "shouldCopyOplogsInput", GoGetter: "ShouldCopyOplogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupScheduleCopySettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleExport",
		reflect.TypeOf((*CloudBackupScheduleExport)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupScheduleExportOutputReference",
		reflect.TypeOf((*CloudBackupScheduleExportOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "exportBucketId", GoGetter: "ExportBucketId"},
			_jsii_.MemberProperty{JsiiProperty: "exportBucketIdInput", GoGetter: "ExportBucketIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyType", GoGetter: "FrequencyType"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyTypeInput", GoGetter: "FrequencyTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetExportBucketId", GoMethod: "ResetExportBucketId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrequencyType", GoMethod: "ResetFrequencyType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupScheduleExportOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemDaily",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemDaily)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemDailyOutputReference",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemDailyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyInterval", GoGetter: "FrequencyInterval"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyIntervalInput", GoGetter: "FrequencyIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyType", GoGetter: "FrequencyType"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnit", GoGetter: "RetentionUnit"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnitInput", GoGetter: "RetentionUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValue", GoGetter: "RetentionValue"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValueInput", GoGetter: "RetentionValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedulePolicyItemDailyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemHourly",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemHourly)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemHourlyOutputReference",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemHourlyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyInterval", GoGetter: "FrequencyInterval"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyIntervalInput", GoGetter: "FrequencyIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyType", GoGetter: "FrequencyType"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnit", GoGetter: "RetentionUnit"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnitInput", GoGetter: "RetentionUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValue", GoGetter: "RetentionValue"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValueInput", GoGetter: "RetentionValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedulePolicyItemHourlyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemMonthly",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemMonthly)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemMonthlyList",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemMonthlyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedulePolicyItemMonthlyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemMonthlyOutputReference",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemMonthlyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyInterval", GoGetter: "FrequencyInterval"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyIntervalInput", GoGetter: "FrequencyIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyType", GoGetter: "FrequencyType"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnit", GoGetter: "RetentionUnit"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnitInput", GoGetter: "RetentionUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValue", GoGetter: "RetentionValue"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValueInput", GoGetter: "RetentionValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedulePolicyItemMonthlyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemWeekly",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemWeekly)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemWeeklyList",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemWeeklyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedulePolicyItemWeeklyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedulePolicyItemWeeklyOutputReference",
		reflect.TypeOf((*CloudBackupSchedulePolicyItemWeeklyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyInterval", GoGetter: "FrequencyInterval"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyIntervalInput", GoGetter: "FrequencyIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "frequencyType", GoGetter: "FrequencyType"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnit", GoGetter: "RetentionUnit"},
			_jsii_.MemberProperty{JsiiProperty: "retentionUnitInput", GoGetter: "RetentionUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValue", GoGetter: "RetentionValue"},
			_jsii_.MemberProperty{JsiiProperty: "retentionValueInput", GoGetter: "RetentionValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudBackupSchedulePolicyItemWeeklyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
