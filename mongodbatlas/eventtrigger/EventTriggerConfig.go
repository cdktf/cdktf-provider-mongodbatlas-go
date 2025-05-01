// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventtrigger

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventTriggerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#app_id EventTrigger#app_id}.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#name EventTrigger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#project_id EventTrigger#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#type EventTrigger#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_collection EventTrigger#config_collection}.
	ConfigCollection *string `field:"optional" json:"configCollection" yaml:"configCollection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_database EventTrigger#config_database}.
	ConfigDatabase *string `field:"optional" json:"configDatabase" yaml:"configDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_full_document EventTrigger#config_full_document}.
	ConfigFullDocument interface{} `field:"optional" json:"configFullDocument" yaml:"configFullDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_full_document_before EventTrigger#config_full_document_before}.
	ConfigFullDocumentBefore interface{} `field:"optional" json:"configFullDocumentBefore" yaml:"configFullDocumentBefore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_match EventTrigger#config_match}.
	ConfigMatch *string `field:"optional" json:"configMatch" yaml:"configMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_operation_type EventTrigger#config_operation_type}.
	ConfigOperationType *string `field:"optional" json:"configOperationType" yaml:"configOperationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_operation_types EventTrigger#config_operation_types}.
	ConfigOperationTypes *[]*string `field:"optional" json:"configOperationTypes" yaml:"configOperationTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_project EventTrigger#config_project}.
	ConfigProject *string `field:"optional" json:"configProject" yaml:"configProject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_providers EventTrigger#config_providers}.
	ConfigProviders *[]*string `field:"optional" json:"configProviders" yaml:"configProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_schedule EventTrigger#config_schedule}.
	ConfigSchedule *string `field:"optional" json:"configSchedule" yaml:"configSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#config_service_id EventTrigger#config_service_id}.
	ConfigServiceId *string `field:"optional" json:"configServiceId" yaml:"configServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#disabled EventTrigger#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// event_processors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#event_processors EventTrigger#event_processors}
	EventProcessors *EventTriggerEventProcessors `field:"optional" json:"eventProcessors" yaml:"eventProcessors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#function_id EventTrigger#function_id}.
	FunctionId *string `field:"optional" json:"functionId" yaml:"functionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#id EventTrigger#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.34.0/docs/resources/event_trigger#unordered EventTrigger#unordered}.
	Unordered interface{} `field:"optional" json:"unordered" yaml:"unordered"`
}

