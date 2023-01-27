package cloudprovideraccessauthorization


type CloudProviderAccessAuthorizationAws struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/cloud_provider_access_authorization#iam_assumed_role_arn CloudProviderAccessAuthorization#iam_assumed_role_arn}.
	IamAssumedRoleArn *string `field:"required" json:"iamAssumedRoleArn" yaml:"iamAssumedRoleArn"`
}

