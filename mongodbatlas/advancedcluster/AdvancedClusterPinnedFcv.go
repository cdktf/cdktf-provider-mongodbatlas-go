// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterPinnedFcv struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.29.0/docs/resources/advanced_cluster#expiration_date AdvancedCluster#expiration_date}.
	ExpirationDate *string `field:"required" json:"expirationDate" yaml:"expirationDate"`
}

