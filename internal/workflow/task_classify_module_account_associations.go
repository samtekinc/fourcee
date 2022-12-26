package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskClassifyModuleAccountAssociations Task = "ClassifyModuleAccountAssociations"
)

type orgAccountWrapper struct {
	OrgAccounts []models.OrganizationalAccount
}

type ClassifyModuleAccountAssociationsInput struct {
	OrgAccountsPerOrgUnit           []orgAccountWrapper
	ActiveModuleAccountAssociations []models.ModuleAccountAssociation
}

type ClassifyModuleAccountAssociationsOutput struct {
	ActiveModuleAccountAssociations          []models.ModuleAccountAssociation
	InactiveModuleAccountAssociations        []models.ModuleAccountAssociation
	AccountsNeedingModuleAccountAssociations []models.OrganizationalAccount
}

func (t *TaskHandler) ClassifyModuleAccountAssociations(ctx context.Context, input ClassifyModuleAccountAssociationsInput) (*ClassifyModuleAccountAssociationsOutput, error) {
	accountsWithModuleAccountAssociations := make(map[string]bool)
	accountsWithActiveModuleAccountAssociations := make(map[string]bool)
	// make note of which accounts have existing module account associations
	for _, moduleAccountAssociation := range input.ActiveModuleAccountAssociations {
		accountsWithModuleAccountAssociations[moduleAccountAssociation.OrgAccountId] = true
	}

	// make note of which accounts do not have existing module account associations
	accountsNeedingModuleAccountAssociations := make([]models.OrganizationalAccount, 0)
	for _, orgAccountWrapper := range input.OrgAccountsPerOrgUnit {
		for _, orgAccount := range orgAccountWrapper.OrgAccounts {
			accountsWithActiveModuleAccountAssociations[orgAccount.OrgAccountId] = true
			if !accountsWithModuleAccountAssociations[orgAccount.OrgAccountId] {
				accountsNeedingModuleAccountAssociations = append(accountsNeedingModuleAccountAssociations, orgAccount)
			}
		}
	}

	// filter out module account associations which don't have an account
	inactiveModuleAccountAssociations := make([]models.ModuleAccountAssociation, 0)
	for _, moduleAccountAssociation := range input.ActiveModuleAccountAssociations {
		if !accountsWithActiveModuleAccountAssociations[moduleAccountAssociation.OrgAccountId] {
			inactiveModuleAccountAssociations = append(inactiveModuleAccountAssociations, moduleAccountAssociation)
		}
	}

	filteredActiveModuleAccountAssociations := make([]models.ModuleAccountAssociation, 0)
	for _, moduleAccountAssociation := range input.ActiveModuleAccountAssociations {
		if accountsWithActiveModuleAccountAssociations[moduleAccountAssociation.OrgAccountId] {
			filteredActiveModuleAccountAssociations = append(filteredActiveModuleAccountAssociations, moduleAccountAssociation)
		}
	}

	return &ClassifyModuleAccountAssociationsOutput{
		ActiveModuleAccountAssociations:          filteredActiveModuleAccountAssociations,
		InactiveModuleAccountAssociations:        inactiveModuleAccountAssociations,
		AccountsNeedingModuleAccountAssociations: accountsNeedingModuleAccountAssociations,
	}, nil
}
