package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/services/modules/database"
	"github.com/sheacloud/tfom/pkg/modules/models"
)

type ModulesAPIClientInterface interface {
	GetModuleGroup(ctx context.Context, moduleGroupId string) (*models.ModuleGroup, error)
	GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error)
	PutModuleGroup(ctx context.Context, input *models.NewModuleGroup) (*models.ModuleGroup, error)
	DeleteModuleGroup(ctx context.Context, moduleGroupId string) error

	GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error)
	GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error)
	PutModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error)
	DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error
}

type ModulesAPIClient struct {
	dbClient database.ModulesDatabaseClientInterface
}

func NewModulesAPIClient(dbClient database.ModulesDatabaseClientInterface) *ModulesAPIClient {
	return &ModulesAPIClient{
		dbClient: dbClient,
	}
}
