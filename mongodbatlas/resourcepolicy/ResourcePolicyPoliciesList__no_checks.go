// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package resourcepolicy

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourcePolicyPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicyPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicyPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicyPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicyPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicyPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicyPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResourcePolicyPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

