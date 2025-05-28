package domain

import (
	"fmt"
)

func PushEventReporter(event Event) string {
	payload := ParsePayload[PushEventPayload](event)

	commits := len(payload.Commits)

	return fmt.Sprintf("Pushed %d commits to %s", commits, event.Repo.Name)
}

func CreateEventReporter(event Event) string {
	payload := ParsePayload[CreateEventPayload](event)

	return fmt.Sprintf("A new '%s' was created in %s", payload.RefType, event.Repo.Name)
}

func IssueCommentEventReporter(event Event) string {
	payload := ParsePayload[IssueCommentEventPayload](event)

	return fmt.Sprintf(
		"%s comments the following in issue '%s' -> %s",
		payload.Comment.User.Login,
		payload.Issue.Title,
		payload.Comment.Body,
	)
}

func PullRequestEventReporter(event Event) string {
	payload := ParsePayload[PullRequestEventPayload](event)

	return fmt.Sprintf(
		"%s %s a pull request with title '%s'",
		event.Actor.Login,
		payload.Action,
		payload.PullRequest.Title,
	)
}
