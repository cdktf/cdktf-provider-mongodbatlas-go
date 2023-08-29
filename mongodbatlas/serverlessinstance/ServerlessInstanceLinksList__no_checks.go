// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serverlessinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerlessInstanceLinksList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServerlessInstanceLinksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceLinksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceLinksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceLinksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceLinksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServerlessInstanceLinksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

