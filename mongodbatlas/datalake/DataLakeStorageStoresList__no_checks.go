// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datalake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataLakeStorageStoresList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataLakeStorageStoresList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataLakeStorageStoresList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataLakeStorageStoresList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataLakeStorageStoresList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataLakeStorageStoresListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

