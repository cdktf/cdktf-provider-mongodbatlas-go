// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package federateddatabaseinstance

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FederatedDatabaseInstanceStorageDatabasesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFederatedDatabaseInstanceStorageDatabasesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

