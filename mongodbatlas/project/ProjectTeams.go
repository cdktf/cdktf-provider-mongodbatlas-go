package project


type ProjectTeams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/project#role_names Project#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/project#team_id Project#team_id}.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

