package rallyapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateResponse_Unmarshal(t *testing.T) {
	// Sample JSON response from Rally API (simplified version)
	jsonResponse := `{
		"OperationResult": {
			"_rallyAPIMajor": "2",
			"_rallyAPIMinor": "0",
			"Errors": [],
			"Warnings": ["Ignored JSON element defect.c_DefectImpactedArea during processing of this request."],
			"Object": {
				"_rallyAPIMajor": "2",
				"_rallyAPIMinor": "0",
				"_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/defect/838568308171",
				"_refObjectUUID": "ed22f7ab-eea3-4b25-b982-419f32ef8d55",
				"_objectVersion": "7",
				"_refObjectName": "New defect to analyse",
				"_type": "Defect",
				"ObjectID": 838568308171,
				"FormattedID": "DE201",
				"Name": "New defect to analyse",
				"Description": "This is a test description",
				"Notes": "Test notes"
			}
		}
	}`

	var response UpdateResponse[Defect]
	err := json.Unmarshal([]byte(jsonResponse), &response)

	assert.NoError(t, err)
	assert.Equal(t, "2", response.OperationResult.RallyAPIMajor)
	assert.Equal(t, "0", response.OperationResult.RallyAPIMinor)
	assert.Empty(t, response.OperationResult.Errors)
	assert.Len(t, response.OperationResult.Warnings, 1)
	assert.Contains(t, response.OperationResult.Warnings[0], "c_DefectImpactedArea")

	// Check the object
	defect := response.GetObject()
	assert.Equal(t, "DE201", defect.FormattedID)
	assert.Equal(t, "New defect to analyse", defect.Name)
	assert.Equal(t, "This is a test description", defect.Description)
	assert.Equal(t, "Test notes", defect.Notes)
}

func TestUpdateResponse_CheckErrors(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		response := UpdateResponse[Defect]{
			OperationResult: OperationResult[Defect]{
				Errors: []string{},
			},
		}

		err := CheckUpdateResponse(response)
		assert.NoError(t, err)
	})

	t.Run("with errors", func(t *testing.T) {
		response := UpdateResponse[Defect]{
			OperationResult: OperationResult[Defect]{
				Errors: []string{"Invalid field value", "Another error"},
			},
		}

		err := CheckUpdateResponse(response)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Invalid field value")
	})
}

func TestUpdateResponse_HasWarnings(t *testing.T) {
	t.Run("no warnings", func(t *testing.T) {
		response := UpdateResponse[Defect]{
			OperationResult: OperationResult[Defect]{
				Warnings: []string{},
			},
		}

		assert.False(t, response.HasWarnings())
	})

	t.Run("with warnings", func(t *testing.T) {
		response := UpdateResponse[Defect]{
			OperationResult: OperationResult[Defect]{
				Warnings: []string{"Warning message"},
			},
		}

		assert.True(t, response.HasWarnings())
		warnings := response.GetWarnings()
		assert.Len(t, warnings, 1)
		assert.Equal(t, "Warning message", warnings[0])
	})
}

func TestRequest_MarshalJSON(t *testing.T) {
	defect := Defect{
		FormattedID: "DE123",
		Name:        "Test Defect",
		Description: "Test Description",
		Notes:       "Test Notes",
	}

	request := NewRequest(defect)
	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)

	// Verify the JSON structure
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	assert.NoError(t, err)

	// Should have "defect" key at top level
	defectData, ok := result["defect"]
	assert.True(t, ok, "Expected 'defect' key in marshaled JSON")

	// Verify defect fields
	defectMap := defectData.(map[string]interface{})
	assert.Equal(t, "DE123", defectMap["FormattedID"])
	assert.Equal(t, "Test Defect", defectMap["Name"])
	assert.Equal(t, "Test Description", defectMap["Description"])
	assert.Equal(t, "Test Notes", defectMap["Notes"])
}

func TestRequest_MarshalJSON_WithObjectReference(t *testing.T) {
	defect := Defect{
		FormattedID: "DE123",
		Name:        "Test Defect",
		C_DefectImpactedArea: &ObjectReference{
			ObjectID: 99999,
		},
	}

	request := NewRequest(defect)
	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)

	// Verify the JSON structure
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	assert.NoError(t, err)

	// Should have "defect" key at top level
	defectData, ok := result["defect"]
	assert.True(t, ok, "Expected 'defect' key in marshaled JSON")

	// Verify defect fields
	defectMap := defectData.(map[string]interface{})
	assert.Equal(t, "DE123", defectMap["FormattedID"])
	assert.Equal(t, "Test Defect", defectMap["Name"])

	// Verify the custom field is properly nested
	customField, ok := defectMap["c_DefectImpactedArea"].(map[string]interface{})
	assert.True(t, ok, "c_DefectImpactedArea should be an object")
	assert.Equal(t, float64(99999), customField["ObjectId"])
}
