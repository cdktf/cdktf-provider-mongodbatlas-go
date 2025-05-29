// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ldapconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LdapConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#authentication_enabled LdapConfiguration#authentication_enabled}.
	AuthenticationEnabled interface{} `field:"required" json:"authenticationEnabled" yaml:"authenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#bind_password LdapConfiguration#bind_password}.
	BindPassword *string `field:"required" json:"bindPassword" yaml:"bindPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#bind_username LdapConfiguration#bind_username}.
	BindUsername *string `field:"required" json:"bindUsername" yaml:"bindUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#hostname LdapConfiguration#hostname}.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#project_id LdapConfiguration#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#authorization_enabled LdapConfiguration#authorization_enabled}.
	AuthorizationEnabled interface{} `field:"optional" json:"authorizationEnabled" yaml:"authorizationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#authz_query_template LdapConfiguration#authz_query_template}.
	AuthzQueryTemplate *string `field:"optional" json:"authzQueryTemplate" yaml:"authzQueryTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#ca_certificate LdapConfiguration#ca_certificate}.
	CaCertificate *string `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#id LdapConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#port LdapConfiguration#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// user_to_dn_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.35.0/docs/resources/ldap_configuration#user_to_dn_mapping LdapConfiguration#user_to_dn_mapping}
	UserToDnMapping interface{} `field:"optional" json:"userToDnMapping" yaml:"userToDnMapping"`
}

