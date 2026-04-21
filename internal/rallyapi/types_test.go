package rallyapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefect_MarshalWithObjectReference(t *testing.T) {
	defect := Defect{
		FormattedID: "DE123",
		Name:        "Test Defect",
		C_DefectImpactedArea: &ObjectReference{
			ObjectID: 12345,
		},
	}

	jsonBytes, err := json.Marshal(defect)
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	assert.NoError(t, err)

	// Verify the custom field is an object with ObjectId
	customField, ok := result["c_DefectImpactedArea"].(map[string]interface{})
	assert.True(t, ok, "c_DefectImpactedArea should be an object")
	assert.Equal(t, float64(12345), customField["ObjectId"])
}

func TestDefect_MarshalWithoutObjectReference(t *testing.T) {
	defect := Defect{
		FormattedID: "DE123",
		Name:        "Test Defect",
		// C_DefectImpactedArea is nil
	}

	jsonBytes, err := json.Marshal(defect)
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	assert.NoError(t, err)

	// Verify the custom field is omitted when nil
	_, exists := result["c_DefectImpactedArea"]
	assert.False(t, exists, "c_DefectImpactedArea should be omitted when nil")
}
