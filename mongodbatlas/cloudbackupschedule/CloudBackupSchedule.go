package cloudbackupschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v2/cloudbackupschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/cloud_backup_schedule mongodbatlas_cloud_backup_schedule}.
type CloudBackupSchedule interface {
	cdktf.TerraformResource
	AutoExportEnabled() interface{}
	SetAutoExportEnabled(val interface{})
	AutoExportEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopySettings() CloudBackupScheduleCopySettingsList
	CopySettingsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Export() CloudBackupScheduleExportOutputReference
	ExportInput() *CloudBackupScheduleExport
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
	IdPolicy() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NextSnapshot() *string
	// The tree node.
	Node() constructs.Node
	PolicyItemDaily() CloudBackupSchedulePolicyItemDailyOutputReference
	PolicyItemDailyInput() *CloudBackupSchedulePolicyItemDaily
	PolicyItemHourly() CloudBackupSchedulePolicyItemHourlyOutputReference
	PolicyItemHourlyInput() *CloudBackupSchedulePolicyItemHourly
	PolicyItemMonthly() CloudBackupSchedulePolicyItemMonthlyList
	PolicyItemMonthlyInput() interface{}
	PolicyItemWeekly() CloudBackupSchedulePolicyItemWeeklyList
	PolicyItemWeeklyInput() interface{}
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
	ReferenceHourOfDay() *float64
	SetReferenceHourOfDay(val *float64)
	ReferenceHourOfDayInput() *float64
	ReferenceMinuteOfHour() *float64
	SetReferenceMinuteOfHour(val *float64)
	ReferenceMinuteOfHourInput() *float64
	RestoreWindowDays() *float64
	SetRestoreWindowDays(val *float64)
	RestoreWindowDaysInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateSnapshots() interface{}
	SetUpdateSnapshots(val interface{})
	UpdateSnapshotsInput() interface{}
	UseOrgAndGroupNamesInExportPrefix() interface{}
	SetUseOrgAndGroupNamesInExportPrefix(val interface{})
	UseOrgAndGroupNamesInExportPrefixInput() interface{}
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
	PutCopySettings(value interface{})
	PutExport(value *CloudBackupScheduleExport)
	PutPolicyItemDaily(value *CloudBackupSchedulePolicyItemDaily)
	PutPolicyItemHourly(value *CloudBackupSchedulePolicyItemHourly)
	PutPolicyItemMonthly(value interface{})
	PutPolicyItemWeekly(value interface{})
	ResetAutoExportEnabled()
	ResetCopySettings()
	ResetExport()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyItemDaily()
	ResetPolicyItemHourly()
	ResetPolicyItemMonthly()
	ResetPolicyItemWeekly()
	ResetReferenceHourOfDay()
	ResetReferenceMinuteOfHour()
	ResetRestoreWindowDays()
	ResetUpdateSnapshots()
	ResetUseOrgAndGroupNamesInExportPrefix()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudBackupSchedule
type jsiiProxy_CloudBackupSchedule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudBackupSchedule) AutoExportEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoExportEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) AutoExportEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoExportEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) CopySettings() CloudBackupScheduleCopySettingsList {
	var returns CloudBackupScheduleCopySettingsList
	_jsii_.Get(
		j,
		"copySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) CopySettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Export() CloudBackupScheduleExportOutputReference {
	var returns CloudBackupScheduleExportOutputReference
	_jsii_.Get(
		j,
		"export",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ExportInput() *CloudBackupScheduleExport {
	var returns *CloudBackupScheduleExport
	_jsii_.Get(
		j,
		"exportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) IdPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) NextSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemDaily() CloudBackupSchedulePolicyItemDailyOutputReference {
	var returns CloudBackupSchedulePolicyItemDailyOutputReference
	_jsii_.Get(
		j,
		"policyItemDaily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemDailyInput() *CloudBackupSchedulePolicyItemDaily {
	var returns *CloudBackupSchedulePolicyItemDaily
	_jsii_.Get(
		j,
		"policyItemDailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemHourly() CloudBackupSchedulePolicyItemHourlyOutputReference {
	var returns CloudBackupSchedulePolicyItemHourlyOutputReference
	_jsii_.Get(
		j,
		"policyItemHourly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemHourlyInput() *CloudBackupSchedulePolicyItemHourly {
	var returns *CloudBackupSchedulePolicyItemHourly
	_jsii_.Get(
		j,
		"policyItemHourlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemMonthly() CloudBackupSchedulePolicyItemMonthlyList {
	var returns CloudBackupSchedulePolicyItemMonthlyList
	_jsii_.Get(
		j,
		"policyItemMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemMonthlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyItemMonthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemWeekly() CloudBackupSchedulePolicyItemWeeklyList {
	var returns CloudBackupSchedulePolicyItemWeeklyList
	_jsii_.Get(
		j,
		"policyItemWeekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) PolicyItemWeeklyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyItemWeeklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ReferenceHourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceHourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ReferenceHourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceHourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ReferenceMinuteOfHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceMinuteOfHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) ReferenceMinuteOfHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"referenceMinuteOfHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) RestoreWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restoreWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) RestoreWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restoreWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) UpdateSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) UpdateSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) UseOrgAndGroupNamesInExportPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOrgAndGroupNamesInExportPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudBackupSchedule) UseOrgAndGroupNamesInExportPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOrgAndGroupNamesInExportPrefixInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/cloud_backup_schedule mongodbatlas_cloud_backup_schedule} Resource.
func NewCloudBackupSchedule(scope constructs.Construct, id *string, config *CloudBackupScheduleConfig) CloudBackupSchedule {
	_init_.Initialize()

	if err := validateNewCloudBackupScheduleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudBackupSchedule{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.9.0/docs/resources/cloud_backup_schedule mongodbatlas_cloud_backup_schedule} Resource.
func NewCloudBackupSchedule_Override(c CloudBackupSchedule, scope constructs.Construct, id *string, config *CloudBackupScheduleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetAutoExportEnabled(val interface{}) {
	if err := j.validateSetAutoExportEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoExportEnabled",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetReferenceHourOfDay(val *float64) {
	if err := j.validateSetReferenceHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceHourOfDay",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetReferenceMinuteOfHour(val *float64) {
	if err := j.validateSetReferenceMinuteOfHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceMinuteOfHour",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetRestoreWindowDays(val *float64) {
	if err := j.validateSetRestoreWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreWindowDays",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetUpdateSnapshots(val interface{}) {
	if err := j.validateSetUpdateSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSnapshots",
		val,
	)
}

func (j *jsiiProxy_CloudBackupSchedule)SetUseOrgAndGroupNamesInExportPrefix(val interface{}) {
	if err := j.validateSetUseOrgAndGroupNamesInExportPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOrgAndGroupNamesInExportPrefix",
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
func CloudBackupSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudBackupSchedule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudBackupSchedule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudBackupSchedule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudBackupSchedule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudBackupSchedule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudBackupSchedule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.cloudBackupSchedule.CloudBackupSchedule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) PutCopySettings(value interface{}) {
	if err := c.validatePutCopySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCopySettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) PutExport(value *CloudBackupScheduleExport) {
	if err := c.validatePutExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExport",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) PutPolicyItemDaily(value *CloudBackupSchedulePolicyItemDaily) {
	if err := c.validatePutPolicyItemDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPolicyItemDaily",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) PutPolicyItemHourly(value *CloudBackupSchedulePolicyItemHourly) {
	if err := c.validatePutPolicyItemHourlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPolicyItemHourly",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) PutPolicyItemMonthly(value interface{}) {
	if err := c.validatePutPolicyItemMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPolicyItemMonthly",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) PutPolicyItemWeekly(value interface{}) {
	if err := c.validatePutPolicyItemWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPolicyItemWeekly",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetAutoExportEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoExportEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetCopySettings() {
	_jsii_.InvokeVoid(
		c,
		"resetCopySettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetExport() {
	_jsii_.InvokeVoid(
		c,
		"resetExport",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetPolicyItemDaily() {
	_jsii_.InvokeVoid(
		c,
		"resetPolicyItemDaily",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetPolicyItemHourly() {
	_jsii_.InvokeVoid(
		c,
		"resetPolicyItemHourly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetPolicyItemMonthly() {
	_jsii_.InvokeVoid(
		c,
		"resetPolicyItemMonthly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetPolicyItemWeekly() {
	_jsii_.InvokeVoid(
		c,
		"resetPolicyItemWeekly",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetReferenceHourOfDay() {
	_jsii_.InvokeVoid(
		c,
		"resetReferenceHourOfDay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetReferenceMinuteOfHour() {
	_jsii_.InvokeVoid(
		c,
		"resetReferenceMinuteOfHour",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetRestoreWindowDays() {
	_jsii_.InvokeVoid(
		c,
		"resetRestoreWindowDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetUpdateSnapshots() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdateSnapshots",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) ResetUseOrgAndGroupNamesInExportPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetUseOrgAndGroupNamesInExportPrefix",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudBackupSchedule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudBackupSchedule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

