// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudprovideraccessauthorization


type CloudProviderAccessAuthorizationAzure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/cloud_provider_access_authorization#atlas_azure_app_id CloudProviderAccessAuthorization#atlas_azure_app_id}.
	AtlasAzureAppId *string `field:"required" json:"atlasAzureAppId" yaml:"atlasAzureAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/cloud_provider_access_authorization#service_principal_id CloudProviderAccessAuthorization#service_principal_id}.
	ServicePrincipalId *string `field:"required" json:"servicePrincipalId" yaml:"servicePrincipalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/cloud_provider_access_authorization#tenant_id CloudProviderAccessAuthorization#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

