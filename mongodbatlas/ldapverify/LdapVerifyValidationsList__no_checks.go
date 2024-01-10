// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ldapverify

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LdapVerifyValidationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LdapVerifyValidationsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LdapVerifyValidationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LdapVerifyValidationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LdapVerifyValidationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LdapVerifyValidationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLdapVerifyValidationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

