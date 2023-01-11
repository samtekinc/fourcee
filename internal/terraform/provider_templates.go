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
	if t.Config.Alias == nil || *t.Config.Alias == "" {
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

type AzureProviderTemplate struct {
	SubscriptionId string
}

func (t *AzureProviderTemplate) GetProviderConfiguration() (string, error) {
	return fmt.Sprintf(`provider "azurerm" {
  subscription_id = "%s"
  features {}
}`, t.SubscriptionId), nil
}

type GCPProviderTemplate struct {
	Config    models.GcpProviderConfiguration
	ProjectId string
}

func (t *GCPProviderTemplate) GetProviderConfiguration() (string, error) {
	if t.Config.Alias == nil || *t.Config.Alias == "" {
		return fmt.Sprintf(`provider "google" {
			project = "%s"
			region = "%s"
		  }`, t.ProjectId, t.Config.Region), nil
	}
	return fmt.Sprintf(`provider "google" {
  alias = "%s"
  project = "%s"
  region = "%s"
}`, t.Config.Alias, t.ProjectId, t.Config.Region), nil
}
