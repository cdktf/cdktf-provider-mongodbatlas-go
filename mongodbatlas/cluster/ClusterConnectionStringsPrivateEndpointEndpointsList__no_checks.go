// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterConnectionStringsPrivateEndpointEndpointsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClusterConnectionStringsPrivateEndpointEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterConnectionStringsPrivateEndpointEndpointsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterConnectionStringsPrivateEndpointEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterConnectionStringsPrivateEndpointEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterConnectionStringsPrivateEndpointEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterConnectionStringsPrivateEndpointEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

