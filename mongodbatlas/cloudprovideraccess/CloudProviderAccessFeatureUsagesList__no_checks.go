// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudprovideraccess

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudProviderAccessFeatureUsagesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudProviderAccessFeatureUsagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessFeatureUsagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessFeatureUsagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessFeatureUsagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudProviderAccessFeatureUsagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

