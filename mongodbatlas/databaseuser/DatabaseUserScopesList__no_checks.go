// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package databaseuser

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseUserScopesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseUserScopesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseUserScopesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserScopesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserScopesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserScopesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserScopesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseUserScopesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

