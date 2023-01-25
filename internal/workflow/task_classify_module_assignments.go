package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskClassifyModuleAssignments Task = "ClassifyModuleAssignments"
)

type orgAccountWrapper struct {
	OrgAccounts []models.OrgAccount
}

type ClassifyModuleAssignmentsInput struct {
	OrgAccountsPerOrgUnit   []orgAccountWrapper
	ActiveModuleAssignments []models.ModuleAssignment
}

type ClassifyModuleAssignmentsOutput struct {
	ActiveModuleAssignments          []models.ModuleAssignment
	InactiveModuleAssignments        []models.ModuleAssignment
	AccountsNeedingModuleAssignments []models.OrgAccount
}

func (t *TaskHandler) ClassifyModuleAssignments(ctx context.Context, input ClassifyModuleAssignmentsInput) (*ClassifyModuleAssignmentsOutput, error) {
	accountsWithModuleAssignments := make(map[string]bool)
	accountsWithActiveModuleAssignments := make(map[string]bool)
	// make note of which accounts have existing module assignments
	for _, moduleAssignment := range input.ActiveModuleAssignments {
		accountsWithModuleAssignments[moduleAssignment.OrgAccountID] = true
	}

	// make note of which accounts do not have existing module assignments
	accountsNeedingModuleAssignments := make([]models.OrgAccount, 0)
	for _, orgAccountWrapper := range input.OrgAccountsPerOrgUnit {
		for _, orgAccount := range orgAccountWrapper.OrgAccounts {
			accountsWithActiveModuleAssignments[orgAccount.OrgAccountID] = true
			if !accountsWithModuleAssignments[orgAccount.OrgAccountID] {
				accountsNeedingModuleAssignments = append(accountsNeedingModuleAssignments, orgAccount)
			}
		}
	}

	// filter out module assignments which don't have an account
	inactiveModuleAssignments := make([]models.ModuleAssignment, 0)
	for _, moduleAssignment := range input.ActiveModuleAssignments {
		if !accountsWithActiveModuleAssignments[moduleAssignment.OrgAccountID] {
			inactiveModuleAssignments = append(inactiveModuleAssignments, moduleAssignment)
		}
	}

	filteredActiveModuleAssignments := make([]models.ModuleAssignment, 0)
	for _, moduleAssignment := range input.ActiveModuleAssignments {
		if accountsWithActiveModuleAssignments[moduleAssignment.OrgAccountID] {
			filteredActiveModuleAssignments = append(filteredActiveModuleAssignments, moduleAssignment)
		}
	}

	return &ClassifyModuleAssignmentsOutput{
		ActiveModuleAssignments:          filteredActiveModuleAssignments,
		InactiveModuleAssignments:        inactiveModuleAssignments,
		AccountsNeedingModuleAssignments: accountsNeedingModuleAssignments,
	}, nil
}
