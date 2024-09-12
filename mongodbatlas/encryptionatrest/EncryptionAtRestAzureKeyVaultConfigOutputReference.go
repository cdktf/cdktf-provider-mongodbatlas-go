// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package encryptionatrest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/encryptionatrest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EncryptionAtRestAzureKeyVaultConfigOutputReference interface {
	cdktf.ComplexObject
	AzureEnvironment() *string
	SetAzureEnvironment(val *string)
	AzureEnvironmentInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyIdentifier() *string
	SetKeyIdentifier(val *string)
	KeyIdentifierInput() *string
	KeyVaultName() *string
	SetKeyVaultName(val *string)
	KeyVaultNameInput() *string
	RequirePrivateNetworking() interface{}
	SetRequirePrivateNetworking(val interface{})
	RequirePrivateNetworkingInput() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Secret() *string
	SetSecret(val *string)
	SecretInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Valid() cdktf.IResolvable
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAzureEnvironment()
	ResetClientId()
	ResetEnabled()
	ResetKeyIdentifier()
	ResetKeyVaultName()
	ResetRequirePrivateNetworking()
	ResetResourceGroupName()
	ResetSecret()
	ResetSubscriptionId()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EncryptionAtRestAzureKeyVaultConfigOutputReference
type jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) AzureEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) AzureEnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureEnvironmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) KeyIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) KeyIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) KeyVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) KeyVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) RequirePrivateNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePrivateNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) RequirePrivateNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePrivateNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) SecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) Valid() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"valid",
		&returns,
	)
	return returns
}


func NewEncryptionAtRestAzureKeyVaultConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EncryptionAtRestAzureKeyVaultConfigOutputReference {
	_init_.Initialize()

	if err := validateNewEncryptionAtRestAzureKeyVaultConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRestAzureKeyVaultConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEncryptionAtRestAzureKeyVaultConfigOutputReference_Override(e EncryptionAtRestAzureKeyVaultConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRestAzureKeyVaultConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetAzureEnvironment(val *string) {
	if err := j.validateSetAzureEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureEnvironment",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetKeyIdentifier(val *string) {
	if err := j.validateSetKeyIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyIdentifier",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetKeyVaultName(val *string) {
	if err := j.validateSetKeyVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultName",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetRequirePrivateNetworking(val interface{}) {
	if err := j.validateSetRequirePrivateNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePrivateNetworking",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetAzureEnvironment() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureEnvironment",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		e,
		"resetClientId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetKeyIdentifier() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyIdentifier",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetKeyVaultName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyVaultName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetRequirePrivateNetworking() {
	_jsii_.InvokeVoid(
		e,
		"resetRequirePrivateNetworking",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetResourceGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		e,
		"resetTenantId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRestAzureKeyVaultConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

