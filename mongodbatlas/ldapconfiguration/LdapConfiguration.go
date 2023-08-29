// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ldapconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v5/ldapconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/ldap_configuration mongodbatlas_ldap_configuration}.
type LdapConfiguration interface {
	cdktf.TerraformResource
	AuthenticationEnabled() interface{}
	SetAuthenticationEnabled(val interface{})
	AuthenticationEnabledInput() interface{}
	AuthorizationEnabled() interface{}
	SetAuthorizationEnabled(val interface{})
	AuthorizationEnabledInput() interface{}
	AuthzQueryTemplate() *string
	SetAuthzQueryTemplate(val *string)
	AuthzQueryTemplateInput() *string
	BindPassword() *string
	SetBindPassword(val *string)
	BindPasswordInput() *string
	BindUsername() *string
	SetBindUsername(val *string)
	BindUsernameInput() *string
	CaCertificate() *string
	SetCaCertificate(val *string)
	CaCertificateInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
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
	UserToDnMapping() LdapConfigurationUserToDnMappingList
	UserToDnMappingInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutUserToDnMapping(value interface{})
	ResetAuthorizationEnabled()
	ResetAuthzQueryTemplate()
	ResetCaCertificate()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetUserToDnMapping()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LdapConfiguration
type jsiiProxy_LdapConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LdapConfiguration) AuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) AuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) AuthorizationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) AuthorizationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) AuthzQueryTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authzQueryTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) AuthzQueryTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authzQueryTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) BindPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) BindPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) BindUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) BindUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) CaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) CaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) UserToDnMapping() LdapConfigurationUserToDnMappingList {
	var returns LdapConfigurationUserToDnMappingList
	_jsii_.Get(
		j,
		"userToDnMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LdapConfiguration) UserToDnMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userToDnMappingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/ldap_configuration mongodbatlas_ldap_configuration} Resource.
func NewLdapConfiguration(scope constructs.Construct, id *string, config *LdapConfigurationConfig) LdapConfiguration {
	_init_.Initialize()

	if err := validateNewLdapConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LdapConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.ldapConfiguration.LdapConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/ldap_configuration mongodbatlas_ldap_configuration} Resource.
func NewLdapConfiguration_Override(l LdapConfiguration, scope constructs.Construct, id *string, config *LdapConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.ldapConfiguration.LdapConfiguration",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetAuthenticationEnabled(val interface{}) {
	if err := j.validateSetAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetAuthorizationEnabled(val interface{}) {
	if err := j.validateSetAuthorizationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEnabled",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetAuthzQueryTemplate(val *string) {
	if err := j.validateSetAuthzQueryTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authzQueryTemplate",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetBindPassword(val *string) {
	if err := j.validateSetBindPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindPassword",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetBindUsername(val *string) {
	if err := j.validateSetBindUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindUsername",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetCaCertificate(val *string) {
	if err := j.validateSetCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificate",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LdapConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
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
func LdapConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.ldapConfiguration.LdapConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LdapConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.ldapConfiguration.LdapConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LdapConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLdapConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.ldapConfiguration.LdapConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LdapConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.ldapConfiguration.LdapConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LdapConfiguration) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LdapConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LdapConfiguration) PutUserToDnMapping(value interface{}) {
	if err := l.validatePutUserToDnMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUserToDnMapping",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetAuthorizationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthorizationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetAuthzQueryTemplate() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthzQueryTemplate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetCaCertificate() {
	_jsii_.InvokeVoid(
		l,
		"resetCaCertificate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetPort() {
	_jsii_.InvokeVoid(
		l,
		"resetPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) ResetUserToDnMapping() {
	_jsii_.InvokeVoid(
		l,
		"resetUserToDnMapping",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LdapConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LdapConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

