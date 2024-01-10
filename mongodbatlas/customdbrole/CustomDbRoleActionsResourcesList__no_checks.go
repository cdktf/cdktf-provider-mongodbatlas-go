// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customdbrole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomDbRoleActionsResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomDbRoleActionsResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomDbRoleActionsResourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsResourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsResourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomDbRoleActionsResourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

