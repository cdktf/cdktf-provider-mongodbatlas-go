package project


type ProjectTeams struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#role_names Project#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#team_id Project#team_id}.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

