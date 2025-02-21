// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/v7/searchindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.27.0/docs/resources/search_index mongodbatlas_search_index}.
type SearchIndex interface {
	cdktf.TerraformResource
	Analyzer() *string
	SetAnalyzer(val *string)
	AnalyzerInput() *string
	Analyzers() *string
	SetAnalyzers(val *string)
	AnalyzersInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	CollectionName() *string
	SetCollectionName(val *string)
	CollectionNameInput() *string
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Fields() *string
	SetFields(val *string)
	FieldsInput() *string
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
	IndexId() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MappingsDynamic() interface{}
	SetMappingsDynamic(val interface{})
	MappingsDynamicInput() interface{}
	MappingsFields() *string
	SetMappingsFields(val *string)
	MappingsFieldsInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	SearchAnalyzer() *string
	SetSearchAnalyzer(val *string)
	SearchAnalyzerInput() *string
	Status() *string
	StoredSource() *string
	SetStoredSource(val *string)
	StoredSourceInput() *string
	Synonyms() SearchIndexSynonymsList
	SynonymsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SearchIndexTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	WaitForIndexBuildCompletion() interface{}
	SetWaitForIndexBuildCompletion(val interface{})
	WaitForIndexBuildCompletionInput() interface{}
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
	PutSynonyms(value interface{})
	PutTimeouts(value *SearchIndexTimeouts)
	ResetAnalyzer()
	ResetAnalyzers()
	ResetFields()
	ResetId()
	ResetMappingsDynamic()
	ResetMappingsFields()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSearchAnalyzer()
	ResetStoredSource()
	ResetSynonyms()
	ResetTimeouts()
	ResetType()
	ResetWaitForIndexBuildCompletion()
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

// The jsii proxy struct for SearchIndex
type jsiiProxy_SearchIndex struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SearchIndex) Analyzer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) AnalyzerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Analyzers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) AnalyzersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analyzersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) CollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) CollectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Fields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) FieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) MappingsDynamic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mappingsDynamic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) MappingsDynamicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mappingsDynamicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) MappingsFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingsFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) MappingsFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingsFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) SearchAnalyzer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchAnalyzer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) SearchAnalyzerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchAnalyzerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) StoredSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storedSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) StoredSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storedSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Synonyms() SearchIndexSynonymsList {
	var returns SearchIndexSynonymsList
	_jsii_.Get(
		j,
		"synonyms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) SynonymsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"synonymsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Timeouts() SearchIndexTimeoutsOutputReference {
	var returns SearchIndexTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) WaitForIndexBuildCompletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForIndexBuildCompletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchIndex) WaitForIndexBuildCompletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForIndexBuildCompletionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.27.0/docs/resources/search_index mongodbatlas_search_index} Resource.
func NewSearchIndex(scope constructs.Construct, id *string, config *SearchIndexConfig) SearchIndex {
	_init_.Initialize()

	if err := validateNewSearchIndexParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SearchIndex{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/mongodb/mongodbatlas/1.27.0/docs/resources/search_index mongodbatlas_search_index} Resource.
func NewSearchIndex_Override(s SearchIndex, scope constructs.Construct, id *string, config *SearchIndexConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SearchIndex)SetAnalyzer(val *string) {
	if err := j.validateSetAnalyzerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyzer",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetAnalyzers(val *string) {
	if err := j.validateSetAnalyzersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyzers",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetCollectionName(val *string) {
	if err := j.validateSetCollectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionName",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetFields(val *string) {
	if err := j.validateSetFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetMappingsDynamic(val interface{}) {
	if err := j.validateSetMappingsDynamicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingsDynamic",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetMappingsFields(val *string) {
	if err := j.validateSetMappingsFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingsFields",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetSearchAnalyzer(val *string) {
	if err := j.validateSetSearchAnalyzerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchAnalyzer",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetStoredSource(val *string) {
	if err := j.validateSetStoredSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storedSource",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SearchIndex)SetWaitForIndexBuildCompletion(val interface{}) {
	if err := j.validateSetWaitForIndexBuildCompletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForIndexBuildCompletion",
		val,
	)
}

// Generates CDKTF code for importing a SearchIndex resource upon running "cdktf plan <stack-name>".
func SearchIndex_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSearchIndex_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
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
func SearchIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSearchIndex_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SearchIndex_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSearchIndex_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SearchIndex_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSearchIndex_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SearchIndex_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-mongodbatlas.searchIndex.SearchIndex",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SearchIndex) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SearchIndex) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SearchIndex) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SearchIndex) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SearchIndex) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SearchIndex) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SearchIndex) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SearchIndex) PutSynonyms(value interface{}) {
	if err := s.validatePutSynonymsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSynonyms",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SearchIndex) PutTimeouts(value *SearchIndexTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SearchIndex) ResetAnalyzer() {
	_jsii_.InvokeVoid(
		s,
		"resetAnalyzer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetAnalyzers() {
	_jsii_.InvokeVoid(
		s,
		"resetAnalyzers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetFields() {
	_jsii_.InvokeVoid(
		s,
		"resetFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetMappingsDynamic() {
	_jsii_.InvokeVoid(
		s,
		"resetMappingsDynamic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetMappingsFields() {
	_jsii_.InvokeVoid(
		s,
		"resetMappingsFields",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetSearchAnalyzer() {
	_jsii_.InvokeVoid(
		s,
		"resetSearchAnalyzer",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetStoredSource() {
	_jsii_.InvokeVoid(
		s,
		"resetStoredSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetSynonyms() {
	_jsii_.InvokeVoid(
		s,
		"resetSynonyms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) ResetWaitForIndexBuildCompletion() {
	_jsii_.InvokeVoid(
		s,
		"resetWaitForIndexBuildCompletion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SearchIndex) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchIndex) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

