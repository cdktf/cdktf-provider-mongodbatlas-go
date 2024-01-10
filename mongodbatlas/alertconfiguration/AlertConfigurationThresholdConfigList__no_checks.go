// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alertconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertConfigurationThresholdConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertConfigurationThresholdConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertConfigurationThresholdConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationThresholdConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationThresholdConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationThresholdConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationThresholdConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertConfigurationThresholdConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

