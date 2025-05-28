package domain

type Event struct {
	Type    string         `json:"type"`
	Actor   Actor          `json:"actor"`
	Repo    Repo           `json:"repo"`
	Payload map[string]any `json:"payload"`
}

type PushEventPayload struct {
	Commits []any `json:"commits"`
}

type CreateEventPayload struct {
	RefType string `json:"ref_type"`
}

type IssueCommentEventPayload struct {
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
}

type PullRequestEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}
