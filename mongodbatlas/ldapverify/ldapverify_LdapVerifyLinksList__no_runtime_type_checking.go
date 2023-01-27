//go:build no_runtime_type_checking

package ldapverify

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LdapVerifyLinksList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LdapVerifyLinksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LdapVerifyLinksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LdapVerifyLinksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LdapVerifyLinksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLdapVerifyLinksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

