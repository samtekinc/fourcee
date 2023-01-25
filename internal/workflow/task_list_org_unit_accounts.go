package workflow

import (
	"context"

	"github.com/sheacloud/tfom/pkg/models"
)

const (
	TaskListOrgUnitAccounts Task = "ListOrgUnitAccounts"
)

type ListOrgUnitAccountsInput struct {
	OrgUnit       models.OrgUnit
	CloudPlatform models.CloudPlatform
}

type ListOrgUnitAccountsOutput struct {
	OrgAccounts []models.OrgAccount
}

func (t *TaskHandler) ListOrgUnitAccounts(ctx context.Context, input ListOrgUnitAccountsInput) (*ListOrgUnitAccountsOutput, error) {
	// get accounts under OU
	accounts := []models.OrgAccount{}
	nextCursor := ""
	for {
		accountsPage, err := t.apiClient.GetOrgUnitMembershipsByOrgUnit(ctx, input.OrgUnit.OrgUnitId, 100, nextCursor)
		if err != nil {
			return nil, err
		}
		for _, accountDetails := range accountsPage.Items {
			account, err := t.apiClient.GetOrgAccount(ctx, accountDetails.OrgAccountID)
			if err != nil {
				return nil, err
			}
			if account.CloudPlatform == input.CloudPlatform {
				accounts = append(accounts, *account)
			}
		}
		if accountsPage.NextCursor == "" {
			break
		}
		nextCursor = accountsPage.NextCursor
	}
	return &ListOrgUnitAccountsOutput{
		OrgAccounts: accounts,
	}, nil
}
