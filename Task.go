package msgraph

import (
	"encoding/json"
	"fmt"
)

type PlannerAssignment struct {
	OrderHint string
}
type Thumbnail struct {
	Height       int
	SourceItemID string
	Url          string
	Width        int
	Content      []byte
}
type ThumbnailSet struct {
	ID     string
	Large  Thumbnail
	Medium Thumbnail
	Small  Thumbnail
	Source Thumbnail
}
type Identity struct {
	DisplayName string
	ID          string
	Thumbnails  ThumbnailSet
}
type IdentitySet struct {
	Application Identity
	Device      Identity
	User        Identity
}

type Plan struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Task represents a single calendar of a user
//
// See https://docs.microsoft.com/en-us/graph/api/resources/plannerTask?view=graph-rest-1.0
type Task struct {
	ID                       string // The task's unique identifier. Read-only.
	ActiveChecklistItemCount int
	AppliedCategories        map[string]bool
	AssigneePriority         string
	Assignments              map[string]PlannerAssignment
	BucketID                 string
	ChecklistItemCount       int
	CompletedBy              IdentitySet
	CompletedDateTime        string
	ConversationThreadID     string
	CreatedBy                IdentitySet
	CreatedDateTime          string
	DueDateTime              string
	HasDescription           bool
	OrderHint                string
	PercentComplete          int
	PlanID                   string
	PreviewType              string
	ReferenceCount           int
	StartDateTime            string
	Title                    string
	graphClient              *GraphClient // the graphClient that created this instance
}

func (t Task) String() string {
	return fmt.Sprintf("Task(ID: \"%v\", Title: \"%v\"", t.ID, t.Title)
}

// setGraphClient sets the graphClient instance in this instance and all child-instances (if any)
func (t *Task) setGraphClient(graphClient *GraphClient) {
	t.graphClient = graphClient
}

// UnmarshalJSON implements the json unmarshal to be used by the json-library
func (t *Task) UnmarshalJSON(data []byte) error {
	tmp := struct {
		ID                       string                       `json:"id"` // The task's unique identifier. Read-only.
		ActiveChecklistItemCount int                          `json:"activeChecklistItemCount"`
		AppliedCategories        map[string]bool              `json:"appliedCategories"`
		AssigneePriority         string                       `json:"assigneePriority"`
		Assignments              map[string]PlannerAssignment `json:"assignments"`
		BucketID                 string                       `json:bucketId"`
		ChecklistItemCount       int                          `json:"checklistItemCount"`
		CompletedBy              IdentitySet                  `json:"completedBy"`
		CompletedDateTime        string                       `json:"completedDateTime"`
		ConversationThreadID     string                       `json:"conversationThreadId"`
		CreatedBy                IdentitySet                  `json:"createdBy"`
		CreatedDateTime          string                       `json:"createdDateTime"`
		DueDateTime              string                       `json:"dueDateTime"`
		HasDescription           bool                         `json:"hasDescription"`
		OrderHint                string                       `json:"orderHint"`
		PercentComplete          int                          `json:"percentComplete"`
		PlanID                   string                       `json:"planId"`
		PreviewType              string                       `json:"previewType"`
		ReferenceCount           int                          `json:"referenceCount"`
		StartDateTime            string                       `json:"startDateTime"`
		Title                    string                       `json:"title"`
	}{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	t.ID = tmp.ID
	t.ActiveChecklistItemCount = tmp.ActiveChecklistItemCount
	t.AppliedCategories = tmp.AppliedCategories
	t.AssigneePriority = tmp.AssigneePriority
	t.Assignments = tmp.Assignments
	t.BucketID = tmp.BucketID
	t.ChecklistItemCount = tmp.ChecklistItemCount
	t.CompletedBy = tmp.CompletedBy
	t.CompletedDateTime = tmp.CompletedDateTime
	t.ConversationThreadID = tmp.ConversationThreadID
	t.CreatedBy = tmp.CreatedBy
	t.CreatedDateTime = tmp.CreatedDateTime
	t.DueDateTime = tmp.DueDateTime
	t.HasDescription = tmp.HasDescription
	t.OrderHint = tmp.OrderHint
	t.PercentComplete = tmp.PercentComplete
	t.PlanID = tmp.PlanID
	t.PreviewType = tmp.PreviewType
	t.ReferenceCount = tmp.ReferenceCount
	t.StartDateTime = tmp.StartDateTime
	t.Title = tmp.Title

	return nil
}
