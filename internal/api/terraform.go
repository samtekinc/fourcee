package api

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/go-getter"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/sheacloud/tfom/pkg/models"
)

func attributeToValue(attr *hclwrite.Attribute) string {
	fullValue := string(attr.BuildTokens(nil).Bytes())
	valueSide := strings.SplitAfterN(fullValue, "=", 2)[1]
	valueSide = strings.TrimSpace(valueSide)
	return valueSide
}
func trimQuotes(s string) string {
	return strings.Trim(s, "\"")
}

func GetVariablesFromModule(moduleURL, workingDirectory string) ([]*models.ModuleVariable, error) {
	err := getter.GetAny(workingDirectory, moduleURL)
	if err != nil {
		return nil, err
	}

	variables := []*models.ModuleVariable{}

	files, err := os.ReadDir(workingDirectory)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range files {
		if strings.HasSuffix(fileInfo.Name(), ".tf") {
			fileData, err := os.ReadFile(workingDirectory + "/" + fileInfo.Name())
			if err != nil {
				return nil, err
			}
			hclFile, diags := hclwrite.ParseConfig(fileData, fileInfo.Name(), hcl.Pos{Line: 1, Column: 1})
			if diags.HasErrors() {
				panic(diags)
			}

			for _, block := range hclFile.Body().Blocks() {
				if block.Type() == "variable" {
					newVariable := &models.ModuleVariable{}
					labels := block.Labels()
					if len(labels) != 1 {
						panic("expected 1 label")
					}
					newVariable.Name = labels[0]

					for attrName, attr := range block.Body().Attributes() {
						switch attrName {
						case "type":
							newVariable.Type = attributeToValue(attr)
						case "default":
							newVariable.Default = aws.String(attributeToValue(attr))
						case "description":
							newVariable.Description = trimQuotes(attributeToValue(attr))
						}
					}

					variables = append(variables, newVariable)
				}
			}
		}
	}

	return variables, nil
}
