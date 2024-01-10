// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package advancedcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdvancedClusterConnectionStringsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AdvancedClusterConnectionStringsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AdvancedClusterConnectionStringsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterConnectionStringsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterConnectionStringsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterConnectionStringsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAdvancedClusterConnectionStringsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

