// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#description Organization#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#name Organization#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#org_owner_id Organization#org_owner_id}.
	OrgOwnerId *string `field:"required" json:"orgOwnerId" yaml:"orgOwnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#role_names Organization#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#api_access_list_required Organization#api_access_list_required}.
	ApiAccessListRequired interface{} `field:"optional" json:"apiAccessListRequired" yaml:"apiAccessListRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#federation_settings_id Organization#federation_settings_id}.
	FederationSettingsId *string `field:"optional" json:"federationSettingsId" yaml:"federationSettingsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#gen_ai_features_enabled Organization#gen_ai_features_enabled}.
	GenAiFeaturesEnabled interface{} `field:"optional" json:"genAiFeaturesEnabled" yaml:"genAiFeaturesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#id Organization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#multi_factor_auth_required Organization#multi_factor_auth_required}.
	MultiFactorAuthRequired interface{} `field:"optional" json:"multiFactorAuthRequired" yaml:"multiFactorAuthRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#restrict_employee_access Organization#restrict_employee_access}.
	RestrictEmployeeAccess interface{} `field:"optional" json:"restrictEmployeeAccess" yaml:"restrictEmployeeAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#security_contact Organization#security_contact}.
	SecurityContact *string `field:"optional" json:"securityContact" yaml:"securityContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.1/docs/resources/organization#skip_default_alerts_settings Organization#skip_default_alerts_settings}.
	SkipDefaultAlertsSettings interface{} `field:"optional" json:"skipDefaultAlertsSettings" yaml:"skipDefaultAlertsSettings"`
}

