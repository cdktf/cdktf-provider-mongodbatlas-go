package datamongodbatlassearchindexes

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasSearchIndexesConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#cluster_name DataMongodbatlasSearchIndexes#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#collection_name DataMongodbatlasSearchIndexes#collection_name}.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#database DataMongodbatlasSearchIndexes#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#project_id DataMongodbatlasSearchIndexes#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#id DataMongodbatlasSearchIndexes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#items_per_page DataMongodbatlasSearchIndexes#items_per_page}.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/search_indexes#page_num DataMongodbatlasSearchIndexes#page_num}.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
}

