// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupcompliancepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/backupcompliancepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.23.0/docs/resources/backup_compliance_policy mongodbatlas_backup_compliance_policy}.
type BackupCompliancePolicy interface {
	cdktf.TerraformResource
	AuthorizedEmail() *string
	SetAuthorizedEmail(val *string)
	AuthorizedEmailInput() *string
	AuthorizedUserFirstName() *string
	SetAuthorizedUserFirstName(val *string)
	AuthorizedUserFirstNameInput() *string
	AuthorizedUserLastName() *string
	SetAuthorizedUserLastName(val *string)
	AuthorizedUserLastNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyProtectionEnabled() interface{}
	SetCopyProtectionEnabled(val interface{})
	CopyProtectionEnabledInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionAtRestEnabled() interface{}
	SetEncryptionAtRestEnabled(val interface{})
	EncryptionAtRestEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnDemandPolicyItem() BackupCompliancePolicyOnDemandPolicyItemOutputReference
	OnDemandPolicyItemInput() *BackupCompliancePolicyOnDemandPolicyItem
	PitEnabled() interface{}
	SetPitEnabled(val interface{})
	PitEnabledInput() interface{}
	PolicyItemDaily() BackupCompliancePolicyPolicyItemDailyOutputReference
	PolicyItemDailyInput() *BackupCompliancePolicyPolicyItemDaily
	PolicyItemHourly() BackupCompliancePolicyPolicyItemHourlyOutputReference
	PolicyItemHourlyInput() *BackupCompliancePolicyPolicyItemHourly
	PolicyItemMonthly() BackupCompliancePolicyPolicyItemMonthlyList
	PolicyItemMonthlyInput() interface{}
	PolicyItemWeekly() BackupCompliancePolicyPolicyItemWeeklyList
	PolicyItemWeeklyInput() interface{}
	PolicyItemYearly() BackupCompliancePolicyPolicyItemYearlyList
	PolicyItemYearlyInput() interface{}
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
	RestoreWindowDays() *float64
	SetRestoreWindowDays(val *float64)
	RestoreWindowDaysInput() *float64
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedDate() *string
	UpdatedUser() *string
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
	PutOnDemandPolicyItem(value *BackupCompliancePolicyOnDemandPolicyItem)
	PutPolicyItemDaily(value *BackupCompliancePolicyPolicyItemDaily)
	PutPolicyItemHourly(value *BackupCompliancePolicyPolicyItemHourly)
	PutPolicyItemMonthly(value interface{})
	PutPolicyItemWeekly(value interface{})
	PutPolicyItemYearly(value interface{})
	ResetCopyProtectionEnabled()
	ResetEncryptionAtRestEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPitEnabled()
	ResetPolicyItemDaily()
	ResetPolicyItemHourly()
	ResetPolicyItemMonthly()
	ResetPolicyItemWeekly()
	ResetPolicyItemYearly()
	ResetRestoreWindowDays()
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

// The jsii proxy struct for BackupCompliancePolicy
type jsiiProxy_BackupCompliancePolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BackupCompliancePolicy) AuthorizedEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) AuthorizedEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) AuthorizedUserFirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedUserFirstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) AuthorizedUserFirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedUserFirstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) AuthorizedUserLastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedUserLastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) AuthorizedUserLastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizedUserLastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) CopyProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) CopyProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) EncryptionAtRestEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRestEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) EncryptionAtRestEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRestEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) OnDemandPolicyItem() BackupCompliancePolicyOnDemandPolicyItemOutputReference {
	var returns BackupCompliancePolicyOnDemandPolicyItemOutputReference
	_jsii_.Get(
		j,
		"onDemandPolicyItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) OnDemandPolicyItemInput() *BackupCompliancePolicyOnDemandPolicyItem {
	var returns *BackupCompliancePolicyOnDemandPolicyItem
	_jsii_.Get(
		j,
		"onDemandPolicyItemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemDaily() BackupCompliancePolicyPolicyItemDailyOutputReference {
	var returns BackupCompliancePolicyPolicyItemDailyOutputReference
	_jsii_.Get(
		j,
		"policyItemDaily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemDailyInput() *BackupCompliancePolicyPolicyItemDaily {
	var returns *BackupCompliancePolicyPolicyItemDaily
	_jsii_.Get(
		j,
		"policyItemDailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemHourly() BackupCompliancePolicyPolicyItemHourlyOutputReference {
	var returns BackupCompliancePolicyPolicyItemHourlyOutputReference
	_jsii_.Get(
		j,
		"policyItemHourly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemHourlyInput() *BackupCompliancePolicyPolicyItemHourly {
	var returns *BackupCompliancePolicyPolicyItemHourly
	_jsii_.Get(
		j,
		"policyItemHourlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemMonthly() BackupCompliancePolicyPolicyItemMonthlyList {
	var returns BackupCompliancePolicyPolicyItemMonthlyList
	_jsii_.Get(
		j,
		"policyItemMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemMonthlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyItemMonthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemWeekly() BackupCompliancePolicyPolicyItemWeeklyList {
	var returns BackupCompliancePolicyPolicyItemWeeklyList
	_jsii_.Get(
		j,
		"policyItemWeekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemWeeklyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyItemWeeklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemYearly() BackupCompliancePolicyPolicyItemYearlyList {
	var returns BackupCompliancePolicyPolicyItemYearlyList
	_jsii_.Get(
		j,
		"policyItemYearly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) PolicyItemYearlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyItemYearlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) RestoreWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restoreWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) RestoreWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restoreWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) UpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupCompliancePolicy) UpdatedUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedUser",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.23.0/docs/resources/backup_compliance_policy mongodbatlas_backup_compliance_policy} Resource.
func NewBackupCompliancePolicy(scope constructs.Construct, id *string, config *BackupCompliancePolicyConfig) BackupCompliancePolicy {
	_init_.Initialize()

	if err := validateNewBackupCompliancePolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupCompliancePolicy{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.23.0/docs/resources/backup_compliance_policy mongodbatlas_backup_compliance_policy} Resource.
func NewBackupCompliancePolicy_Override(b BackupCompliancePolicy, scope constructs.Construct, id *string, config *BackupCompliancePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetAuthorizedEmail(val *string) {
	if err := j.validateSetAuthorizedEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedEmail",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetAuthorizedUserFirstName(val *string) {
	if err := j.validateSetAuthorizedUserFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedUserFirstName",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetAuthorizedUserLastName(val *string) {
	if err := j.validateSetAuthorizedUserLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedUserLastName",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetCopyProtectionEnabled(val interface{}) {
	if err := j.validateSetCopyProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetEncryptionAtRestEnabled(val interface{}) {
	if err := j.validateSetEncryptionAtRestEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtRestEnabled",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetPitEnabled(val interface{}) {
	if err := j.validateSetPitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitEnabled",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BackupCompliancePolicy)SetRestoreWindowDays(val *float64) {
	if err := j.validateSetRestoreWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreWindowDays",
		val,
	)
}

// Generates CDKTF code for importing a BackupCompliancePolicy resource upon running "cdktf plan <stack-name>".
func BackupCompliancePolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBackupCompliancePolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
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
func BackupCompliancePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupCompliancePolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupCompliancePolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupCompliancePolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupCompliancePolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupCompliancePolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BackupCompliancePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.backupCompliancePolicy.BackupCompliancePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) PutOnDemandPolicyItem(value *BackupCompliancePolicyOnDemandPolicyItem) {
	if err := b.validatePutOnDemandPolicyItemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOnDemandPolicyItem",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) PutPolicyItemDaily(value *BackupCompliancePolicyPolicyItemDaily) {
	if err := b.validatePutPolicyItemDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPolicyItemDaily",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) PutPolicyItemHourly(value *BackupCompliancePolicyPolicyItemHourly) {
	if err := b.validatePutPolicyItemHourlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPolicyItemHourly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) PutPolicyItemMonthly(value interface{}) {
	if err := b.validatePutPolicyItemMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPolicyItemMonthly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) PutPolicyItemWeekly(value interface{}) {
	if err := b.validatePutPolicyItemWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPolicyItemWeekly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) PutPolicyItemYearly(value interface{}) {
	if err := b.validatePutPolicyItemYearlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPolicyItemYearly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetCopyProtectionEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetCopyProtectionEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetEncryptionAtRestEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetEncryptionAtRestEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetPitEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetPitEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetPolicyItemDaily() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicyItemDaily",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetPolicyItemHourly() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicyItemHourly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetPolicyItemMonthly() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicyItemMonthly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetPolicyItemWeekly() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicyItemWeekly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetPolicyItemYearly() {
	_jsii_.InvokeVoid(
		b,
		"resetPolicyItemYearly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) ResetRestoreWindowDays() {
	_jsii_.InvokeVoid(
		b,
		"resetRestoreWindowDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupCompliancePolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupCompliancePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

