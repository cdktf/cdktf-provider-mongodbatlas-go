// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cluster


type ClusterPinnedFcv struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.23.0/docs/resources/cluster#expiration_date Cluster#expiration_date}.
	ExpirationDate *string `field:"required" json:"expirationDate" yaml:"expirationDate"`
}

