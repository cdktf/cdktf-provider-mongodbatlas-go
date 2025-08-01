// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudprovideraccesssetup


type CloudProviderAccessSetupAzureConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/cloud_provider_access_setup#atlas_azure_app_id CloudProviderAccessSetup#atlas_azure_app_id}.
	AtlasAzureAppId *string `field:"required" json:"atlasAzureAppId" yaml:"atlasAzureAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/cloud_provider_access_setup#service_principal_id CloudProviderAccessSetup#service_principal_id}.
	ServicePrincipalId *string `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.39.0/docs/resources/cloud_provider_access_setup#tenant_id CloudProviderAccessSetup#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

