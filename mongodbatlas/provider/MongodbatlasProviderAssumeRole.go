// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type MongodbatlasProviderAssumeRole struct {
	// The duration, between 15 minutes and 12 hours, of the role session.
	//
	// Valid time units are ns, us (or µs), ms, s, h, or m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#duration MongodbatlasProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// A unique identifier that might be required when you assume a role in another account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#external_id MongodbatlasProvider#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#policy MongodbatlasProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#policy_arns MongodbatlasProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#role_arn MongodbatlasProvider#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An identifier for the assumed role session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#session_name MongodbatlasProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// Source identity specified by the principal assuming the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#source_identity MongodbatlasProvider#source_identity}
	SourceIdentity *string `field:"optional" json:"sourceIdentity" yaml:"sourceIdentity"`
	// Assume role session tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#tags MongodbatlasProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Assume role session tag keys to pass to any subsequent sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs#transitive_tag_keys MongodbatlasProvider#transitive_tag_keys}
	TransitiveTagKeys *[]*string `field:"optional" json:"transitiveTagKeys" yaml:"transitiveTagKeys"`
}

