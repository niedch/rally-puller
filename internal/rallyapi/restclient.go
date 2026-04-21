package rallyapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"com.github.niedch/internal/conf"
)

type RestClient struct {
	baseURL     string
	zSessionId  string
	projectId   string
	workspaceId string
}

func NewRestClient(config *conf.Config) *RestClient {
	return &RestClient{
		baseURL:     config.CONNECTALL_RALLY_URL,
		workspaceId: config.CONNECTALL_WORKSPACE_ID,
		zSessionId:  config.CONNECTALL_RALLY_API_KEY,
	}
}

func (r *RestClient) FindDefects(ctx context.Context, queryBuilder QueryBuilder) ([]Defect, error) {
	baseUrl := fmt.Sprintf("%s/%s", r.baseURL, "defect")
	queryParams, headers := r.createCommons(queryBuilder)

	log.Printf("Fetching from '%s'\n", baseUrl)
	response, err := get[Response[Defect]](ctx, baseUrl, queryParams, headers)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(response); err != nil {
		return nil, err
	}

	return response.QueryResult.Results, nil
}

func (r *RestClient) GetAttributeDefinition(ctx context.Context, queryBuilder QueryBuilder) ([]AllowedAttributeValue, error) {
	baseUrl := fmt.Sprintf("%s/attributedefinition/%s/allowedValues", r.baseURL, queryBuilder.objectId)
	queryParams, headers := r.createCommons(queryBuilder)

	response, err := get[Response[AllowedAttributeValue]](ctx, baseUrl, queryParams, headers)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(response); err != nil {
		return nil, err
	}

	return response.QueryResult.Results, nil
}

func (r *RestClient) PutDefect(ctx context.Context, queryBuilder QueryBuilder, defect Defect) error {
	baseUrl := fmt.Sprintf("%s/%s/%s", r.baseURL, "defect", queryBuilder.objectId)
	_, headers := r.createCommons(queryBuilder)

	log.Printf("Updating defect at '%s'\n", baseUrl)
	err := put(ctx, baseUrl, headers, defect)
	if err != nil {
		return err
	}

	return nil
}

func (r *RestClient) createCommons(queryBuilder QueryBuilder) (url.Values, http.Header) {
	queryParams := make(url.Values)
	headers := make(http.Header)

	headers.Add("zsessionid", r.zSessionId)
	queryParams.Add("workspace", fmt.Sprintf("workspace/%s", r.workspaceId))
	queryParams.Add("query", queryBuilder.build())
	return queryParams, headers
}

func checkResponse[T any](response Response[T]) error {
	if len(response.QueryResult.Errors) != 0 {
		return fmt.Errorf("Rally Error '%s'", response.QueryResult.Errors[0])
	}

	return nil
}
