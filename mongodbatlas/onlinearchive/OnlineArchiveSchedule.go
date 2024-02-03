// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinearchive


type OnlineArchiveSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#type OnlineArchive#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#day_of_month OnlineArchive#day_of_month}.
	DayOfMonth *float64 `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#day_of_week OnlineArchive#day_of_week}.
	DayOfWeek *float64 `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#end_hour OnlineArchive#end_hour}.
	EndHour *float64 `field:"optional" json:"endHour" yaml:"endHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#end_minute OnlineArchive#end_minute}.
	EndMinute *float64 `field:"optional" json:"endMinute" yaml:"endMinute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#start_hour OnlineArchive#start_hour}.
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/online_archive#start_minute OnlineArchive#start_minute}.
	StartMinute *float64 `field:"optional" json:"startMinute" yaml:"startMinute"`
}

