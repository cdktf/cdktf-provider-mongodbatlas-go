// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package advancedcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdvancedClusterLabelsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AdvancedClusterLabelsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterLabelsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterLabelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterLabelsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterLabelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAdvancedClusterLabelsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

