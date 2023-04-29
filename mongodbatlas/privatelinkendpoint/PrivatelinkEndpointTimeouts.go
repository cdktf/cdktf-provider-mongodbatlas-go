package privatelinkendpoint


type PrivatelinkEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/privatelink_endpoint#create PrivatelinkEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/privatelink_endpoint#delete PrivatelinkEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
