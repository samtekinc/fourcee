package terraform

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/sheacloud/tfom/pkg/models"
)

type TerraformConfigurationInput struct {
	ModuleAssignment  *models.ModuleAssignment
	ModulePropagation *models.ModulePropagation
	ModuleVersion     *models.ModuleVersion
	OrgAccount        *models.OrgAccount
	LockTableName     string
}

func GetTerraformConfiguration(input *TerraformConfigurationInput) ([]byte, error) {
	providers := []ProviderTemplate{}

	var arguments []models.Argument
	var awsProviderConfigurations []models.AwsProviderConfiguration
	var gcpProviderConfigurations []models.GcpProviderConfiguration
	var moduleName string
	if input.ModulePropagation == nil {
		arguments = input.ModuleAssignment.Arguments
		awsProviderConfigurations = input.ModuleAssignment.AwsProviderConfigurations
		gcpProviderConfigurations = input.ModuleAssignment.GcpProviderConfigurations
		moduleName = input.ModuleAssignment.Name
	} else {
		arguments = input.ModulePropagation.Arguments
		awsProviderConfigurations = input.ModulePropagation.AwsProviderConfigurations
		gcpProviderConfigurations = input.ModulePropagation.GcpProviderConfigurations
		moduleName = input.ModulePropagation.Name
	}

	switch input.OrgAccount.CloudPlatform {
	case models.CloudPlatformAWS:
		for _, awsProvider := range awsProviderConfigurations {
			providers = append(providers, &AWSProviderTemplate{
				Config:         awsProvider,
				AssumeRoleName: input.OrgAccount.AssumeRoleName,
				AccountId:      input.OrgAccount.CloudIdentifier,
				SessionName:    fmt.Sprintf("tfom-%v", input.ModuleAssignment.ID),
			})
		}
	case models.CloudPlatformAzure:
		providers = append(providers, &AzureProviderTemplate{
			SubscriptionId: input.OrgAccount.CloudIdentifier,
		})
	case models.CloudPlatformGCP:
		for _, gcpProvider := range gcpProviderConfigurations {
			providers = append(providers, &GCPProviderTemplate{
				Config:    gcpProvider,
				ProjectId: input.OrgAccount.CloudIdentifier,
			})
		}
	default:
		return nil, fmt.Errorf("unknown cloud platform: %s", input.OrgAccount.CloudPlatform)
	}

	templateInput := TemplateInput{
		BackendBucket:         input.ModuleAssignment.RemoteStateBucket,
		BackendKey:            input.ModuleAssignment.RemoteStateKey,
		BackendRegion:         input.ModuleAssignment.RemoteStateRegion,
		LockTableName:         input.LockTableName,
		Providers:             providers,
		AccountMetadata:       append(input.OrgAccount.Metadata, input.OrgAccount.GetInternalMetadata()...),
		ModuleVersionMetadata: input.ModuleVersion.GetInternalMetadata(),
		ModuleName:            moduleName,
		ModuleSource:          input.ModuleVersion.RemoteSource,
		ModuleArguments:       arguments,
	}

	if input.ModulePropagation != nil {
		templateInput.ModulePropagationMetadata = input.ModulePropagation.GetInternalMetadata()
	}

	buf := bytes.NewBuffer([]byte{})
	err := moduleTemplate.Execute(buf, templateInput)
	if err != nil {
		return nil, err
	}

	configBytes := hclwrite.Format(buf.Bytes())
	return configBytes, nil
}

type TemplateInput struct {
	BackendBucket             string
	BackendKey                string
	BackendRegion             string
	LockTableName             string
	Providers                 []ProviderTemplate
	AccountMetadata           []models.Metadata
	ModuleVersionMetadata     []models.Metadata
	ModulePropagationMetadata []models.Metadata
	ModuleName                string
	ModuleSource              string
	ModuleArguments           []models.Argument
}

var moduleTemplateString = `
terraform {
  backend "s3" {
    bucket = "{{.BackendBucket}}"
    key    = "{{.BackendKey}}"
    region = "{{.BackendRegion}}"
	dynamodb_table = "{{.LockTableName}}"
  }
}

locals {
{{range $index, $element := .AccountMetadata}}  org_account_{{$element.Name}} = "{{$element.Value}}"
{{end}}{{range $index, $element := .ModuleVersionMetadata}}  module_version_{{$element.Name}} = "{{$element.Value}}"
{{end}}{{range $index, $element := .ModulePropagationMetadata}}  module_propagation_{{$element.Name}} = "{{$element.Value}}"
{{end}}}

{{range $index, $element := .Providers}}{{$element.GetProviderConfiguration}}
{{end}}

module "{{.ModuleName}}" {
  source = "{{.ModuleSource}}"
{{range $index, $element := .ModuleArguments}}  {{$element.Name}} = {{$element.Value}}
{{end}}}
`

var moduleTemplate = template.Must(template.New("module").Parse(moduleTemplateString))
