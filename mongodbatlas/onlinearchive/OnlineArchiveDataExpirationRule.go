// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onlinearchive


type OnlineArchiveDataExpirationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.13.1/docs/resources/online_archive#expire_after_days OnlineArchive#expire_after_days}.
	ExpireAfterDays *float64 `field:"required" json:"expireAfterDays" yaml:"expireAfterDays"`
}

