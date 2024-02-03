// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advancedcluster


type AdvancedClusterReplicationSpecsRegionConfigsElectableSpecs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/advanced_cluster#instance_size AdvancedCluster#instance_size}.
	InstanceSize *string `field:"required" json:"instanceSize" yaml:"instanceSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/advanced_cluster#disk_iops AdvancedCluster#disk_iops}.
	DiskIops *float64 `field:"optional" json:"diskIops" yaml:"diskIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/advanced_cluster#ebs_volume_type AdvancedCluster#ebs_volume_type}.
	EbsVolumeType *string `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.0/docs/resources/advanced_cluster#node_count AdvancedCluster#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
}

