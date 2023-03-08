package ldapconfiguration


type LdapConfigurationUserToDnMapping struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/ldap_configuration#ldap_query LdapConfiguration#ldap_query}.
	LdapQuery *string `field:"optional" json:"ldapQuery" yaml:"ldapQuery"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/ldap_configuration#match LdapConfiguration#match}.
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/ldap_configuration#substitution LdapConfiguration#substitution}.
	Substitution *string `field:"optional" json:"substitution" yaml:"substitution"`
}

