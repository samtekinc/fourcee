package client

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyModuleVersionFilters(tx *gorm.DB, filters *models.ModuleVersionFilters) *gorm.DB {
	if filters != nil {
		if filters.NameContains != nil {
			tx = tx.Where("name LIKE ?", "%"+*filters.NameContains+"%")
		}
		if filters.RemoteSourceContains != nil {
			tx = tx.Where("remote_source LIKE ?", "%"+*filters.RemoteSourceContains+"%")
		}
		if filters.TerraformVersion != nil {
			tx = tx.Where("terraform_version = ?", *filters.TerraformVersion)
		}
	}
	return tx
}

func applyModuleVersionPreloads(tx *gorm.DB) *gorm.DB {
	return tx.Preload("Variables")
}

func (c *APIClient) GetModuleVersionsByIDs(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var moduleVersions []*models.ModuleVersion
	tx := applyModuleVersionPreloads(c.db)
	err := tx.Find(&moduleVersions, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	var keyToIndex = map[string]int{}
	for i := range keys {
		keyToIndex[keys[i].String()] = i
	}

	response := make([]*dataloader.Result, len(moduleVersions))
	for i := range moduleVersions {
		index := keyToIndex[idToString(moduleVersions[i].ID)]
		response[index] = &dataloader.Result{Data: moduleVersions[i], Error: nil}
	}
	return response
}

func (c *APIClient) GetModuleVersion(ctx context.Context, id uint) (*models.ModuleVersion, error) {
	var moduleVersion models.ModuleVersion
	tx := applyModuleVersionPreloads(c.db)
	err := tx.First(&moduleVersion, id).Error
	if err != nil {
		return nil, err
	}
	return &moduleVersion, nil
}

func (c *APIClient) GetModuleVersionBatched(ctx context.Context, id uint) (*models.ModuleVersion, error) {
	thunk := c.moduleVersionsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModuleVersion), nil
}

func (c *APIClient) GetModuleVersions(ctx context.Context, filters *models.ModuleVersionFilters, limit *int, offset *int) ([]*models.ModuleVersion, error) {
	var moduleVersions []*models.ModuleVersion
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleVersionFilters(tx, filters)
	tx = applyModuleVersionPreloads(tx)
	err := tx.Find(&moduleVersions).Error
	if err != nil {
		return nil, err
	}
	return moduleVersions, nil
}

func (c *APIClient) GetModuleVersionsForModuleGroup(ctx context.Context, moduleGroupId uint, filters *models.ModuleVersionFilters, limit *int, offset *int) ([]*models.ModuleVersion, error) {
	var moduleVersions []*models.ModuleVersion
	tx := applyPagination(c.db, limit, offset)
	tx = applyModuleVersionFilters(tx, filters)
	tx = applyModuleVersionPreloads(tx)
	err := tx.Model(&models.ModuleGroup{Model: gorm.Model{ID: moduleGroupId}}).Association("ModuleVersionsAssociation").Find(&moduleVersions)
	if err != nil {
		return nil, err
	}
	return moduleVersions, nil
}

func (c *APIClient) CreateModuleVersion(ctx context.Context, input *models.NewModuleVersion) (*models.ModuleVersion, error) {
	directory := c.versionInstallationDirectory + uuid.New().String()
	variables, err := GetVariablesFromModule(input.RemoteSource, directory)
	if err != nil {
		return nil, err
	}

	moduleVersion := models.ModuleVersion{
		Name:             input.Name,
		ModuleGroupID:    input.ModuleGroupID,
		RemoteSource:     input.RemoteSource,
		TerraformVersion: input.TerraformVersion,
		Variables:        variables,
	}
	err = c.db.Create(&moduleVersion).Error
	if err != nil {
		return nil, err
	}

	os.RemoveAll(directory)

	return &moduleVersion, nil
}

func (c *APIClient) DeleteModuleVersion(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModuleVersion{}, id).Error
}
