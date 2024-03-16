// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/networkpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs/resources/network_peering mongodbatlas_network_peering}.
type NetworkPeering interface {
	cdktf.TerraformResource
	AccepterRegionName() *string
	SetAccepterRegionName(val *string)
	AccepterRegionNameInput() *string
	AtlasCidrBlock() *string
	SetAtlasCidrBlock(val *string)
	AtlasCidrBlockInput() *string
	AtlasGcpProjectId() *string
	SetAtlasGcpProjectId(val *string)
	AtlasGcpProjectIdInput() *string
	AtlasId() *string
	AtlasVpcName() *string
	SetAtlasVpcName(val *string)
	AtlasVpcNameInput() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	AzureDirectoryId() *string
	SetAzureDirectoryId(val *string)
	AzureDirectoryIdInput() *string
	AzureSubscriptionId() *string
	SetAzureSubscriptionId(val *string)
	AzureSubscriptionIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionId() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerId() *string
	SetContainerId(val *string)
	ContainerIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ErrorMessage() *string
	ErrorState() *string
	ErrorStateName() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpProjectId() *string
	SetGcpProjectId(val *string)
	GcpProjectIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkName() *string
	SetNetworkName(val *string)
	NetworkNameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerId() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderName() *string
	SetProviderName(val *string)
	ProviderNameInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RouteTableCidrBlock() *string
	SetRouteTableCidrBlock(val *string)
	RouteTableCidrBlockInput() *string
	Status() *string
	StatusName() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VnetName() *string
	SetVnetName(val *string)
	VnetNameInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccepterRegionName()
	ResetAtlasCidrBlock()
	ResetAtlasGcpProjectId()
	ResetAtlasVpcName()
	ResetAwsAccountId()
	ResetAzureDirectoryId()
	ResetAzureSubscriptionId()
	ResetGcpProjectId()
	ResetId()
	ResetNetworkName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceGroupName()
	ResetRouteTableCidrBlock()
	ResetVnetName()
	ResetVpcId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NetworkPeering
type jsiiProxy_NetworkPeering struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkPeering) AccepterRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accepterRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AccepterRegionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accepterRegionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasGcpProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasGcpProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasGcpProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasGcpProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasVpcName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasVpcName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AtlasVpcNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atlasVpcNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AzureDirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureDirectoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AzureDirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureDirectoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AzureSubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSubscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) AzureSubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureSubscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ErrorState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ErrorStateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorStateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) GcpProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) GcpProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) NetworkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) NetworkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) PeerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) RouteTableCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) RouteTableCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) StatusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) VnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) VnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkPeering) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs/resources/network_peering mongodbatlas_network_peering} Resource.
func NewNetworkPeering(scope constructs.Construct, id *string, config *NetworkPeeringConfig) NetworkPeering {
	_init_.Initialize()

	if err := validateNewNetworkPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkPeering{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.15.2/docs/resources/network_peering mongodbatlas_network_peering} Resource.
func NewNetworkPeering_Override(n NetworkPeering, scope constructs.Construct, id *string, config *NetworkPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAccepterRegionName(val *string) {
	if err := j.validateSetAccepterRegionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accepterRegionName",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAtlasCidrBlock(val *string) {
	if err := j.validateSetAtlasCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atlasCidrBlock",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAtlasGcpProjectId(val *string) {
	if err := j.validateSetAtlasGcpProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atlasGcpProjectId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAtlasVpcName(val *string) {
	if err := j.validateSetAtlasVpcNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atlasVpcName",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAzureDirectoryId(val *string) {
	if err := j.validateSetAzureDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureDirectoryId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetAzureSubscriptionId(val *string) {
	if err := j.validateSetAzureSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureSubscriptionId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetContainerId(val *string) {
	if err := j.validateSetContainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetGcpProjectId(val *string) {
	if err := j.validateSetGcpProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpProjectId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetNetworkName(val *string) {
	if err := j.validateSetNetworkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkName",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetRouteTableCidrBlock(val *string) {
	if err := j.validateSetRouteTableCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeTableCidrBlock",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetVnetName(val *string) {
	if err := j.validateSetVnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetName",
		val,
	)
}

func (j *jsiiProxy_NetworkPeering)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTF code for importing a NetworkPeering resource upon running "cdktf plan <stack-name>".
func NetworkPeering_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkPeering_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NetworkPeering_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkPeering_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPeering_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkPeering_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkPeering_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.networkPeering.NetworkPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkPeering) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkPeering) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkPeering) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkPeering) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkPeering) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAccepterRegionName() {
	_jsii_.InvokeVoid(
		n,
		"resetAccepterRegionName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAtlasCidrBlock() {
	_jsii_.InvokeVoid(
		n,
		"resetAtlasCidrBlock",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAtlasGcpProjectId() {
	_jsii_.InvokeVoid(
		n,
		"resetAtlasGcpProjectId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAtlasVpcName() {
	_jsii_.InvokeVoid(
		n,
		"resetAtlasVpcName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		n,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAzureDirectoryId() {
	_jsii_.InvokeVoid(
		n,
		"resetAzureDirectoryId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetAzureSubscriptionId() {
	_jsii_.InvokeVoid(
		n,
		"resetAzureSubscriptionId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetGcpProjectId() {
	_jsii_.InvokeVoid(
		n,
		"resetGcpProjectId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetNetworkName() {
	_jsii_.InvokeVoid(
		n,
		"resetNetworkName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetResourceGroupName() {
	_jsii_.InvokeVoid(
		n,
		"resetResourceGroupName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetRouteTableCidrBlock() {
	_jsii_.InvokeVoid(
		n,
		"resetRouteTableCidrBlock",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetVnetName() {
	_jsii_.InvokeVoid(
		n,
		"resetVnetName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) ResetVpcId() {
	_jsii_.InvokeVoid(
		n,
		"resetVpcId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkPeering) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkPeering) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

