package datalake


type DataLakeDataProcessRegion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/data_lake#cloud_provider DataLake#cloud_provider}.
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/data_lake#region DataLake#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
}

