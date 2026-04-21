package rallyapi

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	queryParts []string
	objectId   string
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		queryParts: make([]string, 0),
	}
}

func (qb *QueryBuilder) WithFormattedId(pattern string) *QueryBuilder {
	qb.queryParts = append(qb.queryParts, fmt.Sprintf("(FormattedID contains \"%s\")", pattern))
	return qb
}

func (qb *QueryBuilder) WithObjectId(objectId string) *QueryBuilder {
	qb.objectId = objectId;
	return qb
}

func (qb *QueryBuilder) WithFormattedIds(formattedIds []string) *QueryBuilder {
	if len(formattedIds) == 0 {
		return qb
	}

	query := fmt.Sprintf(`(FormattedID contains "%s")`, formattedIds[0])
	for i := 1; i < len(formattedIds); i++ {
		query = fmt.Sprintf(`(%s OR (FormattedID contains "%s"))`, query, formattedIds[i])
	}

	qb.queryParts = append(qb.queryParts, query)
	return qb
}

func (qb *QueryBuilder) build() string {
	if len(qb.queryParts) == 0 {
		return ""
	}

	return strings.Join(qb.queryParts, " AND ")
}
