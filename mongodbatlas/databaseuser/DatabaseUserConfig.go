package databaseuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseUserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#project_id DatabaseUser#project_id}.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#roles DatabaseUser#roles}
	Roles interface{} `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#username DatabaseUser#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#auth_database_name DatabaseUser#auth_database_name}.
	AuthDatabaseName *string `field:"optional" json:"authDatabaseName" yaml:"authDatabaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#aws_iam_type DatabaseUser#aws_iam_type}.
	AwsIamType *string `field:"optional" json:"awsIamType" yaml:"awsIamType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#database_name DatabaseUser#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#id DatabaseUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#labels DatabaseUser#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#ldap_auth_type DatabaseUser#ldap_auth_type}.
	LdapAuthType *string `field:"optional" json:"ldapAuthType" yaml:"ldapAuthType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#password DatabaseUser#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// scopes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#scopes DatabaseUser#scopes}
	Scopes interface{} `field:"optional" json:"scopes" yaml:"scopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/mongodbatlas/r/database_user#x509_type DatabaseUser#x509_type}.
	X509Type *string `field:"optional" json:"x509Type" yaml:"x509Type"`
}

