// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package projectapikey

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProjectApiKeyProjectAssignmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProjectApiKeyProjectAssignmentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

