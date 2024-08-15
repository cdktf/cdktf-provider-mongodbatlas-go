// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudprovideraccessauthorization


type CloudProviderAccessAuthorizationAws struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.0/docs/resources/cloud_provider_access_authorization#iam_assumed_role_arn CloudProviderAccessAuthorization#iam_assumed_role_arn}.
	IamAssumedRoleArn *string `field:"required" json:"iamAssumedRoleArn" yaml:"iamAssumedRoleArn"`
}

