// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterConnectionStringsPrivateEndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClusterConnectionStringsPrivateEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterConnectionStringsPrivateEndpointList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterConnectionStringsPrivateEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterConnectionStringsPrivateEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterConnectionStringsPrivateEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterConnectionStringsPrivateEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

