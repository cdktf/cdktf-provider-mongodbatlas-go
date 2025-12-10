// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datamongodbatlasalertconfigurations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v8/datamongodbatlasalertconfigurations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference interface {
	cdktf.ComplexObject
	ApiToken() *string
	ChannelName() *string
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
	DatadogRegion() *string
	DelayMin() *float64
	EmailAddress() *string
	EmailEnabled() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	IntegrationId() *string
	InternalValue() *DataMongodbatlasAlertConfigurationsResultsNotification
	SetInternalValue(val *DataMongodbatlasAlertConfigurationsResultsNotification)
	IntervalMin() *float64
	MicrosoftTeamsWebhookUrl() *string
	MobileNumber() *string
	NotifierId() *string
	OpsGenieApiKey() *string
	OpsGenieRegion() *string
	Roles() *[]*string
	ServiceKey() *string
	SmsEnabled() cdktf.IResolvable
	TeamId() *string
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
	Username() *string
	VictorOpsApiKey() *string
	VictorOpsRoutingKey() *string
	WebhookSecret() *string
	WebhookUrl() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference
type jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) DatadogApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) DatadogRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) DelayMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) EmailEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emailEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) InternalValue() *DataMongodbatlasAlertConfigurationsResultsNotification {
	var returns *DataMongodbatlasAlertConfigurationsResultsNotification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) IntervalMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) MicrosoftTeamsWebhookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftTeamsWebhookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) MobileNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) NotifierId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifierId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) OpsGenieApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsGenieApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) OpsGenieRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"opsGenieRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) SmsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"smsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) TeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) TeamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) VictorOpsApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"victorOpsApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) VictorOpsRoutingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"victorOpsRoutingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) WebhookSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) WebhookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookUrl",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasAlertConfigurationsResultsNotificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasAlertConfigurationsResultsNotificationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAlertConfigurations.DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasAlertConfigurationsResultsNotificationOutputReference_Override(d DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAlertConfigurations.DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference)SetInternalValue(val *DataMongodbatlasAlertConfigurationsResultsNotification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAlertConfigurationsResultsNotificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

