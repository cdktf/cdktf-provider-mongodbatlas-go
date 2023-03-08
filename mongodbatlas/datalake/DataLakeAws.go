package datalake


type DataLakeAws struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/data_lake#role_id DataLake#role_id}.
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/data_lake#test_s3_bucket DataLake#test_s3_bucket}.
	TestS3Bucket *string `field:"required" json:"testS3Bucket" yaml:"testS3Bucket"`
}

