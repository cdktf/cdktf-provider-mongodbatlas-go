package datalakepipeline


type DataLakePipelineSinkPartitionFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/data_lake_pipeline#field_name DataLakePipeline#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.2/docs/resources/data_lake_pipeline#order DataLakePipeline#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}
