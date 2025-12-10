// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterSnapshotBackupPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClusterSnapshotBackupPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterSnapshotBackupPolicyList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterSnapshotBackupPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterSnapshotBackupPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterSnapshotBackupPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterSnapshotBackupPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

