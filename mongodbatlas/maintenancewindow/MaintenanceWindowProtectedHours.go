// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenancewindow


type MaintenanceWindowProtectedHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/maintenance_window#end_hour_of_day MaintenanceWindow#end_hour_of_day}.
	EndHourOfDay *float64 `field:"required" json:"endHourOfDay" yaml:"endHourOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.38.0/docs/resources/maintenance_window#start_hour_of_day MaintenanceWindow#start_hour_of_day}.
	StartHourOfDay *float64 `field:"required" json:"startHourOfDay" yaml:"startHourOfDay"`
}

