// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package advancedcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AdvancedClusterAdvancedConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterAdvancedConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAdvancedClusterAdvancedConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

