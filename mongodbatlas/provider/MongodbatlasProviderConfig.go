// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type MongodbatlasProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#alias MongodbatlasProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#assume_role MongodbatlasProvider#assume_role}
	AssumeRole interface{} `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// AWS API Access Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#aws_access_key_id MongodbatlasProvider#aws_access_key_id}
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// AWS API Access Secret Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#aws_secret_access_key MongodbatlasProvider#aws_secret_access_key}
	AwsSecretAccessKey *string `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// AWS Security Token Service provided session token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#aws_session_token MongodbatlasProvider#aws_session_token}
	AwsSessionToken *string `field:"optional" json:"awsSessionToken" yaml:"awsSessionToken"`
	// MongoDB Atlas Base URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#base_url MongodbatlasProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// MongoDB Atlas Base URL default to gov.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#is_mongodbgov_cloud MongodbatlasProvider#is_mongodbgov_cloud}
	IsMongodbgovCloud interface{} `field:"optional" json:"isMongodbgovCloud" yaml:"isMongodbgovCloud"`
	// MongoDB Atlas Programmatic Private Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#private_key MongodbatlasProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// MongoDB Atlas Programmatic Public Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#public_key MongodbatlasProvider#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// MongoDB Realm Base URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#realm_base_url MongodbatlasProvider#realm_base_url}
	RealmBaseUrl *string `field:"optional" json:"realmBaseUrl" yaml:"realmBaseUrl"`
	// Region where secret is stored as part of AWS Secret Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#region MongodbatlasProvider#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Name of secret stored in AWS Secret Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#secret_name MongodbatlasProvider#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// AWS Security Token Service endpoint. Required for cross-AWS region or cross-AWS account secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs#sts_endpoint MongodbatlasProvider#sts_endpoint}
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
}

