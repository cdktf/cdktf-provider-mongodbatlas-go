// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedCluster",
		reflect.TypeOf((*AdvancedCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptDataRisksAndForceReplicaSetReconfig", GoGetter: "AcceptDataRisksAndForceReplicaSetReconfig"},
			_jsii_.MemberProperty{JsiiProperty: "acceptDataRisksAndForceReplicaSetReconfigInput", GoGetter: "AcceptDataRisksAndForceReplicaSetReconfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "advancedConfiguration", GoGetter: "AdvancedConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "advancedConfigurationInput", GoGetter: "AdvancedConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupEnabled", GoGetter: "BackupEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "backupEnabledInput", GoGetter: "BackupEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "biConnectorConfig", GoGetter: "BiConnectorConfig"},
			_jsii_.MemberProperty{JsiiProperty: "biConnectorConfigInput", GoGetter: "BiConnectorConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterType", GoGetter: "ClusterType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterTypeInput", GoGetter: "ClusterTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionStrings", GoGetter: "ConnectionStrings"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createDate", GoGetter: "CreateDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "diskSizeGb", GoGetter: "DiskSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "diskSizeGbInput", GoGetter: "DiskSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtRestProvider", GoGetter: "EncryptionAtRestProvider"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtRestProviderInput", GoGetter: "EncryptionAtRestProviderInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mongoDbMajorVersion", GoGetter: "MongoDbMajorVersion"},
			_jsii_.MemberProperty{JsiiProperty: "mongoDbMajorVersionInput", GoGetter: "MongoDbMajorVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "mongoDbVersion", GoGetter: "MongoDbVersion"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "paused", GoGetter: "Paused"},
			_jsii_.MemberProperty{JsiiProperty: "pausedInput", GoGetter: "PausedInput"},
			_jsii_.MemberProperty{JsiiProperty: "pitEnabled", GoGetter: "PitEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "pitEnabledInput", GoGetter: "PitEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAdvancedConfiguration", GoMethod: "PutAdvancedConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putBiConnectorConfig", GoMethod: "PutBiConnectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putLabels", GoMethod: "PutLabels"},
			_jsii_.MemberMethod{JsiiMethod: "putReplicationSpecs", GoMethod: "PutReplicationSpecs"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSpecs", GoGetter: "ReplicationSpecs"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSpecsInput", GoGetter: "ReplicationSpecsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceptDataRisksAndForceReplicaSetReconfig", GoMethod: "ResetAcceptDataRisksAndForceReplicaSetReconfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedConfiguration", GoMethod: "ResetAdvancedConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupEnabled", GoMethod: "ResetBackupEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetBiConnectorConfig", GoMethod: "ResetBiConnectorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskSizeGb", GoMethod: "ResetDiskSizeGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionAtRestProvider", GoMethod: "ResetEncryptionAtRestProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetMongoDbMajorVersion", GoMethod: "ResetMongoDbMajorVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPaused", GoMethod: "ResetPaused"},
			_jsii_.MemberMethod{JsiiMethod: "resetPitEnabled", GoMethod: "ResetPitEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetainBackupsEnabled", GoMethod: "ResetRetainBackupsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootCertType", GoMethod: "ResetRootCertType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerminationProtectionEnabled", GoMethod: "ResetTerminationProtectionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersionReleaseSystem", GoMethod: "ResetVersionReleaseSystem"},
			_jsii_.MemberProperty{JsiiProperty: "retainBackupsEnabled", GoGetter: "RetainBackupsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "retainBackupsEnabledInput", GoGetter: "RetainBackupsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "rootCertType", GoGetter: "RootCertType"},
			_jsii_.MemberProperty{JsiiProperty: "rootCertTypeInput", GoGetter: "RootCertTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "stateName", GoGetter: "StateName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtectionEnabled", GoGetter: "TerminationProtectionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtectionEnabledInput", GoGetter: "TerminationProtectionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "versionReleaseSystem", GoGetter: "VersionReleaseSystem"},
			_jsii_.MemberProperty{JsiiProperty: "versionReleaseSystemInput", GoGetter: "VersionReleaseSystemInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterAdvancedConfiguration",
		reflect.TypeOf((*AdvancedClusterAdvancedConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterAdvancedConfigurationList",
		reflect.TypeOf((*AdvancedClusterAdvancedConfigurationList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AdvancedClusterAdvancedConfigurationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterAdvancedConfigurationOutputReference",
		reflect.TypeOf((*AdvancedClusterAdvancedConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultReadConcern", GoGetter: "DefaultReadConcern"},
			_jsii_.MemberProperty{JsiiProperty: "defaultReadConcernInput", GoGetter: "DefaultReadConcernInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultWriteConcern", GoGetter: "DefaultWriteConcern"},
			_jsii_.MemberProperty{JsiiProperty: "defaultWriteConcernInput", GoGetter: "DefaultWriteConcernInput"},
			_jsii_.MemberProperty{JsiiProperty: "failIndexKeyTooLong", GoGetter: "FailIndexKeyTooLong"},
			_jsii_.MemberProperty{JsiiProperty: "failIndexKeyTooLongInput", GoGetter: "FailIndexKeyTooLongInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "javascriptEnabled", GoGetter: "JavascriptEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "javascriptEnabledInput", GoGetter: "JavascriptEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumEnabledTlsProtocol", GoGetter: "MinimumEnabledTlsProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "minimumEnabledTlsProtocolInput", GoGetter: "MinimumEnabledTlsProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "noTableScan", GoGetter: "NoTableScan"},
			_jsii_.MemberProperty{JsiiProperty: "noTableScanInput", GoGetter: "NoTableScanInput"},
			_jsii_.MemberProperty{JsiiProperty: "oplogMinRetentionHours", GoGetter: "OplogMinRetentionHours"},
			_jsii_.MemberProperty{JsiiProperty: "oplogMinRetentionHoursInput", GoGetter: "OplogMinRetentionHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "oplogSizeMb", GoGetter: "OplogSizeMb"},
			_jsii_.MemberProperty{JsiiProperty: "oplogSizeMbInput", GoGetter: "OplogSizeMbInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultReadConcern", GoMethod: "ResetDefaultReadConcern"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultWriteConcern", GoMethod: "ResetDefaultWriteConcern"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailIndexKeyTooLong", GoMethod: "ResetFailIndexKeyTooLong"},
			_jsii_.MemberMethod{JsiiMethod: "resetJavascriptEnabled", GoMethod: "ResetJavascriptEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumEnabledTlsProtocol", GoMethod: "ResetMinimumEnabledTlsProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoTableScan", GoMethod: "ResetNoTableScan"},
			_jsii_.MemberMethod{JsiiMethod: "resetOplogMinRetentionHours", GoMethod: "ResetOplogMinRetentionHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetOplogSizeMb", GoMethod: "ResetOplogSizeMb"},
			_jsii_.MemberMethod{JsiiMethod: "resetSampleRefreshIntervalBiConnector", GoMethod: "ResetSampleRefreshIntervalBiConnector"},
			_jsii_.MemberMethod{JsiiMethod: "resetSampleSizeBiConnector", GoMethod: "ResetSampleSizeBiConnector"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransactionLifetimeLimitSeconds", GoMethod: "ResetTransactionLifetimeLimitSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sampleRefreshIntervalBiConnector", GoGetter: "SampleRefreshIntervalBiConnector"},
			_jsii_.MemberProperty{JsiiProperty: "sampleRefreshIntervalBiConnectorInput", GoGetter: "SampleRefreshIntervalBiConnectorInput"},
			_jsii_.MemberProperty{JsiiProperty: "sampleSizeBiConnector", GoGetter: "SampleSizeBiConnector"},
			_jsii_.MemberProperty{JsiiProperty: "sampleSizeBiConnectorInput", GoGetter: "SampleSizeBiConnectorInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transactionLifetimeLimitSeconds", GoGetter: "TransactionLifetimeLimitSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "transactionLifetimeLimitSecondsInput", GoGetter: "TransactionLifetimeLimitSecondsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterAdvancedConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterBiConnectorConfig",
		reflect.TypeOf((*AdvancedClusterBiConnectorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterBiConnectorConfigOutputReference",
		reflect.TypeOf((*AdvancedClusterBiConnectorConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "readPreference", GoGetter: "ReadPreference"},
			_jsii_.MemberProperty{JsiiProperty: "readPreferenceInput", GoGetter: "ReadPreferenceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadPreference", GoMethod: "ResetReadPreference"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterBiConnectorConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConfig",
		reflect.TypeOf((*AdvancedClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStrings",
		reflect.TypeOf((*AdvancedClusterConnectionStrings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsList",
		reflect.TypeOf((*AdvancedClusterConnectionStringsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterConnectionStringsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsOutputReference",
		reflect.TypeOf((*AdvancedClusterConnectionStringsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "private", GoGetter: "Private"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpoint", GoGetter: "PrivateEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "privateSrv", GoGetter: "PrivateSrv"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "standard", GoGetter: "Standard"},
			_jsii_.MemberProperty{JsiiProperty: "standardSrv", GoGetter: "StandardSrv"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterConnectionStringsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsPrivateEndpoint",
		reflect.TypeOf((*AdvancedClusterConnectionStringsPrivateEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsPrivateEndpointEndpoints",
		reflect.TypeOf((*AdvancedClusterConnectionStringsPrivateEndpointEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsPrivateEndpointEndpointsList",
		reflect.TypeOf((*AdvancedClusterConnectionStringsPrivateEndpointEndpointsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterConnectionStringsPrivateEndpointEndpointsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsPrivateEndpointEndpointsOutputReference",
		reflect.TypeOf((*AdvancedClusterConnectionStringsPrivateEndpointEndpointsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpointId", GoGetter: "EndpointId"},
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
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterConnectionStringsPrivateEndpointEndpointsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsPrivateEndpointList",
		reflect.TypeOf((*AdvancedClusterConnectionStringsPrivateEndpointList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterConnectionStringsPrivateEndpointList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterConnectionStringsPrivateEndpointOutputReference",
		reflect.TypeOf((*AdvancedClusterConnectionStringsPrivateEndpointOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionString", GoGetter: "ConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "srvConnectionString", GoGetter: "SrvConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "srvShardOptimizedConnectionString", GoGetter: "SrvShardOptimizedConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterConnectionStringsPrivateEndpointOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterLabels",
		reflect.TypeOf((*AdvancedClusterLabels)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterLabelsList",
		reflect.TypeOf((*AdvancedClusterLabelsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AdvancedClusterLabelsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterLabelsOutputReference",
		reflect.TypeOf((*AdvancedClusterLabelsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterLabelsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecs",
		reflect.TypeOf((*AdvancedClusterReplicationSpecs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsList",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AdvancedClusterReplicationSpecsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerId", GoGetter: "ContainerId"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "numShards", GoGetter: "NumShards"},
			_jsii_.MemberProperty{JsiiProperty: "numShardsInput", GoGetter: "NumShardsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRegionConfigs", GoMethod: "PutRegionConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "regionConfigs", GoGetter: "RegionConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "regionConfigsInput", GoGetter: "RegionConfigsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumShards", GoMethod: "ResetNumShards"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneName", GoMethod: "ResetZoneName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "zoneName", GoGetter: "ZoneName"},
			_jsii_.MemberProperty{JsiiProperty: "zoneNameInput", GoGetter: "ZoneNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigs",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScaling)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScalingOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScalingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnabled", GoGetter: "ComputeEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnabledInput", GoGetter: "ComputeEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "computeMaxInstanceSize", GoGetter: "ComputeMaxInstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "computeMaxInstanceSizeInput", GoGetter: "ComputeMaxInstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeMinInstanceSize", GoGetter: "ComputeMinInstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "computeMinInstanceSizeInput", GoGetter: "ComputeMinInstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeScaleDownEnabled", GoGetter: "ComputeScaleDownEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "computeScaleDownEnabledInput", GoGetter: "ComputeScaleDownEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskGbEnabled", GoGetter: "DiskGbEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "diskGbEnabledInput", GoGetter: "DiskGbEnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetComputeEnabled", GoMethod: "ResetComputeEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeMaxInstanceSize", GoMethod: "ResetComputeMaxInstanceSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeMinInstanceSize", GoMethod: "ResetComputeMinInstanceSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeScaleDownEnabled", GoMethod: "ResetComputeScaleDownEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskGbEnabled", GoMethod: "ResetDiskGbEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAnalyticsAutoScalingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecsOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskIops", GoGetter: "DiskIops"},
			_jsii_.MemberProperty{JsiiProperty: "diskIopsInput", GoGetter: "DiskIopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsVolumeType", GoGetter: "EbsVolumeType"},
			_jsii_.MemberProperty{JsiiProperty: "ebsVolumeTypeInput", GoGetter: "EbsVolumeTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "instanceSize", GoGetter: "InstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "instanceSizeInput", GoGetter: "InstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCount", GoGetter: "NodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCountInput", GoGetter: "NodeCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskIops", GoMethod: "ResetDiskIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsVolumeType", GoMethod: "ResetEbsVolumeType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNodeCount", GoMethod: "ResetNodeCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAnalyticsSpecsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAutoScaling",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsAutoScaling)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnabled", GoGetter: "ComputeEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnabledInput", GoGetter: "ComputeEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "computeMaxInstanceSize", GoGetter: "ComputeMaxInstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "computeMaxInstanceSizeInput", GoGetter: "ComputeMaxInstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeMinInstanceSize", GoGetter: "ComputeMinInstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "computeMinInstanceSizeInput", GoGetter: "ComputeMinInstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeScaleDownEnabled", GoGetter: "ComputeScaleDownEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "computeScaleDownEnabledInput", GoGetter: "ComputeScaleDownEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskGbEnabled", GoGetter: "DiskGbEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "diskGbEnabledInput", GoGetter: "DiskGbEnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetComputeEnabled", GoMethod: "ResetComputeEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeMaxInstanceSize", GoMethod: "ResetComputeMaxInstanceSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeMinInstanceSize", GoMethod: "ResetComputeMinInstanceSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeScaleDownEnabled", GoMethod: "ResetComputeScaleDownEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskGbEnabled", GoMethod: "ResetDiskGbEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsAutoScalingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskIops", GoGetter: "DiskIops"},
			_jsii_.MemberProperty{JsiiProperty: "diskIopsInput", GoGetter: "DiskIopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsVolumeType", GoGetter: "EbsVolumeType"},
			_jsii_.MemberProperty{JsiiProperty: "ebsVolumeTypeInput", GoGetter: "EbsVolumeTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "instanceSize", GoGetter: "InstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "instanceSizeInput", GoGetter: "InstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCount", GoGetter: "NodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCountInput", GoGetter: "NodeCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskIops", GoMethod: "ResetDiskIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsVolumeType", GoMethod: "ResetEbsVolumeType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNodeCount", GoMethod: "ResetNodeCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsElectableSpecsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsList",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "analyticsAutoScaling", GoGetter: "AnalyticsAutoScaling"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsAutoScalingInput", GoGetter: "AnalyticsAutoScalingInput"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsSpecs", GoGetter: "AnalyticsSpecs"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsSpecsInput", GoGetter: "AnalyticsSpecsInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoScaling", GoGetter: "AutoScaling"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingInput", GoGetter: "AutoScalingInput"},
			_jsii_.MemberProperty{JsiiProperty: "backingProviderName", GoGetter: "BackingProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "backingProviderNameInput", GoGetter: "BackingProviderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "electableSpecs", GoGetter: "ElectableSpecs"},
			_jsii_.MemberProperty{JsiiProperty: "electableSpecsInput", GoGetter: "ElectableSpecsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "providerNameInput", GoGetter: "ProviderNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAnalyticsAutoScaling", GoMethod: "PutAnalyticsAutoScaling"},
			_jsii_.MemberMethod{JsiiMethod: "putAnalyticsSpecs", GoMethod: "PutAnalyticsSpecs"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoScaling", GoMethod: "PutAutoScaling"},
			_jsii_.MemberMethod{JsiiMethod: "putElectableSpecs", GoMethod: "PutElectableSpecs"},
			_jsii_.MemberMethod{JsiiMethod: "putReadOnlySpecs", GoMethod: "PutReadOnlySpecs"},
			_jsii_.MemberProperty{JsiiProperty: "readOnlySpecs", GoGetter: "ReadOnlySpecs"},
			_jsii_.MemberProperty{JsiiProperty: "readOnlySpecsInput", GoGetter: "ReadOnlySpecsInput"},
			_jsii_.MemberProperty{JsiiProperty: "regionName", GoGetter: "RegionName"},
			_jsii_.MemberProperty{JsiiProperty: "regionNameInput", GoGetter: "RegionNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnalyticsAutoScaling", GoMethod: "ResetAnalyticsAutoScaling"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnalyticsSpecs", GoMethod: "ResetAnalyticsSpecs"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoScaling", GoMethod: "ResetAutoScaling"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackingProviderName", GoMethod: "ResetBackingProviderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetElectableSpecs", GoMethod: "ResetElectableSpecs"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadOnlySpecs", GoMethod: "ResetReadOnlySpecs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecsOutputReference",
		reflect.TypeOf((*AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskIops", GoGetter: "DiskIops"},
			_jsii_.MemberProperty{JsiiProperty: "diskIopsInput", GoGetter: "DiskIopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsVolumeType", GoGetter: "EbsVolumeType"},
			_jsii_.MemberProperty{JsiiProperty: "ebsVolumeTypeInput", GoGetter: "EbsVolumeTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "instanceSize", GoGetter: "InstanceSize"},
			_jsii_.MemberProperty{JsiiProperty: "instanceSizeInput", GoGetter: "InstanceSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCount", GoGetter: "NodeCount"},
			_jsii_.MemberProperty{JsiiProperty: "nodeCountInput", GoGetter: "NodeCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskIops", GoMethod: "ResetDiskIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsVolumeType", GoMethod: "ResetEbsVolumeType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNodeCount", GoMethod: "ResetNodeCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterReplicationSpecsRegionConfigsReadOnlySpecsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterTags",
		reflect.TypeOf((*AdvancedClusterTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterTagsList",
		reflect.TypeOf((*AdvancedClusterTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_AdvancedClusterTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterTagsOutputReference",
		reflect.TypeOf((*AdvancedClusterTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterTimeouts",
		reflect.TypeOf((*AdvancedClusterTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.advancedCluster.AdvancedClusterTimeoutsOutputReference",
		reflect.TypeOf((*AdvancedClusterTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_AdvancedClusterTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
