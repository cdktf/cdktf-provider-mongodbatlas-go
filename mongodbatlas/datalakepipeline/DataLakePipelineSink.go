package datalakepipeline


type DataLakePipelineSink struct {
	// partition_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/data_lake_pipeline#partition_fields DataLakePipeline#partition_fields}
	PartitionFields interface{} `field:"optional" json:"partitionFields" yaml:"partitionFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/data_lake_pipeline#provider DataLakePipeline#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/data_lake_pipeline#region DataLakePipeline#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/data_lake_pipeline#type DataLakePipeline#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

