package terraform

import (
	"fmt"

	"github.com/sheacloud/tfom/pkg/models"
)

type ProviderTemplate interface {
	GetProviderConfiguration() (string, error)
}

type AWSProviderTemplate struct {
	Config         models.AwsProviderConfiguration
	AssumeRoleName string
	AccountId      string
	SessionName    string
}

func (t *AWSProviderTemplate) GetProviderConfiguration() (string, error) {
	if t.Config.Alias == "" {
		return fmt.Sprintf(`provider "aws" {
			region = "%s"
			assume_role {
			  role_arn = "arn:aws:iam::%s:role/%s"
			  session_name = "%s"
			}
		  }`, t.Config.Region, t.AccountId, t.AssumeRoleName, t.SessionName), nil
	}
	return fmt.Sprintf(`provider "aws" {
  alias = "%s"
  region = "%s"
  assume_role {
	role_arn = "arn:aws:iam::%s:role/%s"
	session_name = "%s"
  }
}`, t.Config.Alias, t.Config.Region, t.AccountId, t.AssumeRoleName, t.SessionName), nil
}
