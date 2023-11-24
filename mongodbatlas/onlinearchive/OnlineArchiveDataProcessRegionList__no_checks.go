// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package onlinearchive

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OnlineArchiveDataProcessRegionList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OnlineArchiveDataProcessRegionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OnlineArchiveDataProcessRegionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OnlineArchiveDataProcessRegionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OnlineArchiveDataProcessRegionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OnlineArchiveDataProcessRegionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOnlineArchiveDataProcessRegionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

