package networkcontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkContainerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#atlas_cidr_block NetworkContainer#atlas_cidr_block}.
	AtlasCidrBlock *string `field:"required" json:"atlasCidrBlock" yaml:"atlasCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#project_id NetworkContainer#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#id NetworkContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#provider_name NetworkContainer#provider_name}.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#region NetworkContainer#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#region_name NetworkContainer#region_name}.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/network_container#regions NetworkContainer#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

