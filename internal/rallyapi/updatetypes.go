package rallyapi

import (
	"encoding/json"
	"fmt"
	"log"
)

type RallyEntity interface {
	RallyType() string
}

type Request[T RallyEntity] struct {
	Entity T
}

func (r Request[T]) MarshalJSON() ([]byte, error) {
	entityType := r.Entity.RallyType()

	wrapper := map[string]T{
		entityType: r.Entity,
	}
	return json.Marshal(wrapper)
}

func NewRequest[T RallyEntity](entity T) Request[T] {
	return Request[T]{Entity: entity}
}

type Update[T RallyEntity] = Request[T]

// {"OperationResult": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "Errors": [], "Warnings": ["Ignored JSON element defect.c_DefectImpactedArea during processing of this request."], "Object": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/defect/838568308171", "_refObjectUUID": "ed22f7ab-eea3-4b25-b982-419f32ef8d55", "_objectVersion": "7", "_refObjectName": "New defect to analyse", "CreationDate": "2025-12-04T08:13:23.201Z", "_CreatedAt": "today at 3:13 am", "ObjectID": 838568308171, "ObjectUUID": "ed22f7ab-eea3-4b25-b982-419f32ef8d55", "VersionId": "7", "Subscription": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/subscription/21034609048", "_refObjectUUID": "8b5dec86-7178-4b4b-8d52-0c088a49019b", "_refObjectName": "ConnectALL Team", "_type": "Subscription"}, "Workspace": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/workspace/836073279859", "_refObjectUUID": "bb975f8f-99a8-4fcc-be91-9c19ac848f4c", "_refObjectName": "Chris Playground", "_type": "Workspace"}, "AIAssisted": false, "Changesets": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Changesets", "_type": "Changeset", "Count": 0}, "Connections": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Connections", "_type": "Connection", "Count": 0}, "CreatedBy": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/user/21034609179", "_refObjectUUID": "fa4a021a-26f5-42e8-b2a5-2c1409cd4765", "_refObjectName": "connectall@go2group.com", "_type": "User"}, "Description": "\u003Cp\u003EThis is a UI - Automation config issue\u003C/p\u003E", "Discussion": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Discussion", "_type": "ConversationPost", "Count": 0}, "DisplayColor": "#f9a814", "Expedite": false, "FormattedID": "DE201", "LastUpdateDate": "2025-12-04T09:37:57.738Z", "LatestDiscussionAgeInMinutes": null, "Milestones": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Milestones", "_type": "Milestone", "_tagsNameArray": [], "Count": 0}, "Name": "New defect to analyse", "Notes": "", "Owner": null, "Project": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/project/836073294171", "_refObjectUUID": "dd74df03-f66a-4f98-977b-d57df5f2c2be", "_refObjectName": "Project 1", "_type": "Project"}, "Ready": false, "RevisionHistory": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/revisionhistory/838568308173", "_refObjectUUID": "92a0c27c-bdad-4406-baad-6236c0af99c8", "_type": "RevisionHistory"}, "Tags": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Tags", "_type": "Tag", "_tagsNameArray": [], "Count": 0}, "FlowState": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/flowstate/836073294179", "_refObjectUUID": "3d84f314-9f97-4d7d-817a-2e0c60527ad9", "_refObjectName": "Defined", "_type": "FlowState"}, "FlowStateChangedDate": "2025-12-04T08:13:23.201Z", "LastBuild": null, "LastRun": null, "PassingTestCaseCount": 0, "ScheduleState": "Defined", "ScheduleStatePrefix": "D", "TestCaseCount": 0, "AcceptedDate": null, "AffectsDoc": false, "Attachments": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Attachments", "_type": "Attachment", "Count": 0}, "Blocked": false, "BlockedReason": null, "Blocker": null, "ClosedDate": null, "DefectSuites": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/DefectSuites", "_type": "DefectSuite", "Count": 0}, "DragAndDropRank": "P!!!/!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1~F;AWL", "Duplicates": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Duplicates", "_type": "Defect", "Count": 0}, "Environment": "None", "FixedInBuild": null, "FoundInBuild": null, "InProgressDate": null, "Iteration": null, "OpenedDate": null, "Package": null, "PlanEstimate": null, "Priority": "None", "Recycled": false, "Release": null, "ReleaseNote": false, "Requirement": null, "Resolution": "None", "SalesforceCaseID": null, "SalesforceCaseNumber": null, "Severity": "None", "State": "Submitted", "SubmittedBy": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/user/21034609179", "_refObjectUUID": "fa4a021a-26f5-42e8-b2a5-2c1409cd4765", "_refObjectName": "connectall@go2group.com", "_type": "User"}, "TargetBuild": null, "TargetDate": null, "TaskActualTotal": 0.0, "TaskEstimateTotal": 0.0, "TaskRemainingTotal": 0.0, "TaskStatus": "NONE", "Tasks": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/Tasks", "_type": "Task", "Count": 0}, "TestCase": null, "TestCaseResult": null, "TestCaseStatus": "NONE", "TestCases": {"_rallyAPIMajor": "2", "_rallyAPIMinor": "0", "_ref": "https://rally1.rallydev.com/slm/webservice/v2.0/Defect/838568308171/TestCases", "_type": "TestCase", "Count": 0}, "VerifiedInBuild": null, "c_DefectArea": "Core - Update", "_type": "Defect"}}}

type UpdateResponse[T RallyEntity] struct {
	OperationResult OperationResult[T] `json:"OperationResult"`
}

type OperationResult[T RallyEntity] struct {
	Version
	Errors   []string `json:"Errors"`
	Warnings []string `json:"Warnings"`
	Object   T        `json:"Object"`
}

// CheckUpdateResponse checks if the update response contains errors and prints warnings
func CheckUpdateResponse[T RallyEntity](response UpdateResponse[T]) error {
	// Print warnings if any
	if response.HasWarnings() {
		for _, warning := range response.GetWarnings() {
			log.Printf("Rally Warning: %s\n", warning)
		}
	}
	
	// Check for errors
	if len(response.OperationResult.Errors) != 0 {
		return fmt.Errorf("Rally Error: %s", response.OperationResult.Errors[0])
	}
	return nil
}

// GetObject returns the updated object from the response
func (r UpdateResponse[T]) GetObject() T {
	return r.OperationResult.Object
}

// HasWarnings returns true if the response contains warnings
func (r UpdateResponse[T]) HasWarnings() bool {
	return len(r.OperationResult.Warnings) > 0
}

// GetWarnings returns all warnings from the response
func (r UpdateResponse[T]) GetWarnings() []string {
	return r.OperationResult.Warnings
}
