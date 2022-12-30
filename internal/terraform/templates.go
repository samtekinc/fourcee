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
	for _, awsProvider := range input.ModulePropagation.AwsProviderConfigurations {
		providers = append(providers, &AWSProviderTemplate{
			Config:         awsProvider,
			AssumeRoleName: input.OrgAccount.AssumeRoleName,
			AccountId:      input.OrgAccount.CloudIdentifier,
			SessionName:    fmt.Sprintf("tfom-%s", input.ModulePropagation.ModulePropagationId),
		})
	}
	templateInput := TemplateInput{
		BackendBucket:   input.ModuleAccountAssociation.RemoteStateBucket,
		BackendKey:      input.ModuleAccountAssociation.RemoteStateKey,
		BackendRegion:   input.ModuleAccountAssociation.RemoteStateRegion,
		Providers:       providers,
		ModuleName:      input.ModulePropagation.Name,
		ModuleSource:    input.ModuleVersion.RemoteSource,
		ModuleArguments: input.ModulePropagation.Arguments,
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
	BackendBucket   string
	BackendKey      string
	BackendRegion   string
	Providers       []ProviderTemplate
	ModuleName      string
	ModuleSource    string
	ModuleArguments []models.Argument
}

var moduleTemplateString = `
terraform {
  backend "s3" {
    bucket = "{{.BackendBucket}}"
    key    = "{{.BackendKey}}"
    region = "{{.BackendRegion}}"
  }
}

{{range $index, $element := .Providers}}{{$element.GetProviderConfiguration}}
{{end}}

module "{{.ModuleName}}" {
  source = "{{.ModuleSource}}"
{{range $index, $element := .ModuleArguments}}  {{$element.Name}} = {{$element.Value}}
{{end}}}
`

var moduleTemplate = template.Must(template.New("module").Parse(moduleTemplateString))
