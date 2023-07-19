package project


type ProjectApiKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/project#api_key_id Project#api_key_id}.
	ApiKeyId *string `field:"required" json:"apiKeyId" yaml:"apiKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/project#role_names Project#role_names}.
	RoleNames *[]*string `field:"required" json:"roleNames" yaml:"roleNames"`
}

