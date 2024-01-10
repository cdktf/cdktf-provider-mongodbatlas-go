// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datalakepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataLakePipelineSnapshotsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataLakePipelineSnapshotsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataLakePipelineSnapshotsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineSnapshotsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineSnapshotsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineSnapshotsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataLakePipelineSnapshotsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

