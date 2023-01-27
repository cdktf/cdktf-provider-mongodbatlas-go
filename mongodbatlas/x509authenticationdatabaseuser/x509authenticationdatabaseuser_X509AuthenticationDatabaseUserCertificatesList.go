package x509authenticationdatabaseuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/cdktf/cdktf-provider-mongodbatlas-go/mongodbatlas/x509authenticationdatabaseuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type X509AuthenticationDatabaseUserCertificatesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) X509AuthenticationDatabaseUserCertificatesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for X509AuthenticationDatabaseUserCertificatesList
type jsiiProxy_X509AuthenticationDatabaseUserCertificatesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewX509AuthenticationDatabaseUserCertificatesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) X509AuthenticationDatabaseUserCertificatesList {
	_init_.Initialize()

	if err := validateNewX509AuthenticationDatabaseUserCertificatesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_X509AuthenticationDatabaseUserCertificatesList{}

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.x509AuthenticationDatabaseUser.X509AuthenticationDatabaseUserCertificatesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewX509AuthenticationDatabaseUserCertificatesList_Override(x X509AuthenticationDatabaseUserCertificatesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-mongodbatlas.x509AuthenticationDatabaseUser.X509AuthenticationDatabaseUserCertificatesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		x,
	)
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (x *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) Get(index *float64) X509AuthenticationDatabaseUserCertificatesOutputReference {
	if err := x.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns X509AuthenticationDatabaseUserCertificatesOutputReference

	_jsii_.Invoke(
		x,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := x.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		x,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_X509AuthenticationDatabaseUserCertificatesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

