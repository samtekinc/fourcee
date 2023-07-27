package activities

import (
	"context"

	"github.com/samtekinc/fourcee/pkg/models"
)

func (r *Activities) ListModulePropagationOrgAccounts(ctx context.Context, modulePropagationID uint) ([]*models.OrgAccount, error) {
	// get the module propagation
	modulePropagation, err := r.apiClient.GetModulePropagation(ctx, modulePropagationID)
	if err != nil {
		return nil, err
	}

	// get the module propagations org unit
	modulePropagationOrgUnit, err := r.apiClient.GetOrgUnit(ctx, modulePropagation.OrgUnitID)
	if err != nil {
		return nil, err
	}

	orgUnits := []*models.OrgUnit{modulePropagationOrgUnit}

	// get the module propagations org unit downsteream org units
	downstreamOrgUnits, err := r.apiClient.GetDownstreamOrgUnits(ctx, modulePropagationOrgUnit.ID, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	orgUnits = append(orgUnits, downstreamOrgUnits...)

	// get the org accounts for the org units
	orgAccounts := []*models.OrgAccount{}
	for _, orgUnit := range orgUnits {
		orgUnitAccounts, err := r.apiClient.GetOrgAccountsForOrgUnit(ctx, orgUnit.ID, nil, nil, nil)
		if err != nil {
			return nil, err
		}
		orgAccounts = append(orgAccounts, orgUnitAccounts...)
	}

	return orgAccounts, nil
}
