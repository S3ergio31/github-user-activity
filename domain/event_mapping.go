package domain

const pushEvent = "PushEvent"
const createEvent = "CreateEvent"
const issueCommentEvent = "IssueCommentEvent"
const pullRequestEvent = "PullRequestEvent"

var Reporters = map[string]EventReport{
	pushEvent:         PushEventReporter,
	createEvent:       CreateEventReporter,
	issueCommentEvent: IssueCommentEventReporter,
	pullRequestEvent:  PullRequestEventReporter,
}
