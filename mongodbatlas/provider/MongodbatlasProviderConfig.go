package provider


type MongodbatlasProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#alias MongodbatlasProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#assume_role MongodbatlasProvider#assume_role}
	AssumeRole *MongodbatlasProviderAssumeRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#aws_access_key_id MongodbatlasProvider#aws_access_key_id}.
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#aws_secret_access_key MongodbatlasProvider#aws_secret_access_key}.
	AwsSecretAccessKey *string `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#aws_session_token MongodbatlasProvider#aws_session_token}.
	AwsSessionToken *string `field:"optional" json:"awsSessionToken" yaml:"awsSessionToken"`
	// MongoDB Atlas Base URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#base_url MongodbatlasProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// MongoDB Atlas Base URL default to gov.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#is_mongodbgov_cloud MongodbatlasProvider#is_mongodbgov_cloud}
	IsMongodbgovCloud interface{} `field:"optional" json:"isMongodbgovCloud" yaml:"isMongodbgovCloud"`
	// MongoDB Atlas Programmatic Private Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#private_key MongodbatlasProvider#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// MongoDB Atlas Programmatic Public Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#public_key MongodbatlasProvider#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// MongoDB Realm Base URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#realm_base_url MongodbatlasProvider#realm_base_url}
	RealmBaseUrl *string `field:"optional" json:"realmBaseUrl" yaml:"realmBaseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#region MongodbatlasProvider#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#secret_name MongodbatlasProvider#secret_name}.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs#sts_endpoint MongodbatlasProvider#sts_endpoint}.
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
}
