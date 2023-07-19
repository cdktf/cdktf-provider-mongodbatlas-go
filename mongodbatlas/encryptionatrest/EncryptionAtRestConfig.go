package encryptionatrest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EncryptionAtRestConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#project_id EncryptionAtRest#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#aws_kms EncryptionAtRest#aws_kms}.
	AwsKms *map[string]*string `field:"optional" json:"awsKms" yaml:"awsKms"`
	// aws_kms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#aws_kms_config EncryptionAtRest#aws_kms_config}
	AwsKmsConfig *EncryptionAtRestAwsKmsConfig `field:"optional" json:"awsKmsConfig" yaml:"awsKmsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#azure_key_vault EncryptionAtRest#azure_key_vault}.
	AzureKeyVault *map[string]*string `field:"optional" json:"azureKeyVault" yaml:"azureKeyVault"`
	// azure_key_vault_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#azure_key_vault_config EncryptionAtRest#azure_key_vault_config}
	AzureKeyVaultConfig *EncryptionAtRestAzureKeyVaultConfig `field:"optional" json:"azureKeyVaultConfig" yaml:"azureKeyVaultConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#google_cloud_kms EncryptionAtRest#google_cloud_kms}.
	GoogleCloudKms *map[string]*string `field:"optional" json:"googleCloudKms" yaml:"googleCloudKms"`
	// google_cloud_kms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#google_cloud_kms_config EncryptionAtRest#google_cloud_kms_config}
	GoogleCloudKmsConfig *EncryptionAtRestGoogleCloudKmsConfig `field:"optional" json:"googleCloudKmsConfig" yaml:"googleCloudKmsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/encryption_at_rest#id EncryptionAtRest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

