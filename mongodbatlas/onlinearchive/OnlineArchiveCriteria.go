package onlinearchive


type OnlineArchiveCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/online_archive#type OnlineArchive#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/online_archive#date_field OnlineArchive#date_field}.
	DateField *string `field:"optional" json:"dateField" yaml:"dateField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/online_archive#date_format OnlineArchive#date_format}.
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/online_archive#expire_after_days OnlineArchive#expire_after_days}.
	ExpireAfterDays *float64 `field:"optional" json:"expireAfterDays" yaml:"expireAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/online_archive#query OnlineArchive#query}.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

