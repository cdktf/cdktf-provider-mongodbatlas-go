package encryptionatrest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/encryptionatrest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/encryption_at_rest mongodbatlas_encryption_at_rest}.
type EncryptionAtRest interface {
	cdktf.TerraformResource
	AwsKms() *map[string]*string
	SetAwsKms(val *map[string]*string)
	AwsKmsConfig() EncryptionAtRestAwsKmsConfigOutputReference
	AwsKmsConfigInput() *EncryptionAtRestAwsKmsConfig
	AwsKmsInput() *map[string]*string
	AzureKeyVault() *map[string]*string
	SetAzureKeyVault(val *map[string]*string)
	AzureKeyVaultConfig() EncryptionAtRestAzureKeyVaultConfigOutputReference
	AzureKeyVaultConfigInput() *EncryptionAtRestAzureKeyVaultConfig
	AzureKeyVaultInput() *map[string]*string
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
	GoogleCloudKms() *map[string]*string
	SetGoogleCloudKms(val *map[string]*string)
	GoogleCloudKmsConfig() EncryptionAtRestGoogleCloudKmsConfigOutputReference
	GoogleCloudKmsConfigInput() *EncryptionAtRestGoogleCloudKmsConfig
	GoogleCloudKmsInput() *map[string]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	PutAwsKmsConfig(value *EncryptionAtRestAwsKmsConfig)
	PutAzureKeyVaultConfig(value *EncryptionAtRestAzureKeyVaultConfig)
	PutGoogleCloudKmsConfig(value *EncryptionAtRestGoogleCloudKmsConfig)
	ResetAwsKms()
	ResetAwsKmsConfig()
	ResetAzureKeyVault()
	ResetAzureKeyVaultConfig()
	ResetGoogleCloudKms()
	ResetGoogleCloudKmsConfig()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EncryptionAtRest
type jsiiProxy_EncryptionAtRest struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EncryptionAtRest) AwsKms() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"awsKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AwsKmsConfig() EncryptionAtRestAwsKmsConfigOutputReference {
	var returns EncryptionAtRestAwsKmsConfigOutputReference
	_jsii_.Get(
		j,
		"awsKmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AwsKmsConfigInput() *EncryptionAtRestAwsKmsConfig {
	var returns *EncryptionAtRestAwsKmsConfig
	_jsii_.Get(
		j,
		"awsKmsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AwsKmsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"awsKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AzureKeyVault() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"azureKeyVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AzureKeyVaultConfig() EncryptionAtRestAzureKeyVaultConfigOutputReference {
	var returns EncryptionAtRestAzureKeyVaultConfigOutputReference
	_jsii_.Get(
		j,
		"azureKeyVaultConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AzureKeyVaultConfigInput() *EncryptionAtRestAzureKeyVaultConfig {
	var returns *EncryptionAtRestAzureKeyVaultConfig
	_jsii_.Get(
		j,
		"azureKeyVaultConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) AzureKeyVaultInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"azureKeyVaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) GoogleCloudKms() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleCloudKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) GoogleCloudKmsConfig() EncryptionAtRestGoogleCloudKmsConfigOutputReference {
	var returns EncryptionAtRestGoogleCloudKmsConfigOutputReference
	_jsii_.Get(
		j,
		"googleCloudKmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) GoogleCloudKmsConfigInput() *EncryptionAtRestGoogleCloudKmsConfig {
	var returns *EncryptionAtRestGoogleCloudKmsConfig
	_jsii_.Get(
		j,
		"googleCloudKmsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) GoogleCloudKmsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"googleCloudKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EncryptionAtRest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/encryption_at_rest mongodbatlas_encryption_at_rest} Resource.
func NewEncryptionAtRest(scope constructs.Construct, id *string, config *EncryptionAtRestConfig) EncryptionAtRest {
	_init_.Initialize()

	if err := validateNewEncryptionAtRestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EncryptionAtRest{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.11.0/docs/resources/encryption_at_rest mongodbatlas_encryption_at_rest} Resource.
func NewEncryptionAtRest_Override(e EncryptionAtRest, scope constructs.Construct, id *string, config *EncryptionAtRestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRest",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetAwsKms(val *map[string]*string) {
	if err := j.validateSetAwsKmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsKms",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetAzureKeyVault(val *map[string]*string) {
	if err := j.validateSetAzureKeyVaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureKeyVault",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetGoogleCloudKms(val *map[string]*string) {
	if err := j.validateSetGoogleCloudKmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleCloudKms",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EncryptionAtRest)SetProvisioners(val *[]interface{}) {
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
func EncryptionAtRest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEncryptionAtRest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EncryptionAtRest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEncryptionAtRest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EncryptionAtRest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEncryptionAtRest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EncryptionAtRest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.encryptionAtRest.EncryptionAtRest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EncryptionAtRest) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EncryptionAtRest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EncryptionAtRest) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EncryptionAtRest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EncryptionAtRest) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EncryptionAtRest) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EncryptionAtRest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EncryptionAtRest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EncryptionAtRest) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EncryptionAtRest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EncryptionAtRest) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRest) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EncryptionAtRest) PutAwsKmsConfig(value *EncryptionAtRestAwsKmsConfig) {
	if err := e.validatePutAwsKmsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAwsKmsConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EncryptionAtRest) PutAzureKeyVaultConfig(value *EncryptionAtRestAzureKeyVaultConfig) {
	if err := e.validatePutAzureKeyVaultConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAzureKeyVaultConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EncryptionAtRest) PutGoogleCloudKmsConfig(value *EncryptionAtRestGoogleCloudKmsConfig) {
	if err := e.validatePutGoogleCloudKmsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putGoogleCloudKmsConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetAwsKms() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsKms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetAwsKmsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsKmsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetAzureKeyVault() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureKeyVault",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetAzureKeyVaultConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureKeyVaultConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetGoogleCloudKms() {
	_jsii_.InvokeVoid(
		e,
		"resetGoogleCloudKms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetGoogleCloudKmsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetGoogleCloudKmsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EncryptionAtRest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EncryptionAtRest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

