// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package globalclusterconfig

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlobalClusterConfigManagedNamespacesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlobalClusterConfigManagedNamespacesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

