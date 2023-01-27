package searchindex


type SearchIndexSynonyms struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/search_index#analyzer SearchIndex#analyzer}.
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/search_index#name SearchIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/search_index#source_collection SearchIndex#source_collection}.
	SourceCollection *string `field:"required" json:"sourceCollection" yaml:"sourceCollection"`
}

