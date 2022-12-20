package database

import (
	"context"

	"github.com/sheacloud/tfom/internal/awsclients"
	"github.com/sheacloud/tfom/pkg/modules/models"
)

type ModulesDatabaseClientInterface interface {
	GetModuleGroup(ctx context.Context, moduleGroupId string) (*models.ModuleGroup, error)
	GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error)
	PutModuleGroup(ctx context.Context, input *models.ModuleGroup) error
	DeleteModuleGroup(ctx context.Context, moduleGroupId string) error

	GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error)
	GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error)
	PutModuleVersion(ctx context.Context, input *models.ModuleVersion) error
	DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error
}

type ModulesDatabaseClient struct {
	dynamodb          awsclients.DynamoDBInterface
	groupsTableName   string
	versionsTableName string
}

func NewModulesDatabaseClient(dynamodb awsclients.DynamoDBInterface, groupsTableName string, versionsTableName string) *ModulesDatabaseClient {
	return &ModulesDatabaseClient{
		dynamodb:          dynamodb,
		groupsTableName:   groupsTableName,
		versionsTableName: versionsTableName,
	}
}
