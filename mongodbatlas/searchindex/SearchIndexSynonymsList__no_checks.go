// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package searchindex

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SearchIndexSynonymsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SearchIndexSynonymsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SearchIndexSynonymsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SearchIndexSynonymsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SearchIndexSynonymsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SearchIndexSynonymsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SearchIndexSynonymsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSearchIndexSynonymsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

