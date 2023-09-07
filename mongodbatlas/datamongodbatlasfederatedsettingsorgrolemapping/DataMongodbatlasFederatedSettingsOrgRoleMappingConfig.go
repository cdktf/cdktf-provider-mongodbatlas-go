// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasfederatedsettingsorgrolemapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasFederatedSettingsOrgRoleMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/data-sources/federated_settings_org_role_mapping#federation_settings_id DataMongodbatlasFederatedSettingsOrgRoleMapping#federation_settings_id}.
	FederationSettingsId *string `field:"required" json:"federationSettingsId" yaml:"federationSettingsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/data-sources/federated_settings_org_role_mapping#org_id DataMongodbatlasFederatedSettingsOrgRoleMapping#org_id}.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.1/docs/data-sources/federated_settings_org_role_mapping#role_mapping_id DataMongodbatlasFederatedSettingsOrgRoleMapping#role_mapping_id}.
	RoleMappingId *string `field:"required" json:"roleMappingId" yaml:"roleMappingId"`
}

