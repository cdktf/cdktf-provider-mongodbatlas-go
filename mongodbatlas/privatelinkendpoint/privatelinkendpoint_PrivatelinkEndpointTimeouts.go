package privatelinkendpoint


type PrivatelinkEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/privatelink_endpoint#create PrivatelinkEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/privatelink_endpoint#delete PrivatelinkEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

