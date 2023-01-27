//go:build no_runtime_type_checking

package databaseuser

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseUserRolesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseUserRolesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserRolesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserRolesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserRolesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseUserRolesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseUserRolesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

