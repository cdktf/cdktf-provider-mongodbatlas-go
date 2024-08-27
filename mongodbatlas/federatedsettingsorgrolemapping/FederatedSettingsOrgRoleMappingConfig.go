// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package federatedsettingsorgrolemapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FederatedSettingsOrgRoleMappingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/federated_settings_org_role_mapping#external_group_name FederatedSettingsOrgRoleMapping#external_group_name}.
	ExternalGroupName *string `field:"required" json:"externalGroupName" yaml:"externalGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/federated_settings_org_role_mapping#federation_settings_id FederatedSettingsOrgRoleMapping#federation_settings_id}.
	FederationSettingsId *string `field:"required" json:"federationSettingsId" yaml:"federationSettingsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/federated_settings_org_role_mapping#org_id FederatedSettingsOrgRoleMapping#org_id}.
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// role_assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.18.1/docs/resources/federated_settings_org_role_mapping#role_assignments FederatedSettingsOrgRoleMapping#role_assignments}
	RoleAssignments interface{} `field:"required" json:"roleAssignments" yaml:"roleAssignments"`
}

