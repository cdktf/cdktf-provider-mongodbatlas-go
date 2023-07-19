package onlinearchive


type OnlineArchivePartitionFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/online_archive#field_name OnlineArchive#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.10.1/docs/resources/online_archive#order OnlineArchive#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

