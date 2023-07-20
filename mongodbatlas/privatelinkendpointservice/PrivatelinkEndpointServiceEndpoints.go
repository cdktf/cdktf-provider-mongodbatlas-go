package privatelinkendpointservice


type PrivatelinkEndpointServiceEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/privatelink_endpoint_service#endpoint_name PrivatelinkEndpointService#endpoint_name}.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/privatelink_endpoint_service#ip_address PrivatelinkEndpointService#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

