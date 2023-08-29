// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datalake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataLakeStorageDatabasesCollectionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataLakeStorageDatabasesCollectionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataLakeStorageDatabasesCollectionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataLakeStorageDatabasesCollectionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataLakeStorageDatabasesCollectionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataLakeStorageDatabasesCollectionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

