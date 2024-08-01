// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SearchIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#cluster_name SearchIndex#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#collection_name SearchIndex#collection_name}.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#database SearchIndex#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#name SearchIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#project_id SearchIndex#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#analyzer SearchIndex#analyzer}.
	Analyzer *string `field:"optional" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#analyzers SearchIndex#analyzers}.
	Analyzers *string `field:"optional" json:"analyzers" yaml:"analyzers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#fields SearchIndex#fields}.
	Fields *string `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#id SearchIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#mappings_dynamic SearchIndex#mappings_dynamic}.
	MappingsDynamic interface{} `field:"optional" json:"mappingsDynamic" yaml:"mappingsDynamic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#mappings_fields SearchIndex#mappings_fields}.
	MappingsFields *string `field:"optional" json:"mappingsFields" yaml:"mappingsFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#search_analyzer SearchIndex#search_analyzer}.
	SearchAnalyzer *string `field:"optional" json:"searchAnalyzer" yaml:"searchAnalyzer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#stored_source SearchIndex#stored_source}.
	StoredSource *string `field:"optional" json:"storedSource" yaml:"storedSource"`
	// synonyms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#synonyms SearchIndex#synonyms}
	Synonyms interface{} `field:"optional" json:"synonyms" yaml:"synonyms"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#timeouts SearchIndex#timeouts}
	Timeouts *SearchIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#type SearchIndex#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.17.5/docs/resources/search_index#wait_for_index_build_completion SearchIndex#wait_for_index_build_completion}.
	WaitForIndexBuildCompletion interface{} `field:"optional" json:"waitForIndexBuildCompletion" yaml:"waitForIndexBuildCompletion"`
}

