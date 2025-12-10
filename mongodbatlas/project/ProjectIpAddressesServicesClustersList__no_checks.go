// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package project

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectIpAddressesServicesClustersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectIpAddressesServicesClustersList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectIpAddressesServicesClustersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectIpAddressesServicesClustersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectIpAddressesServicesClustersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectIpAddressesServicesClustersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectIpAddressesServicesClustersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

