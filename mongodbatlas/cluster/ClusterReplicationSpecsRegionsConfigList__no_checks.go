// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cluster

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterReplicationSpecsRegionsConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterReplicationSpecsRegionsConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

