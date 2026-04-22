package branch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicketIDFromBranchName(t *testing.T) {
	tests := []struct {
		branch string
		want   string
	}{
		{"feature/DE1234-description", "DE1234"},
		{"fix/de99", "DE99"},
		{"DE1", "DE1"},
		{"feature/US1234-description", "US1234"},
		{"fix/us99", "US99"},
		{"US1", "US1"},
		{"feature/S1234-description", "S1234"},
		{"fix/s99", "S99"},
		{"S1", "S1"},
		{"no-id-here", ""},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.branch, func(t *testing.T) {
			assert.Equal(t, tt.want, ticketIDFromBranchName(tt.branch))
		})
	}
}
