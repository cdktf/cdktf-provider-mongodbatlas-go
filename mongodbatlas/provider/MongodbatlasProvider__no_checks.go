// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MongodbatlasProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_MongodbatlasProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateMongodbatlasProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMongodbatlasProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateMongodbatlasProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_MongodbatlasProvider) validateSetAssumeRoleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MongodbatlasProvider) validateSetIsMongodbgovCloudParameters(val interface{}) error {
	return nil
}

func validateNewMongodbatlasProviderParameters(scope constructs.Construct, id *string, config *MongodbatlasProviderConfig) error {
	return nil
}

