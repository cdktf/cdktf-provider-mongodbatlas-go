// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudprovideraccesssetup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudProviderAccessSetupGcpConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudProviderAccessSetupGcpConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudProviderAccessSetupGcpConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessSetupGcpConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessSetupGcpConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudProviderAccessSetupGcpConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudProviderAccessSetupGcpConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

