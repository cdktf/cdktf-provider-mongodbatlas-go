// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package searchdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SearchDeploymentSpecsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SearchDeploymentSpecsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SearchDeploymentSpecsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SearchDeploymentSpecsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SearchDeploymentSpecsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SearchDeploymentSpecsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SearchDeploymentSpecsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSearchDeploymentSpecsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

