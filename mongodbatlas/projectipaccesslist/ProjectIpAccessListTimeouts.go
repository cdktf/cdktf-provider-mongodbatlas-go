package projectipaccesslist


type ProjectIpAccessListTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project_ip_access_list#delete ProjectIpAccessList#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/project_ip_access_list#read ProjectIpAccessList#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

