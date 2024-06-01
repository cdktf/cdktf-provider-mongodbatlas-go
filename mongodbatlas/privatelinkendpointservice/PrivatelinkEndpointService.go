// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatelinkendpointservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/privatelinkendpointservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.2/docs/resources/privatelink_endpoint_service mongodbatlas_privatelink_endpoint_service}.
type PrivatelinkEndpointService interface {
	cdktf.TerraformResource
	AwsConnectionStatus() *string
	AzureStatus() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeleteRequested() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndpointGroupName() *string
	Endpoints() PrivatelinkEndpointServiceEndpointsList
	EndpointServiceId() *string
	SetEndpointServiceId(val *string)
	EndpointServiceIdInput() *string
	EndpointsInput() interface{}
	ErrorMessage() *string
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
	GcpStatus() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterfaceEndpointId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrivateEndpointConnectionName() *string
	PrivateEndpointIpAddress() *string
	SetPrivateEndpointIpAddress(val *string)
	PrivateEndpointIpAddressInput() *string
	PrivateEndpointResourceId() *string
	PrivateLinkId() *string
	SetPrivateLinkId(val *string)
	PrivateLinkIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PrivatelinkEndpointServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutEndpoints(value interface{})
	PutTimeouts(value *PrivatelinkEndpointServiceTimeouts)
	ResetEndpoints()
	ResetGcpProjectId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateEndpointIpAddress()
	ResetTimeouts()
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

// The jsii proxy struct for PrivatelinkEndpointService
type jsiiProxy_PrivatelinkEndpointService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PrivatelinkEndpointService) AwsConnectionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsConnectionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) AzureStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) DeleteRequested() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deleteRequested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) EndpointGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Endpoints() PrivatelinkEndpointServiceEndpointsList {
	var returns PrivatelinkEndpointServiceEndpointsList
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) EndpointServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) EndpointServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) EndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) GcpProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) GcpProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) GcpStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) InterfaceEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) PrivateEndpointConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointConnectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) PrivateEndpointIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) PrivateEndpointIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) PrivateEndpointResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) PrivateLinkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) PrivateLinkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) ProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) Timeouts() PrivatelinkEndpointServiceTimeoutsOutputReference {
	var returns PrivatelinkEndpointServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatelinkEndpointService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.2/docs/resources/privatelink_endpoint_service mongodbatlas_privatelink_endpoint_service} Resource.
func NewPrivatelinkEndpointService(scope constructs.Construct, id *string, config *PrivatelinkEndpointServiceConfig) PrivatelinkEndpointService {
	_init_.Initialize()

	if err := validateNewPrivatelinkEndpointServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatelinkEndpointService{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.16.2/docs/resources/privatelink_endpoint_service mongodbatlas_privatelink_endpoint_service} Resource.
func NewPrivatelinkEndpointService_Override(p PrivatelinkEndpointService, scope constructs.Construct, id *string, config *PrivatelinkEndpointServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetEndpointServiceId(val *string) {
	if err := j.validateSetEndpointServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointServiceId",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetGcpProjectId(val *string) {
	if err := j.validateSetGcpProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpProjectId",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetPrivateEndpointIpAddress(val *string) {
	if err := j.validateSetPrivateEndpointIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointIpAddress",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetPrivateLinkId(val *string) {
	if err := j.validateSetPrivateLinkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkId",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetProviderName(val *string) {
	if err := j.validateSetProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerName",
		val,
	)
}

func (j *jsiiProxy_PrivatelinkEndpointService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a PrivatelinkEndpointService resource upon running "cdktf plan <stack-name>".
func PrivatelinkEndpointService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePrivatelinkEndpointService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
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
func PrivatelinkEndpointService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivatelinkEndpointService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PrivatelinkEndpointService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivatelinkEndpointService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PrivatelinkEndpointService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrivatelinkEndpointService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PrivatelinkEndpointService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.privatelinkEndpointService.PrivatelinkEndpointService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) PutEndpoints(value interface{}) {
	if err := p.validatePutEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEndpoints",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) PutTimeouts(value *PrivatelinkEndpointServiceTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) ResetEndpoints() {
	_jsii_.InvokeVoid(
		p,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) ResetGcpProjectId() {
	_jsii_.InvokeVoid(
		p,
		"resetGcpProjectId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) ResetPrivateEndpointIpAddress() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateEndpointIpAddress",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatelinkEndpointService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatelinkEndpointService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

