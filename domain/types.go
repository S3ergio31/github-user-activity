package domain

import "time"

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
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

type Event struct {
	Id        string         `json:"id"`
	Type      string         `json:"type"`
	Actor     Actor          `json:"actor"`
	Repo      Repo           `json:"repo"`
	Payload   map[string]any `json:"payload"`
	Public    bool           `json:"public"`
	CreatedAt time.Time      `json:"created_at"`
}
