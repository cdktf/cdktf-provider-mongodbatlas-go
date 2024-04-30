// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/alertconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertConfigurationNotificationOutputReference interface {
	cdktf.ComplexObject
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	ChannelName() *string
	SetChannelName(val *string)
	ChannelNameInput() *string
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
	DatadogApiKey() *string
	SetDatadogApiKey(val *string)
	DatadogApiKeyInput() *string
	DatadogRegion() *string
	SetDatadogRegion(val *string)
	DatadogRegionInput() *string
	DelayMin() *float64
	SetDelayMin(val *float64)
	DelayMinInput() *float64
	EmailAddress() *string
	SetEmailAddress(val *string)
	EmailAddressInput() *string
	EmailEnabled() interface{}
	SetEmailEnabled(val interface{})
	EmailEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IntegrationId() *string
	SetIntegrationId(val *string)
	IntegrationIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IntervalMin() *float64
	SetIntervalMin(val *float64)
	IntervalMinInput() *float64
	MicrosoftTeamsWebhookUrl() *string
	SetMicrosoftTeamsWebhookUrl(val *string)
	MicrosoftTeamsWebhookUrlInput() *string
	MobileNumber() *string
	SetMobileNumber(val *string)
	MobileNumberInput() *string
	NotifierId() *string
	SetNotifierId(val *string)
	NotifierIdInput() *string
	OpsGenieApiKey() *string
	SetOpsGenieApiKey(val *string)
	OpsGenieApiKeyInput() *string
	OpsGenieRegion() *string
	SetOpsGenieRegion(val *string)
	OpsGenieRegionInput() *string
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	ServiceKey() *string
	SetServiceKey(val *string)
	ServiceKeyInput() *string
	SmsEnabled() interface{}
	SetSmsEnabled(val interface{})
	SmsEnabledInput() interface{}
	TeamId() *string
	SetTeamId(val *string)
	TeamIdInput() *string
	TeamName() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VictorOpsApiKey() *string
	SetVictorOpsApiKey(val *string)
	VictorOpsApiKeyInput() *string
	VictorOpsRoutingKey() *string
	SetVictorOpsRoutingKey(val *string)
	VictorOpsRoutingKeyInput() *string
	WebhookSecret() *string
	SetWebhookSecret(val *string)
	WebhookSecretInput() *string
	WebhookUrl() *string
	SetWebhookUrl(val *string)
	WebhookUrlInput() *string
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
	ResetApiToken()
	ResetChannelName()
	ResetDatadogApiKey()
	ResetDatadogRegion()
	ResetDelayMin()
	ResetEmailAddress()
	ResetEmailEnabled()
	ResetIntegrationId()
	ResetIntervalMin()
	ResetMicrosoftTeamsWebhookUrl()
	ResetMobileNumber()
	ResetNotifierId()
	ResetOpsGenieApiKey()
	ResetOpsGenieRegion()
	ResetRoles()
	ResetServiceKey()
	ResetSmsEnabled()
	ResetTeamId()
	ResetUsername()
	ResetVictorOpsApiKey()
	ResetVictorOpsRoutingKey()
	ResetWebhookSecret()
	ResetWebhookUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertConfigurationNotificationOutputReference
type jsiiProxy_AlertConfigurationNotificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) DatadogApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) DatadogApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) DatadogRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) DatadogRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) DelayMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) DelayMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) EmailEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) EmailEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"emailEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) IntegrationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) IntervalMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) IntervalMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) MicrosoftTeamsWebhookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftTeamsWebhookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) MicrosoftTeamsWebhookUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftTeamsWebhookUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) MobileNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) MobileNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) NotifierId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifierId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) NotifierIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifierIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) OpsGenieApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsGenieApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) OpsGenieApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsGenieApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) OpsGenieRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsGenieRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) OpsGenieRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsGenieRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) ServiceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) SmsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) SmsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TeamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) VictorOpsApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"victorOpsApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) VictorOpsApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"victorOpsApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) VictorOpsRoutingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"victorOpsRoutingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) VictorOpsRoutingKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"victorOpsRoutingKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) WebhookSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) WebhookSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) WebhookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference) WebhookUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookUrlInput",
		&returns,
	)
	return returns
}


func NewAlertConfigurationNotificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertConfigurationNotificationOutputReference {
	_init_.Initialize()

	if err := validateNewAlertConfigurationNotificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertConfigurationNotificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.alertConfiguration.AlertConfigurationNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertConfigurationNotificationOutputReference_Override(a AlertConfigurationNotificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.alertConfiguration.AlertConfigurationNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetApiToken(val *string) {
	if err := j.validateSetApiTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetDatadogApiKey(val *string) {
	if err := j.validateSetDatadogApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogApiKey",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetDatadogRegion(val *string) {
	if err := j.validateSetDatadogRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadogRegion",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetDelayMin(val *float64) {
	if err := j.validateSetDelayMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delayMin",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetEmailEnabled(val interface{}) {
	if err := j.validateSetEmailEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailEnabled",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetIntegrationId(val *string) {
	if err := j.validateSetIntegrationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationId",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetIntervalMin(val *float64) {
	if err := j.validateSetIntervalMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalMin",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetMicrosoftTeamsWebhookUrl(val *string) {
	if err := j.validateSetMicrosoftTeamsWebhookUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftTeamsWebhookUrl",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetMobileNumber(val *string) {
	if err := j.validateSetMobileNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNumber",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetNotifierId(val *string) {
	if err := j.validateSetNotifierIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifierId",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetOpsGenieApiKey(val *string) {
	if err := j.validateSetOpsGenieApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsGenieApiKey",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetOpsGenieRegion(val *string) {
	if err := j.validateSetOpsGenieRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"opsGenieRegion",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetRoles(val *[]*string) {
	if err := j.validateSetRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetServiceKey(val *string) {
	if err := j.validateSetServiceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKey",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetSmsEnabled(val interface{}) {
	if err := j.validateSetSmsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smsEnabled",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetTeamId(val *string) {
	if err := j.validateSetTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teamId",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetVictorOpsApiKey(val *string) {
	if err := j.validateSetVictorOpsApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"victorOpsApiKey",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetVictorOpsRoutingKey(val *string) {
	if err := j.validateSetVictorOpsRoutingKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"victorOpsRoutingKey",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetWebhookSecret(val *string) {
	if err := j.validateSetWebhookSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookSecret",
		val,
	)
}

func (j *jsiiProxy_AlertConfigurationNotificationOutputReference)SetWebhookUrl(val *string) {
	if err := j.validateSetWebhookUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookUrl",
		val,
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetApiToken() {
	_jsii_.InvokeVoid(
		a,
		"resetApiToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetChannelName() {
	_jsii_.InvokeVoid(
		a,
		"resetChannelName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetDatadogApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadogApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetDatadogRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadogRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetDelayMin() {
	_jsii_.InvokeVoid(
		a,
		"resetDelayMin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetEmailEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEmailEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetIntegrationId() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetIntervalMin() {
	_jsii_.InvokeVoid(
		a,
		"resetIntervalMin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetMicrosoftTeamsWebhookUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetMicrosoftTeamsWebhookUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetMobileNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetMobileNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetNotifierId() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifierId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetOpsGenieApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsGenieApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetOpsGenieRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetOpsGenieRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetRoles() {
	_jsii_.InvokeVoid(
		a,
		"resetRoles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetServiceKey() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetSmsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSmsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetTeamId() {
	_jsii_.InvokeVoid(
		a,
		"resetTeamId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetVictorOpsApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetVictorOpsApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetVictorOpsRoutingKey() {
	_jsii_.InvokeVoid(
		a,
		"resetVictorOpsRoutingKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetWebhookSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetWebhookSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ResetWebhookUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetWebhookUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertConfigurationNotificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

