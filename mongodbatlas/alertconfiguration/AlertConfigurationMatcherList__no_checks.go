// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alertconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertConfigurationMatcherList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertConfigurationMatcherList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationMatcherList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationMatcherList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationMatcherList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationMatcherList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertConfigurationMatcherListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

