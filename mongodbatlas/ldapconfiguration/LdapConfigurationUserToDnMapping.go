// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ldapconfiguration


type LdapConfigurationUserToDnMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/ldap_configuration#ldap_query LdapConfiguration#ldap_query}.
	LdapQuery *string `field:"optional" json:"ldapQuery" yaml:"ldapQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/ldap_configuration#match LdapConfiguration#match}.
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.21.0/docs/resources/ldap_configuration#substitution LdapConfiguration#substitution}.
	Substitution *string `field:"optional" json:"substitution" yaml:"substitution"`
}

