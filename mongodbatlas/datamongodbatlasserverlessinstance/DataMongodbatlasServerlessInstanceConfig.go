package datamongodbatlasserverlessinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasServerlessInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/serverless_instance#name DataMongodbatlasServerlessInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/serverless_instance#project_id DataMongodbatlasServerlessInstance#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/serverless_instance#continuous_backup_enabled DataMongodbatlasServerlessInstance#continuous_backup_enabled}.
	ContinuousBackupEnabled interface{} `field:"optional" json:"continuousBackupEnabled" yaml:"continuousBackupEnabled"`
	// links block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/serverless_instance#links DataMongodbatlasServerlessInstance#links}
	Links interface{} `field:"optional" json:"links" yaml:"links"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/serverless_instance#state_name DataMongodbatlasServerlessInstance#state_name}.
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
}

