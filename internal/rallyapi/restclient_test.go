package rallyapi

import (
	"com.github.niedch/internal/conf"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/defect", r.URL.Path)
		assert.Equal(t, "zsessionid", r.Header.Get("zsessionid"))
		assert.Equal(t, "((FormattedID contains \"DE123\") OR (FormattedID contains \"DE124\"))", r.URL.Query().Get("query"))

		response := Response[Defect]{
			QueryResult: QueryResult[Defect]{
				Results: []Defect{
					{FormattedID: "DE123", Name: "Test Defect", Description: "Test Description"},
					{FormattedID: "DE124", Name: "Test Defect 2", Description: "Test Description 2"},
				},
				TotalResultCount: 2,
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := &conf.Config{
		CONNECTALL_RALLY_URL:     server.URL,
		CONNECTALL_WORKSPACE_ID:  "workspace-id",
		CONNECTALL_RALLY_API_KEY: "zsessionid",
	}
	client := NewRestClient(config)

	queryBuilder := NewQueryBuilder().WithFormattedIds([]string{"DE123", "DE124"})
	defects, err := client.FindDefects(context.Background(), *queryBuilder)

	assert.NoError(t, err)
	assert.Len(t, defects, 2)
	assert.Equal(t, "DE123", defects[0].FormattedID)
	assert.Equal(t, "Test Defect", defects[0].Name)
	assert.Equal(t, "Test Description", defects[0].Description)

	assert.Equal(t, "DE124", defects[1].FormattedID)
	assert.Equal(t, "Test Defect 2", defects[1].Name)
	assert.Equal(t, "Test Description 2", defects[1].Description)
}

func TestPutDefect(t *testing.T) {
	var receivedDefect map[string]Defect

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/defect/12345", r.URL.Path)
		assert.Equal(t, "zsessionid", r.Header.Get("zsessionid"))

		// Verify request body structure
		err := json.NewDecoder(r.Body).Decode(&receivedDefect)
		assert.NoError(t, err)

		// Verify the defect is wrapped correctly
		defect, ok := receivedDefect["defect"]
		assert.True(t, ok, "Expected 'defect' key in request body")
		assert.Equal(t, "DE123", defect.FormattedID)
		assert.Equal(t, "Updated Name", defect.Name)
		assert.Equal(t, "Updated Description", defect.Description)
		assert.Equal(t, "Updated Notes", defect.Notes)

		// Return success response with warnings
		response := UpdateResponse[Defect]{
			OperationResult: OperationResult[Defect]{
				Version: Version{
					RallyAPIMajor: "2",
					RallyAPIMinor: "0",
				},
				Errors:   []string{},
				Warnings: []string{"Ignored JSON element defect.c_DefectImpactedArea during processing of this request."},
				Object:   defect,
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := &conf.Config{
		CONNECTALL_RALLY_URL:     server.URL,
		CONNECTALL_WORKSPACE_ID:  "workspace-id",
		CONNECTALL_RALLY_API_KEY: "zsessionid",
	}
	client := NewRestClient(config)

	// Create defect to update
	defect := Defect{
		FormattedID: "DE123",
		Name:        "Updated Name",
		Description: "Updated Description",
		Notes:       "Updated Notes",
	}

	queryBuilder := NewQueryBuilder().WithObjectId("12345")
	err := client.PutDefect(context.Background(), *queryBuilder, defect)

	assert.NoError(t, err)
}

func TestPutDefect_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return error status code
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	config := &conf.Config{
		CONNECTALL_RALLY_URL:     server.URL,
		CONNECTALL_WORKSPACE_ID:  "workspace-id",
		CONNECTALL_RALLY_API_KEY: "zsessionid",
	}
	client := NewRestClient(config)

	defect := Defect{
		FormattedID: "DE123",
		Name:        "Test",
	}

	queryBuilder := NewQueryBuilder().WithObjectId("12345")
	err := client.PutDefect(context.Background(), *queryBuilder, defect)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Status Code Returned: 400")
}

func TestPutDefect_WithRallyErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return success status but with Rally errors in response
		response := UpdateResponse[Defect]{
			OperationResult: OperationResult[Defect]{
				Version: Version{
					RallyAPIMajor: "2",
					RallyAPIMinor: "0",
				},
				Errors:   []string{"Invalid field value for 'Name'"},
				Warnings: []string{},
				Object:   Defect{},
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	config := &conf.Config{
		CONNECTALL_RALLY_URL:     server.URL,
		CONNECTALL_WORKSPACE_ID:  "workspace-id",
		CONNECTALL_RALLY_API_KEY: "zsessionid",
	}
	client := NewRestClient(config)

	defect := Defect{
		FormattedID: "DE123",
		Name:        "Invalid Name",
	}

	queryBuilder := NewQueryBuilder().WithObjectId("12345")
	err := client.PutDefect(context.Background(), *queryBuilder, defect)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Rally Error")
	assert.Contains(t, err.Error(), "Invalid field value")
}
