package domain

import (
	"fmt"
)

func PushEventResumer(event Event) string {
	commits := len(event.Payload["commits"].([]any))

	return fmt.Sprintf("Pushed %d commits to %s", commits, event.Repo.Name)
}

func CreateEventResumer(event Event) string {
	refType := event.Payload["ref_type"].(string)

	return fmt.Sprintf("A new '%s' was created in %s", refType, event.Repo.Name)
}
