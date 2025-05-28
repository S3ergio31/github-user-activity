package domain

import "time"

type Event struct {
	Id        string         `json:"id"`
	Type      string         `json:"type"`
	Actor     Actor          `json:"actor"`
	Repo      Repo           `json:"repo"`
	Payload   map[string]any `json:"payload"`
	Public    bool           `json:"public"`
	CreatedAt time.Time      `json:"created_at"`
}

type PushEventPayload struct {
	RepositoryId int      `json:"repository_id"`
	PushId       int      `json:"push_id"`
	Size         int      `json:"size"`
	DistinctSize int      `json:"distinct_size"`
	Ref          string   `json:"ref"`
	Head         string   `json:"head"`
	Before       string   `json:"before"`
	Commits      []Commit `json:"commits"`
}

type CreateEventPayload struct {
	Ref          string  `json:"ref"`
	RefType      string  `json:"ref_type"`
	MasterBranch string  `json:"master_branch"`
	Description  *string `json:"description"`
	PusherType   string  `json:"pusher_type"`
}

type IssueCommentEventPayload struct {
	Action  string  `json:"action"`
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
}

type PullRequestEventPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
}
