package branch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefectIDFromBranchName(t *testing.T) {
	tests := []struct {
		branch string
		want   string
	}{
		{"feature/DE1234-description", "DE1234"},
		{"fix/de99", "DE99"},
		{"DE1", "DE1"},
		{"no-id-here", ""},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.branch, func(t *testing.T) {
			assert.Equal(t, tt.want, defectIDFromBranchName(tt.branch))
		})
	}
}
