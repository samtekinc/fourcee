package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sheacloud/tfom/pkg/models"
)

type StateFileJSON struct {
	Version          int    `json:"version"`
	TerraformVersion string `json:"terraform_version"`
	Lineage          string `json:"lineage"`
	Resources        []struct {
		Module    *string `json:"module"`
		Mode      string  `json:"mode"`
		Type      string  `json:"type"`
		Name      string  `json:"name"`
		Provider  string  `json:"provider"`
		Instances []struct {
			SchemaVersion int                    `json:"schema_version"`
			Attributes    map[string]interface{} `json:"attributes"`
		} `json:"instances"`
	} `json:"resources"`
}

func (c *APIClient) GetStateFileVersions(ctx context.Context, stateBucket string, stateKey string, limit *int) ([]*models.StateVersion, error) {
	versions, err := c.s3Client.ListObjectVersions(ctx, &s3.ListObjectVersionsInput{
		Bucket: aws.String(stateBucket),
		Prefix: aws.String(stateKey),
	})
	if err != nil {
		return nil, err
	}

	var stateVersions []*models.StateVersion
	for _, version := range versions.Versions {
		stateVersions = append(stateVersions, &models.StateVersion{
			VersionID:    *version.VersionId,
			LastModified: *version.LastModified,
			IsCurrent:    version.IsLatest,
			Bucket:       stateBucket,
			Key:          stateKey,
		})
	}

	return stateVersions, nil
}

func (c *APIClient) GetStateFileVersion(ctx context.Context, stateBucket string, stateKey string, versionID string) (*models.StateFile, error) {
	obj, err := c.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket:    aws.String(stateBucket),
		Key:       aws.String(stateKey),
		VersionId: aws.String(versionID),
	})
	if err != nil {
		return nil, err
	}

	var stateFileJSON StateFileJSON
	err = json.NewDecoder(obj.Body).Decode(&stateFileJSON)
	if err != nil {
		return nil, err
	}

	var stateFile models.StateFile
	stateFile.VersionID = versionID
	for _, resource := range stateFileJSON.Resources {
		for i, instance := range resource.Instances {
			var namePrefix string
			if resource.Module != nil {
				namePrefix = *resource.Module + "."
			}
			stateFile.Resources = append(stateFile.Resources, models.StateResource{
				Type:       resource.Type,
				Name:       fmt.Sprintf("%s%s[%v]", namePrefix, resource.Name, i),
				ID:         instance.Attributes["id"].(string),
				Attributes: instance.Attributes,
			})
		}
	}

	return &stateFile, nil
}
