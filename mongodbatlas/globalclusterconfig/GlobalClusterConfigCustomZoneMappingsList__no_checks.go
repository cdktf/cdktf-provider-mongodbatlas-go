// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package globalclusterconfig

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlobalClusterConfigCustomZoneMappingsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlobalClusterConfigCustomZoneMappingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigCustomZoneMappingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigCustomZoneMappingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigCustomZoneMappingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigCustomZoneMappingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlobalClusterConfigCustomZoneMappingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

