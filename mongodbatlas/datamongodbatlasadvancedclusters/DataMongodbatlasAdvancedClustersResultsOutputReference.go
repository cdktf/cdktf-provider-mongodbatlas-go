package datamongodbatlasadvancedclusters

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v4/datamongodbatlasadvancedclusters/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataMongodbatlasAdvancedClustersResultsOutputReference interface {
	cdktf.ComplexObject
	AdvancedConfiguration() DataMongodbatlasAdvancedClustersResultsAdvancedConfigurationList
	BackupEnabled() cdktf.IResolvable
	BiConnectorConfig() DataMongodbatlasAdvancedClustersResultsBiConnectorConfigList
	ClusterType() *string
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
	ConnectionStrings() DataMongodbatlasAdvancedClustersResultsConnectionStringsList
	CreateDate() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSizeGb() *float64
	EncryptionAtRestProvider() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataMongodbatlasAdvancedClustersResults
	SetInternalValue(val *DataMongodbatlasAdvancedClustersResults)
	Labels() DataMongodbatlasAdvancedClustersResultsLabelsList
	MongoDbMajorVersion() *string
	MongoDbVersion() *string
	Name() *string
	Paused() cdktf.IResolvable
	PitEnabled() cdktf.IResolvable
	ReplicationSpecs() DataMongodbatlasAdvancedClustersResultsReplicationSpecsList
	RootCertType() *string
	StateName() *string
	TerminationProtectionEnabled() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VersionReleaseSystem() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataMongodbatlasAdvancedClustersResultsOutputReference
type jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) AdvancedConfiguration() DataMongodbatlasAdvancedClustersResultsAdvancedConfigurationList {
	var returns DataMongodbatlasAdvancedClustersResultsAdvancedConfigurationList
	_jsii_.Get(
		j,
		"advancedConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) BackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"backupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) BiConnectorConfig() DataMongodbatlasAdvancedClustersResultsBiConnectorConfigList {
	var returns DataMongodbatlasAdvancedClustersResultsBiConnectorConfigList
	_jsii_.Get(
		j,
		"biConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ConnectionStrings() DataMongodbatlasAdvancedClustersResultsConnectionStringsList {
	var returns DataMongodbatlasAdvancedClustersResultsConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) EncryptionAtRestProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAtRestProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) InternalValue() *DataMongodbatlasAdvancedClustersResults {
	var returns *DataMongodbatlasAdvancedClustersResults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) Labels() DataMongodbatlasAdvancedClustersResultsLabelsList {
	var returns DataMongodbatlasAdvancedClustersResultsLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) MongoDbMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbMajorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) MongoDbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) Paused() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) PitEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"pitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ReplicationSpecs() DataMongodbatlasAdvancedClustersResultsReplicationSpecsList {
	var returns DataMongodbatlasAdvancedClustersResultsReplicationSpecsList
	_jsii_.Get(
		j,
		"replicationSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) RootCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) TerminationProtectionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"terminationProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) VersionReleaseSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionReleaseSystem",
		&returns,
	)
	return returns
}


func NewDataMongodbatlasAdvancedClustersResultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataMongodbatlasAdvancedClustersResultsOutputReference {
	_init_.Initialize()

	if err := validateNewDataMongodbatlasAdvancedClustersResultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedClusters.DataMongodbatlasAdvancedClustersResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataMongodbatlasAdvancedClustersResultsOutputReference_Override(d DataMongodbatlasAdvancedClustersResultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.dataMongodbatlasAdvancedClusters.DataMongodbatlasAdvancedClustersResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference)SetInternalValue(val *DataMongodbatlasAdvancedClustersResults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataMongodbatlasAdvancedClustersResultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

