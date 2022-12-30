package workflow

import (
	"context"
	"fmt"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskCreateMissingModuleAccountAssociations Task = "CreateMissingModuleAccountAssociations"
)

type CreateMissingModuleAccountAssociationsInput struct {
	ModulePropagationId                      string
	AccountsNeedingModuleAccountAssociations []models.OrganizationalAccount
	ActiveModuleAccountAssociations          []models.ModuleAccountAssociation
}

type CreateMissingModuleAccountAssociationsOutput struct {
	ActiveModuleAccountAssociations []models.ModuleAccountAssociation
}

func (t *TaskHandler) CreateMissingModuleAccountAssociations(ctx context.Context, input CreateMissingModuleAccountAssociationsInput) (*CreateMissingModuleAccountAssociationsOutput, error) {
	for _, orgAccount := range input.AccountsNeedingModuleAccountAssociations {
		newModuleAccountAssociation := models.NewModuleAccountAssociation{
			ModulePropagationId: input.ModulePropagationId,
			OrgAccountId:        orgAccount.OrgAccountId,
			RemoteStateBucket:   t.remoteStateBucket,
			RemoteStateKey:      fmt.Sprintf("backends/%s/%s/state.tfstate", orgAccount.OrgAccountId, input.ModulePropagationId),
			RemoteStateRegion:   t.remoteStateRegion,
		}
		moduleAccountAssociation, err := t.apiClient.PutModuleAccountAssociation(ctx, &newModuleAccountAssociation)
		if err != nil {
			return nil, err
		}
		input.ActiveModuleAccountAssociations = append(input.ActiveModuleAccountAssociations, *moduleAccountAssociation)
	}
	return &CreateMissingModuleAccountAssociationsOutput{
		ActiveModuleAccountAssociations: input.ActiveModuleAccountAssociations,
	}, nil
}
