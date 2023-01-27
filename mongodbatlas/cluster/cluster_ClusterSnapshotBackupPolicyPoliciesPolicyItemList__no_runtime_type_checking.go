//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterSnapshotBackupPolicyPoliciesPolicyItemList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterSnapshotBackupPolicyPoliciesPolicyItemList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterSnapshotBackupPolicyPoliciesPolicyItemList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterSnapshotBackupPolicyPoliciesPolicyItemList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterSnapshotBackupPolicyPoliciesPolicyItemList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterSnapshotBackupPolicyPoliciesPolicyItemListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

