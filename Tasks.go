package msgraph

import (
	"strings"
)

type Tasks []Task

func (t Tasks) String() string {
	var tasks = make([]string, len(t))
	for i, task := range t {
		tasks[i] = task.String()
	}
	return "Tasks(" + strings.Join(tasks, " | ") + ")"
}

// setGraphClient sets the graphClient instance in this instance and all child-instances (if any)
func (t Tasks) setGraphClient(gC *GraphClient) Tasks {
	for i := range t {
		t[i].setGraphClient(gC)
	}
	return t
}

// GetByName returns the calendar obj of that array whose DisplayName matches
// the given name. Returns an ErrFindCalendar if no calendar exists that matches the given
// name.
func (t Tasks) GetByTitle(title string) (Task, error) {
	for _, task := range t {
		if task.Title == title {
			return task, nil
		}
	}
	return Task{}, ErrFindTask
}
