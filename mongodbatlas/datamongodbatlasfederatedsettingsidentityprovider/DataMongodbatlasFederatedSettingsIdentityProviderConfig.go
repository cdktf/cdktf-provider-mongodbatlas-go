// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasfederatedsettingsidentityprovider

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasFederatedSettingsIdentityProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/data-sources/federated_settings_identity_provider#federation_settings_id DataMongodbatlasFederatedSettingsIdentityProvider#federation_settings_id}.
	FederationSettingsId *string `field:"required" json:"federationSettingsId" yaml:"federationSettingsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/data-sources/federated_settings_identity_provider#identity_provider_id DataMongodbatlasFederatedSettingsIdentityProvider#identity_provider_id}.
	IdentityProviderId *string `field:"required" json:"identityProviderId" yaml:"identityProviderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/data-sources/federated_settings_identity_provider#id DataMongodbatlasFederatedSettingsIdentityProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

