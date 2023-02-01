package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/api/client"
	tfomConfig "github.com/sheacloud/tfom/internal/config"
	"github.com/sheacloud/tfom/pkg/models"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("unable to create logger, " + err.Error())
	}
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	conf := tfomConfig.ConfigFromEnv()

	apiClient, err := client.APIClientFromConfig(&conf, cfg)
	if err != nil {
		log.Fatalln("unable to create API Client", err)
	}

	orgDimensions, err := createOrgDimensions(ctx, apiClient, 5)
	if err != nil {
		log.Fatalln("unable to create org dimensions", err)
	}

	for i, orgDimension := range orgDimensions {
		orgUnits, err := createOrgUnits(ctx, apiClient, orgDimension.ID, (i+1)*50)
		if err != nil {
			log.Fatalln("unable to create org units", err)
		}
		fmt.Println(orgUnits)
	}
}

func createOrgDimensions(ctx context.Context, apiClient api.APIClientInterface, num int) ([]*models.OrgDimension, error) {
	orgDimensions := make([]*models.OrgDimension, num)
	for i := 0; i < num; i++ {
		orgDimension, err := apiClient.CreateOrgDimension(ctx, &models.NewOrgDimension{
			Name: fmt.Sprintf("Test Org Dimension %v", i),
		})
		if err != nil {
			return nil, err
		}
		orgDimensions[i] = orgDimension
	}
	return orgDimensions, nil
}

func createOrgUnits(ctx context.Context, apiClient api.APIClientInterface, orgDimensionID uint, num int) ([]*models.OrgUnit, error) {
	existingOrgUnits, err := apiClient.GetOrgUnitsForDimension(ctx, orgDimensionID, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	rootOrgUnit := existingOrgUnits[0]

	orgUnits := make([]*models.OrgUnit, num+1)
	orgUnits[0] = rootOrgUnit
	for i := 1; i <= num; i++ {
		parentOrgUnitIDIndex := rand.Intn(i)
		parentOrgUnitID := orgUnits[parentOrgUnitIDIndex].ID
		orgUnit, err := apiClient.CreateOrgUnit(ctx, &models.NewOrgUnit{
			Name:            fmt.Sprintf("Test Org Unit %v:%v", orgDimensionID, i),
			OrgDimensionID:  orgDimensionID,
			ParentOrgUnitID: parentOrgUnitID,
		})
		if err != nil {
			return nil, err
		}
		orgUnits[i] = orgUnit
	}
	return orgUnits, nil
}
