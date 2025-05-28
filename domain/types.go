package domain

type Actor struct {
	Login string `json:"login"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

type Issue struct {
	Title string `json:"title"`
}

type User struct {
	Login string `json:"login"`
}

type Comment struct {
	User User   `json:"user"`
	Body string `json:"body"`
}

type PullRequest struct {
	Title string `json:"title"`
}

type Repo struct {
	Name string `json:"name"`
}
