package cloudprovidersnapshotrestorejob

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJob",
		reflect.TypeOf((*CloudProviderSnapshotRestoreJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cancelled", GoGetter: "Cancelled"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNameInput", GoGetter: "ClusterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryType", GoGetter: "DeliveryType"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryTypeConfig", GoGetter: "DeliveryTypeConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryTypeConfigInput", GoGetter: "DeliveryTypeConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryTypeInput", GoGetter: "DeliveryTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryUrl", GoGetter: "DeliveryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "expired", GoGetter: "Expired"},
			_jsii_.MemberProperty{JsiiProperty: "expiresAt", GoGetter: "ExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "finishedAt", GoGetter: "FinishedAt"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDeliveryTypeConfig", GoMethod: "PutDeliveryTypeConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryType", GoMethod: "ResetDeliveryType"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeliveryTypeConfig", GoMethod: "ResetDeliveryTypeConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotId", GoGetter: "SnapshotId"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotIdInput", GoGetter: "SnapshotIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotRestoreJobId", GoGetter: "SnapshotRestoreJobId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timestamp", GoGetter: "Timestamp"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudProviderSnapshotRestoreJob{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJobConfig",
		reflect.TypeOf((*CloudProviderSnapshotRestoreJobConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJobDeliveryTypeConfig",
		reflect.TypeOf((*CloudProviderSnapshotRestoreJobDeliveryTypeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-mongodbatlas.cloudProviderSnapshotRestoreJob.CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference",
		reflect.TypeOf((*CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automated", GoGetter: "Automated"},
			_jsii_.MemberProperty{JsiiProperty: "automatedInput", GoGetter: "AutomatedInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "download", GoGetter: "Download"},
			_jsii_.MemberProperty{JsiiProperty: "downloadInput", GoGetter: "DownloadInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "oplogInc", GoGetter: "OplogInc"},
			_jsii_.MemberProperty{JsiiProperty: "oplogIncInput", GoGetter: "OplogIncInput"},
			_jsii_.MemberProperty{JsiiProperty: "oplogTs", GoGetter: "OplogTs"},
			_jsii_.MemberProperty{JsiiProperty: "oplogTsInput", GoGetter: "OplogTsInput"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTime", GoGetter: "PointInTime"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeInput", GoGetter: "PointInTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeUtcSeconds", GoGetter: "PointInTimeUtcSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "pointInTimeUtcSecondsInput", GoGetter: "PointInTimeUtcSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomated", GoMethod: "ResetAutomated"},
			_jsii_.MemberMethod{JsiiMethod: "resetDownload", GoMethod: "ResetDownload"},
			_jsii_.MemberMethod{JsiiMethod: "resetOplogInc", GoMethod: "ResetOplogInc"},
			_jsii_.MemberMethod{JsiiMethod: "resetOplogTs", GoMethod: "ResetOplogTs"},
			_jsii_.MemberMethod{JsiiMethod: "resetPointInTime", GoMethod: "ResetPointInTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetPointInTimeUtcSeconds", GoMethod: "ResetPointInTimeUtcSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetClusterName", GoMethod: "ResetTargetClusterName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetProjectId", GoMethod: "ResetTargetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetClusterName", GoGetter: "TargetClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "targetClusterNameInput", GoGetter: "TargetClusterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetProjectId", GoGetter: "TargetProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "targetProjectIdInput", GoGetter: "TargetProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudProviderSnapshotRestoreJobDeliveryTypeConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
