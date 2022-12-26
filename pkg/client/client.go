package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sheacloud/tfom/pkg/models"
)

type OrganizationsClientInterface interface {
	GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error)
}

type OrganizationsClient struct {
	baseUrl    string
	httpClient *http.Client
}

func NewOrganizationsClient(baseUrl string) *OrganizationsClient {
	return &OrganizationsClient{
		baseUrl:    baseUrl,
		httpClient: http.DefaultClient,
	}
}

func (c *OrganizationsClient) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req.WithContext(ctx))
}

func (c *OrganizationsClient) GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error) {
	values := url.Values{}
	values.Add("limit", strconv.Itoa(int(limit)))
	if cursor != "" {
		values.Add("cursor", cursor)
	}

	path := "/organizations/dimensions?" + values.Encode()
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var ods models.OrganizationalDimensions
	err = json.NewDecoder(resp.Body).Decode(&ods)
	if err != nil {
		return nil, err
	}

	return &ods, nil
}
