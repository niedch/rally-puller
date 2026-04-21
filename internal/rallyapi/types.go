package rallyapi

type Version struct {
	RallyAPIMajor string `json:"_rallyAPIMajor,omitempty"`
	RallyAPIMinor string `json:"_rallyAPIMinor,omitempty"`
}

type Ref struct {
	Ref           string `json:"_ref,omitempty"`
	RefObjectUUID string `json:"_refObjectUUID,omitempty"`
	ObjectVersion string `json:"_objectVersion,omitempty"`
	RefObjectName string `json:"_refObjectName,omitempty"`
	Type          string `json:"_type,omitempty"`
}

type Response[T any] struct {
	QueryResult QueryResult[T] `json:"QueryResult"`
}

type QueryResult[T any] struct {
	Version
	Errors           []string `json:"Errors"`
	Warnings         []string `json:"Warnings"`
	TotalResultCount int      `json:"TotalResultCount"`
	StartIndex       int      `json:"StartIndex"`
	PageSize         int      `json:"PageSize"`
	Results          []T      `json:"Results"`
}

type Defect struct {
	Version
	Ref
	ObjectID             *int              `json:"ObjectID,omitempty" rallyapi:"ObjectID"`
	FormattedID          string            `json:"FormattedID,omitempty" rallyapi:"FormattedID"`
	Name                 string            `json:"Name,omitempty" rallyapi:"Name"`
	Description          string            `json:"Description,omitempty" rallyapi:"Description"`
	Notes                string            `json:"Notes,omitempty" rallyapi:"Notes"`
	C_DefectImpactedArea *ObjectReference  `json:"c_DefectImpactedArea,omitempty" rallyapi:"c_DefectImpactedArea"`
}

// ObjectReference represents a reference to another Rally object by ObjectID
type ObjectReference struct {
	ObjectID int `json:"ObjectId"`
}

func (d Defect) RallyType() string {
	return "defect"
}

type AllowedAttributeValue struct {
	Version
	Ref
	ObjectID    *int   `json:"ObjectID" rallyapi:"ObjectID"`
	StringValue string `json:"StringValue" rallyapi:"StringValue"`
}
