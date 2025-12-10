// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package advancedcluster

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AdvancedClusterTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AdvancedClusterTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AdvancedClusterTagsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AdvancedClusterTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAdvancedClusterTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

