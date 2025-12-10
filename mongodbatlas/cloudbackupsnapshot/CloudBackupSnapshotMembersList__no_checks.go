// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudbackupsnapshot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudBackupSnapshotMembersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudBackupSnapshotMembersList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudBackupSnapshotMembersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudBackupSnapshotMembersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudBackupSnapshotMembersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudBackupSnapshotMembersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudBackupSnapshotMembersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

