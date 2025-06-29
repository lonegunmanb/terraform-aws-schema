package schema

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ahmetb/go-linq/v3"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/iancoleman/strcase"
)

const tfProviderCode = `terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "us-east-1"
}

# Create a VPC
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
}
`

const registerTemplate = `package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/v6/generated/data"
	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/v6/generated/resource"
    "github.com/{{ .RepoOwner }}/{{ .GoModule }}/v6/generated/ephemeral"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema
var EphemeralResources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)
    ephemeralResources := make(map[string]*tfjson.Schema, 0)
	{{- range $name, $expression := .ResourceSchemas }}  
	resources["{{$name}}"] = {{$expression}}  
	{{- end }}  
	{{- range $name, $expression := .DataSourceSchemas }}  
	dataSources["{{$name}}"] = {{$expression}}  
	{{- end }}  
    {{- range $name, $expression := .EphemeralResourceSchemas }}
    ephemeralResources["{{$name}}"] = {{$expression}}
    {{- end }}
	Resources = resources
	DataSources = dataSources
    EphemeralResources = ephemeralResources
}`

const registerTestTemplate = `package generated_test

import (
	"testing"

	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/v6/generated"
	"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
	assert.NotEmpty(t, generated.Resources)
	assert.NotEmpty(t, generated.DataSources)
    assert.NotEmpty(t, generated.EphemeralResources)
}`

type RegisterParameter struct {
	ResourceSchemas          map[string]string
	DataSourceSchemas        map[string]string
	EphemeralResourceSchemas map[string]string
	RepoOwner                string
	GoModule                 string
}

type ProviderSchema struct {
	*tfjson.ProviderSchema
	Version *version.Version
}

func RefreshSchema(folder string) (version *version.Version, err error) {
	s, err := ExtractAzureRMProviderSchema()
	if err != nil {
		return nil, err
	}
	return s.Version, SaveProviderSchema(folder, s.ProviderSchema)
}

