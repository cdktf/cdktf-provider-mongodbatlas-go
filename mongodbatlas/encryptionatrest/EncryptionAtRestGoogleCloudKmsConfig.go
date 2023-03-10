package encryptionatrest


type EncryptionAtRestGoogleCloudKmsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/encryption_at_rest#enabled EncryptionAtRest#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/encryption_at_rest#key_version_resource_id EncryptionAtRest#key_version_resource_id}.
	KeyVersionResourceId *string `field:"optional" json:"keyVersionResourceId" yaml:"keyVersionResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/encryption_at_rest#service_account_key EncryptionAtRest#service_account_key}.
	ServiceAccountKey *string `field:"optional" json:"serviceAccountKey" yaml:"serviceAccountKey"`
}

