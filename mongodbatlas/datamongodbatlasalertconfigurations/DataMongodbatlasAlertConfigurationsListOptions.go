package datamongodbatlasalertconfigurations


type DataMongodbatlasAlertConfigurationsListOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/alert_configurations#include_count DataMongodbatlasAlertConfigurations#include_count}.
	IncludeCount interface{} `field:"optional" json:"includeCount" yaml:"includeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/alert_configurations#items_per_page DataMongodbatlasAlertConfigurations#items_per_page}.
	ItemsPerPage *float64 `field:"optional" json:"itemsPerPage" yaml:"itemsPerPage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/d/alert_configurations#page_num DataMongodbatlasAlertConfigurations#page_num}.
	PageNum *float64 `field:"optional" json:"pageNum" yaml:"pageNum"`
}

