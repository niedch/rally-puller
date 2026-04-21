package rallyapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock test structs - these won't change when you modify actual Rally types
type testEntity struct {
	Field1 string `rallyapi:"Field1"`
	Field2 string `rallyapi:"Field2"`
	Field3 string `rallyapi:"Field3"`
	Field4 string // no tag - should be ignored
}

type testNestedEntity struct {
	Inner testEntity
}

func TestStructReflection(t *testing.T) {
	tests := map[string]struct {
		fieldType any
		fetch     string
	}{
		"simple struct": {
			fieldType: testEntity{},
			fetch:     "Field1,Field2,Field3",
		},
		"nested struct": {
			fieldType: testNestedEntity{},
			fetch:     "Field1,Field2,Field3",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			fetch, err := constructFetch(test.fieldType)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, test.fetch, fetch)
		})
	}
}
