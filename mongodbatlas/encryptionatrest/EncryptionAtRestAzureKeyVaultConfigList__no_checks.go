// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package encryptionatrest

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEncryptionAtRestAzureKeyVaultConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

