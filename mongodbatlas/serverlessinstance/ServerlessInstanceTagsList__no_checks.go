// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serverlessinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerlessInstanceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServerlessInstanceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServerlessInstanceTagsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServerlessInstanceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServerlessInstanceTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

