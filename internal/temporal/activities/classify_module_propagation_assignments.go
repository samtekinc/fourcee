package activities

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

type ModulePropagationAssignments struct {
	ActiveModuleAssignments          []*models.ModuleAssignment
	AccountsNeedingModuleAssignments []*models.OrgAccount
	InactiveModuleAssignments        []*models.ModuleAssignment
}

func (r *Activities) ClassifyModulePropagationAssignments(ctx context.Context, modulePropagationAccounts []*models.OrgAccount, existingModuleAssignments []*models.ModuleAssignment) (*ModulePropagationAssignments, error) {
	accountsWithModuleAssignments := make(map[uint]bool)
	activeAccounts := make(map[uint]bool)
	// make note of which accounts have existing module assignments
	for _, moduleAssignment := range existingModuleAssignments {
		accountsWithModuleAssignments[moduleAssignment.OrgAccountID] = true
	}

	// make note of which accounts do not have existing module assignments
	accountsNeedingModuleAssignments := make([]*models.OrgAccount, 0)
	for _, orgAccount := range modulePropagationAccounts {
		activeAccounts[orgAccount.ID] = true
		if !accountsWithModuleAssignments[orgAccount.ID] {
			accountsNeedingModuleAssignments = append(accountsNeedingModuleAssignments, orgAccount)
		}
	}

	// determine which of the existing module assignments need to be deleted vs. kept
	inactiveModuleAssignments := make([]*models.ModuleAssignment, 0)
	activeModuleAssignments := make([]*models.ModuleAssignment, 0)
	for _, moduleAssignment := range existingModuleAssignments {
		if !activeAccounts[moduleAssignment.OrgAccountID] {
			inactiveModuleAssignments = append(inactiveModuleAssignments, moduleAssignment)
		} else {
			activeModuleAssignments = append(activeModuleAssignments, moduleAssignment)
		}
	}

	return &ModulePropagationAssignments{
		ActiveModuleAssignments:          activeModuleAssignments,
		InactiveModuleAssignments:        inactiveModuleAssignments,
		AccountsNeedingModuleAssignments: accountsNeedingModuleAssignments,
	}, nil
}
