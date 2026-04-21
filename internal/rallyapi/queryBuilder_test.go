package rallyapi

import (
	"testing"
)

func TestQueryBuilder_FormattedIDContains(t *testing.T) {
	qb := NewQueryBuilder()
	query := qb.WithFormattedId("US123").build()
	expected := "(FormattedID contains \"US123\")"
	if query != expected {
		t.Errorf("Expected: %s, Got: %s", expected, query)
	}
}

func TestQueryBuilder_FormattedIDsContains(t *testing.T) {
	qb := NewQueryBuilder()

	query := qb.WithFormattedIds([]string{"US1", "US2", "US3"}).build()
	expected := "(((FormattedID contains \"US1\") OR (FormattedID contains \"US2\")) OR (FormattedID contains \"US3\"))"
	if query != expected {
		t.Errorf("Expected: %s, Got: %s", expected, query)
	}
}

func TestQueryBuilder_EmptyQuery(t *testing.T) {
	qb := NewQueryBuilder()
	query := qb.build()
	expected := ""
	if query != expected {
		t.Errorf("Expected: %s, Got: %s", expected, query)
	}
}
