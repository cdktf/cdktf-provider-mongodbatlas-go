// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutAdvancedConfigurationParameters(value *ClusterAdvancedConfiguration) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutBiConnectorConfigParameters(value *ClusterBiConnectorConfig) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutLabelsParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutReplicationSpecsParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Cluster) validatePutTimeoutsParameters(value *ClusterTimeouts) error {
	return nil
}

func validateCluster_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCluster_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAcceptDataRisksAndForceReplicaSetReconfigParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAutoScalingComputeEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAutoScalingComputeScaleDownEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetAutoScalingDiskGbEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetBackingProviderNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetBackupEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetCloudBackupParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetClusterTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetDiskSizeGbParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetEncryptionAtRestProviderParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetMongoDbMajorVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetNumShardsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetPausedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetPitEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderAutoScalingComputeMaxInstanceSizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderAutoScalingComputeMinInstanceSizeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderDiskIopsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderDiskTypeNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderEncryptEbsVolumeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderInstanceSizeNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderRegionNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProviderVolumeTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetRedactClientLogDataParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetReplicationFactorParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetRetainBackupsEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetTerminationProtectionEnabledParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Cluster) validateSetVersionReleaseSystemParameters(val *string) error {
	return nil
}

func validateNewClusterParameters(scope constructs.Construct, id *string, config *ClusterConfig) error {
	return nil
}