func ExtractAzureRMProviderSchema() (*ProviderSchema, error) {
	tmpFolder, err := os.MkdirTemp("", "*")
	if err != nil {
		return nil, fmt.Errorf("error creating temp TF code folder: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tmpFolder)
	}()

	err = os.WriteFile(filepath.Join(tmpFolder, "main.tf"), []byte(tfProviderCode), 0600)
	if err != nil {
		return nil, fmt.Errorf("error writing temp TF code file: %s", err)
	}

	execPath, teardown, err := terraformExecPath()
	if err != nil {
		return nil, err
	}
	if teardown != nil {
		defer teardown()
	}
	workingDir := tmpFolder
	tf, err := tfexec.NewTerraform(workingDir, *execPath)
	if err != nil {
		return nil, fmt.Errorf("error running NewTerraform: %w", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return nil, fmt.Errorf("error running Init: %s", err)
	}
	schema, err := tf.ProvidersSchema(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error running providers: %w", err)
	}
	r := schema.Schemas["registry.terraform.io/hashicorp/aws"]

	_, versions, err := tf.Version(context.Background(), true)
	if err != nil {
		return nil, fmt.Errorf("error running version: %w", err)
	}
	v, ok := versions["registry.terraform.io/hashicorp/aws"]
	if !ok {
		return nil, fmt.Errorf("error getting azurerm version")
	}

	return &ProviderSchema{
		ProviderSchema: r,
		Version:        v,
	}, nil
}

func terraformExecPath() (*string, func(), error) {
	var teardown func()
	execPath, err := findTerraformExecPath()
	if err != nil {
		return nil, nil, fmt.Errorf("error finding Terraform: %s", err)
	}
	if execPath == nil {
		installer := &releases.LatestVersion{
			Product:                  product.Terraform,
			IncludePrereleases:       false,
			SkipChecksumVerification: false,
		}
		tp, err := installer.Install(context.Background())
		if err != nil {
			return nil, nil, fmt.Errorf("error installing Terraform: %s", err)
		}
		teardown = func() {
			_ = os.Remove(tp)
		}
		execPath = &tp
	}
	return execPath, teardown, nil
}

func findTerraformExecPath() (*string, error) {
	terraformExecName := "terraform"
	if os.Getenv("GOOS") == "windows" {
		terraformExecName = "terraform.exe"
	}

	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return nil, fmt.Errorf("PATH environment variable not set")
	}

	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, p := range paths {
		execPath := filepath.Join(p, terraformExecName)
		if _, err := os.Stat(execPath); err == nil {
			return &execPath, nil
		}
	}

	return nil, nil
}

func SaveProviderSchema(folder string, s *tfjson.ProviderSchema) error {
	err := saveResourceSchemas(folder, s)
	if err != nil {
		return fmt.Errorf("error saving resource schemas: %w", err)
	}
	err = saveDataSourceSchemas(folder, s)
	if err != nil {
		return fmt.Errorf("error saving data source schemas: %w", err)
	}
	err = saveEphemeralResourceSchemas(folder, s)
	if err != nil {
		return fmt.Errorf("error saving ephemeral resource schemas: %w", err)
	}
	err = saveRegisterCode(folder, s)
	if err != nil {
		return fmt.Errorf("error saving register code: %w", err)
	}
	return nil
}

func saveRegisterCode(folder string, s *tfjson.ProviderSchema) error {
	parameter := buildRegisterParameter(s)
	err := saveRegister(registerTemplate, parameter, filepath.Join(folder, "register.go"))
	if err != nil {
		return fmt.Errorf("error saving register code: %w", err)
	}
	err = saveRegister(registerTestTemplate, parameter, filepath.Join(folder, "register_test.go"))
	if err != nil {
		return fmt.Errorf("error saving register test code: %w", err)
	}

	return nil
}

func buildRegisterParameter(s *tfjson.ProviderSchema) RegisterParameter {
	parameter := RegisterParameter{
		ResourceSchemas:          make(map[string]string, 0),
		DataSourceSchemas:        make(map[string]string, 0),
		EphemeralResourceSchemas: make(map[string]string, 0),
		RepoOwner:                "lonegunmanb",
		GoModule:                 "terraform-aws-schema",
	}
	linq.From(s.ResourceSchemas).OrderBy(byKey).ToMapBy(&parameter.ResourceSchemas, byKey, func(i interface{}) interface{} {
		pair := i.(linq.KeyValue)
		return fmt.Sprintf("resource.%sSchema()", strcase.ToCamel(pair.Key.(string)))
	})
	linq.From(s.DataSourceSchemas).OrderBy(byKey).ToMapBy(&parameter.DataSourceSchemas, byKey, func(i interface{}) interface{} {
		pair := i.(linq.KeyValue)
		return fmt.Sprintf("data.%sSchema()", strcase.ToCamel(pair.Key.(string)))
	})
	linq.From(s.EphemeralResourceSchemas).OrderBy(byKey).ToMapBy(&parameter.EphemeralResourceSchemas, byKey, func(i interface{}) interface{} {
		pair := i.(linq.KeyValue)
		return fmt.Sprintf("ephemeral.%sSchema()", strcase.ToCamel(pair.Key.(string)))
	})
	return parameter
}

func saveRegister(registerTemplate string, parameter RegisterParameter, destFilePath string) error {
	tmpl, err := template.New("register").Parse(registerTemplate)
	if err != nil {
		return err
	}
	buff := bytes.Buffer{}
	err = tmpl.Execute(&buff, parameter)
	if err != nil {
		return err
	}
	err = save(destFilePath, buff.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func saveDataSourceSchemas(folder string, s *tfjson.ProviderSchema) error {
	for dataName, schema := range s.DataSourceSchemas {
		err := SaveDataSourceSchema(dataName, folder, schema)
		if err != nil {
			return fmt.Errorf("error saving data source schema: %s", err)
		}
	}
	return nil
}

func saveResourceSchemas(folder string, s *tfjson.ProviderSchema) error {
	for resourceName, schema := range s.ResourceSchemas {
		err := SaveResourceSchema(resourceName, folder, schema)
		if err != nil {
			return fmt.Errorf("error saving resource schema: %s", err)
		}
	}
	return nil
}

func saveEphemeralResourceSchemas(folder string, s *tfjson.ProviderSchema) error {
	for resourceName, schema := range s.EphemeralResourceSchemas {
		err := SaveEphemeralResourceSchema(resourceName, folder, schema)
		if err != nil {
			return fmt.Errorf("error saving ephemeral resource schema: %s", err)
		}
	}
	return nil
}

func SaveDataSourceSchema(name, folder string, s *tfjson.Schema) error {
	return saveSchema(name, folder, s, PackageData)
}

func SaveResourceSchema(name, folder string, s *tfjson.Schema) error {
	return saveSchema(name, folder, s, PackageResource)
}

func SaveEphemeralResourceSchema(name, folder string, s *tfjson.Schema) error {
	return saveSchema(name, folder, s, PackageEphemeral)
}

func saveSchema(name, folder string, s *tfjson.Schema, pkg Package) error {
	jsonSchema, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("error marshalling schema: %s", err)
	}
	content, err := GenerateGoFileContent(name, string(jsonSchema), pkg)
	if err != nil {
		return fmt.Errorf("error generating go file content: %s", err)
	}
	fileName := strcase.ToLowerCamel(name)
	err = save(filepath.Join(folder, string(pkg), fmt.Sprintf("%s.go", fileName)), []byte(content))
	if err != nil {
		return fmt.Errorf("error saving file generated/%s/%s.go: %s", pkg, fileName, err)
	}
	content, err = GenerateGoTestFileContent(name, pkg)
	if err != nil {
		return fmt.Errorf("error generating go test file content: %w", err)
	}
	err = save(filepath.Join(folder, string(pkg), fmt.Sprintf("%s_test.go", fileName)), []byte(content))
	if err != nil {
		return fmt.Errorf("error saving file generated/%s/%s_test.go: %w", pkg, fileName, err)
	}
	return nil
}

func save(path string, bytes []byte) error {
	dir := filepath.Dir(path)
	// Create all folders in the path if they don't exist
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err == nil {
		// File exists, truncate it
		err = os.Truncate(path, 0)
		if err != nil {
			return err
		}
	}

	// Write JSON bytes to file
	err = os.WriteFile(path, bytes, 0600)
	if err != nil {
		return err
	}

	return nil
}

func byKey(i interface{}) interface{} {
	pair := i.(linq.KeyValue)
	return pair.Key
}
