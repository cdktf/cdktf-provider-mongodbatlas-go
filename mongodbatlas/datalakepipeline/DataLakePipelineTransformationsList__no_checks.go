// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datalakepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataLakePipelineTransformationsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataLakePipelineTransformationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineTransformationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineTransformationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineTransformationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataLakePipelineTransformationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataLakePipelineTransformationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

