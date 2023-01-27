//go:build no_runtime_type_checking

package alertconfiguration

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertConfigurationNotificationList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertConfigurationNotificationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationNotificationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationNotificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationNotificationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertConfigurationNotificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertConfigurationNotificationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

