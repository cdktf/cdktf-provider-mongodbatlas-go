// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datalakepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataLakePipelineIngestionSchedulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataLakePipelineIngestionSchedulesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataLakePipelineIngestionSchedulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineIngestionSchedulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineIngestionSchedulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineIngestionSchedulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataLakePipelineIngestionSchedulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

