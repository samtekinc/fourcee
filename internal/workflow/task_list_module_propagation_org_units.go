package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskListModulePropagationOrgUnits Task = "ListModulePropagationOrgUnits"
)

type ListModulePropagationOrgUnitsInput struct {
	ModulePropagationId                 string
	ModulePropagationExecutionRequestId string
}

type ListModulePropagationOrgUnitsOutput struct {
	OrgUnits []models.OrganizationalUnit
}

func (t *TaskHandler) ListModulePropagationOrgUnits(ctx context.Context, input ListModulePropagationOrgUnitsInput) (*ListModulePropagationOrgUnitsOutput, error) {
	// get module propagation
	modulePropagation, err := t.apiClient.GetModulePropagation(ctx, input.ModulePropagationId)
	if err != nil {
		return nil, err
	}

	// get org unit from module propagation
	orgUnit, err := t.apiClient.GetOrganizationalUnit(ctx, modulePropagation.OrgDimensionId, modulePropagation.OrgUnitId)
	if err != nil {
		return nil, err
	}

	// get OUs under module propagation
	ouList := []models.OrganizationalUnit{}
	nextCursor := ""
	for {
		ouListPage, err := t.apiClient.GetOrganizationalUnitsByHierarchy(ctx, modulePropagation.OrgDimensionId, orgUnit.Hierarchy+orgUnit.OrgUnitId, 100, nextCursor)
		if err != nil {
			return nil, err
		}
		ouList = append(ouList, ouListPage.Items...)
		if ouListPage.NextCursor == "" {
			break
		}
		nextCursor = ouListPage.NextCursor
	}

	return &ListModulePropagationOrgUnitsOutput{
		OrgUnits: ouList,
	}, nil
}
