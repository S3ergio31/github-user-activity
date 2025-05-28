package domain

const pushEvent = "PushEvent"
const createEvent = "CreateEvent"
const issueCommentEvent = "IssueCommentEvent"
const pullRequestEvent = "PullRequestEvent"

var Resumers = map[string]EventResume{
	pushEvent:         PushEventResumer,
	createEvent:       CreateEventResumer,
	issueCommentEvent: IssueCommentEventResumer,
	pullRequestEvent:  PullRequestEventResumer,
}
