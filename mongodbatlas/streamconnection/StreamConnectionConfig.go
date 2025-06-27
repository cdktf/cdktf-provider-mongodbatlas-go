// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#connection_name StreamConnection#connection_name}.
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#instance_name StreamConnection#instance_name}.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#project_id StreamConnection#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#type StreamConnection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#authentication StreamConnection#authentication}.
	Authentication *StreamConnectionAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#aws StreamConnection#aws}.
	Aws *StreamConnectionAws `field:"optional" json:"aws" yaml:"aws"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#bootstrap_servers StreamConnection#bootstrap_servers}.
	BootstrapServers *string `field:"optional" json:"bootstrapServers" yaml:"bootstrapServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#cluster_name StreamConnection#cluster_name}.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#cluster_project_id StreamConnection#cluster_project_id}.
	ClusterProjectId *string `field:"optional" json:"clusterProjectId" yaml:"clusterProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#config StreamConnection#config}.
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#db_role_to_execute StreamConnection#db_role_to_execute}.
	DbRoleToExecute *StreamConnectionDbRoleToExecute `field:"optional" json:"dbRoleToExecute" yaml:"dbRoleToExecute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#headers StreamConnection#headers}.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#networking StreamConnection#networking}.
	Networking *StreamConnectionNetworking `field:"optional" json:"networking" yaml:"networking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#security StreamConnection#security}.
	Security *StreamConnectionSecurity `field:"optional" json:"security" yaml:"security"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.37.0/docs/resources/stream_connection#url StreamConnection#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

