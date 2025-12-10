// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package customdbrole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomDbRoleActionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomDbRoleActionsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomDbRoleActionsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomDbRoleActionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomDbRoleActionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

