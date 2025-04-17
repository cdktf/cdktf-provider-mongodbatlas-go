// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package resourcepolicy

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourcePolicy) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validatePutPoliciesParameters(value interface{}) error {
	return nil
}

func validateResourcePolicy_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateResourcePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourcePolicy_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateResourcePolicy_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetOrgIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourcePolicy) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewResourcePolicyParameters(scope constructs.Construct, id *string, config *ResourcePolicyConfig) error {
	return nil
}

