package domain

import (
	"encoding/json"
	"fmt"
	"log"
)

func mapToStruct[T any](m map[string]any) T {
	var result T
	jsonBytes, err := json.Marshal(m)

	handleError(err)

	err = json.Unmarshal(jsonBytes, &result)

	handleError(err)

	return result
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("mapToStruct error: %v", err)
	}
}

func PushEventResumer(event Event) string {
	payload := mapToStruct[PushEventPayload](event.Payload)

	commits := len(payload.Commits)

	return fmt.Sprintf("Pushed %d commits to %s", commits, event.Repo.Name)
}

func CreateEventResumer(event Event) string {
	payload := mapToStruct[CreateEventPayload](event.Payload)

	return fmt.Sprintf("A new '%s' was created in %s", payload.RefType, event.Repo.Name)
}

func IssueCommentEventResumer(event Event) string {
	payload := mapToStruct[IssueCommentEventPayload](event.Payload)

	return fmt.Sprintf(
		"%s comments the following in issue '%s' -> %s",
		payload.Comment.User.Login,
		payload.Issue.Title,
		payload.Comment.Body,
	)
}

func PullRequestEventResumer(event Event) string {
	payload := mapToStruct[PullRequestEventPayload](event.Payload)

	return fmt.Sprintf(
		"%s %s a pull request with title '%s'",
		event.Actor.Login,
		payload.Action,
		payload.PullRequest.Title,
	)
}
