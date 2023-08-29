// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudprovideraccesssetup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudProviderAccessSetupAwsConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudProviderAccessSetupAwsConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessSetupAwsConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessSetupAwsConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessSetupAwsConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudProviderAccessSetupAwsConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

