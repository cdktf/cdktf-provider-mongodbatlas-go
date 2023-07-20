package projectapikey


type ProjectApiKeyProjectAssignment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/project_api_key#project_id ProjectApiKey#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/project_api_key#role_names ProjectApiKey#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
}
