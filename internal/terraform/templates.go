package terraform

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"text/template"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/sheacloud/tfom/pkg/models"
)

type TerraformConfigurationInput struct {
	ModuleAccountAssociation *models.ModuleAccountAssociation
	ModulePropagation        *models.ModulePropagation
	ModuleVersion            *models.ModuleVersion
	OrgAccount               *models.OrganizationalAccount
}

func GetTerraformConfigurationBase64(input *TerraformConfigurationInput) (string, error) {
	providers := []ProviderTemplate{}

	switch input.OrgAccount.CloudPlatform {
	case models.CloudPlatformAWS:
		for _, awsProvider := range input.ModulePropagation.AwsProviderConfigurations {
			providers = append(providers, &AWSProviderTemplate{
				Config:         awsProvider,
				AssumeRoleName: input.OrgAccount.AssumeRoleName,
				AccountId:      input.OrgAccount.CloudIdentifier,
				SessionName:    fmt.Sprintf("tfom-%s", input.ModulePropagation.ModulePropagationId),
			})
		}
	case models.CloudPlatformAzure:
		providers = append(providers, &AzureProviderTemplate{
			SubscriptionId: input.OrgAccount.CloudIdentifier,
		})
	case models.CloudPlatformGCP:
		for _, gcpProvider := range input.ModulePropagation.GcpProviderConfigurations {
			providers = append(providers, &GCPProviderTemplate{
				Config:    gcpProvider,
				ProjectId: input.OrgAccount.CloudIdentifier,
			})
		}
	default:
		return "", fmt.Errorf("unknown cloud platform: %s", input.OrgAccount.CloudPlatform)
	}

	templateInput := TemplateInput{
		BackendBucket:             input.ModuleAccountAssociation.RemoteStateBucket,
		BackendKey:                input.ModuleAccountAssociation.RemoteStateKey,
		BackendRegion:             input.ModuleAccountAssociation.RemoteStateRegion,
		Providers:                 providers,
		AccountMetadata:           append(input.OrgAccount.Metadata, input.OrgAccount.GetInternalMetadata()...),
		ModuleVersionMetadata:     input.ModuleVersion.GetInternalMetadata(),
		ModulePropagationMetadata: input.ModulePropagation.GetInternalMetadata(),
		ModuleName:                input.ModulePropagation.Name,
		ModuleSource:              input.ModuleVersion.RemoteSource,
		ModuleArguments:           input.ModulePropagation.Arguments,
	}

	buf := bytes.NewBuffer([]byte{})
	err := moduleTemplate.Execute(buf, templateInput)
	if err != nil {
		return "", err
	}

	configBytes := hclwrite.Format(buf.Bytes())
	return base64.StdEncoding.EncodeToString(configBytes), nil
}

type TemplateInput struct {
	BackendBucket             string
	BackendKey                string
	BackendRegion             string
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
