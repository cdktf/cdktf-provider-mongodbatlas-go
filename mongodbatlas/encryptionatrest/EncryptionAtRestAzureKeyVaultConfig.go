// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest


type EncryptionAtRestAzureKeyVaultConfig struct {
	// Azure environment in which your account credentials reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#azure_environment EncryptionAtRest#azure_environment}
	AzureEnvironment *string `field:"optional" json:"azureEnvironment" yaml:"azureEnvironment"`
	// Unique 36-hexadecimal character string that identifies an Azure application associated with your Azure Active Directory tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#client_id EncryptionAtRest#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Flag that indicates whether someone enabled encryption at rest for the specified  project.
	//
	// To disable encryption at rest using customer key management and remove the configuration details, pass only this parameter with a value of `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#enabled EncryptionAtRest#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Web address with a unique key that identifies for your Azure Key Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#key_identifier EncryptionAtRest#key_identifier}
	KeyIdentifier *string `field:"optional" json:"keyIdentifier" yaml:"keyIdentifier"`
	// Unique string that identifies the Azure Key Vault that contains your key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#key_vault_name EncryptionAtRest#key_vault_name}
	KeyVaultName *string `field:"optional" json:"keyVaultName" yaml:"keyVaultName"`
	// Enable connection to your Azure Key Vault over private networking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#require_private_networking EncryptionAtRest#require_private_networking}
	RequirePrivateNetworking interface{} `field:"optional" json:"requirePrivateNetworking" yaml:"requirePrivateNetworking"`
	// Name of the Azure resource group that contains your Azure Key Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#resource_group_name EncryptionAtRest#resource_group_name}
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Private data that you need secured and that belongs to the specified Azure Key Vault (AKV) tenant (**azureKeyVault.tenantID**). This data can include any type of sensitive data such as passwords, database connection strings, API keys, and the like. AKV stores this information as encrypted binary data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#secret EncryptionAtRest#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Unique 36-hexadecimal character string that identifies your Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#subscription_id EncryptionAtRest#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// Unique 36-hexadecimal character string that identifies the Azure Active Directory tenant within your Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/encryption_at_rest#tenant_id EncryptionAtRest#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

