package project


type ProjectApiKeys struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#api_key_id Project#api_key_id}.
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project#role_names Project#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
}

