package federatedsettingsorgrolemapping


type FederatedSettingsOrgRoleMappingRoleAssignments struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/federated_settings_org_role_mapping#group_id FederatedSettingsOrgRoleMapping#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/federated_settings_org_role_mapping#org_id FederatedSettingsOrgRoleMapping#org_id}.
	OrgId *string `field:"optional" json:"orgId" yaml:"orgId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/federated_settings_org_role_mapping#roles FederatedSettingsOrgRoleMapping#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

