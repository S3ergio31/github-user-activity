package domain

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
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

type Issue struct {
	URL                   string      `json:"url"`
	RepositoryURL         string      `json:"repository_url"`
	LabelsURL             string      `json:"labels_url"`
	CommentsURL           string      `json:"comments_url"`
	EventsURL             string      `json:"events_url"`
	HTMLURL               string      `json:"html_url"`
	ID                    int64       `json:"id"`
	NodeID                string      `json:"node_id"`
	Number                int         `json:"number"`
	Title                 string      `json:"title"`
	User                  User        `json:"user"`
	Labels                []Label     `json:"labels"`
	State                 string      `json:"state"`
	Locked                bool        `json:"locked"`
	Assignee              *User       `json:"assignee"`
	Assignees             []User      `json:"assignees"`
	Milestone             *string     `json:"milestone"`
	Comments              int         `json:"comments"`
	CreatedAt             string      `json:"created_at"`
	UpdatedAt             string      `json:"updated_at"`
	ClosedAt              string      `json:"closed_at"`
	AuthorAssociation     string      `json:"author_association"`
	Type                  *string     `json:"type"`
	ActiveLockReason      *string     `json:"active_lock_reason"`
	Draft                 bool        `json:"draft"`
	PullRequest           PullRequest `json:"pull_request"`
	Body                  string      `json:"body"`
	Reactions             Reactions   `json:"reactions"`
	TimelineURL           string      `json:"timeline_url"`
	PerformedViaGitHubApp *string     `json:"performed_via_github_app"`
	StateReason           *string     `json:"state_reason"`
}

type User struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	UserViewType      string `json:"user_view_type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Label struct {
	Name string `json:"name"` // Vacío en este ejemplo, pero puede tener nombre
}

type Reactions struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type Comment struct {
	URL                   string    `json:"url"`
	HTMLURL               string    `json:"html_url"`
	IssueURL              string    `json:"issue_url"`
	ID                    int64     `json:"id"`
	NodeID                string    `json:"node_id"`
	User                  User      `json:"user"`
	CreatedAt             string    `json:"created_at"`
	UpdatedAt             string    `json:"updated_at"`
	AuthorAssociation     string    `json:"author_association"`
	Body                  string    `json:"body"`
	Reactions             Reactions `json:"reactions"`
	PerformedViaGitHubApp *string   `json:"performed_via_github_app"`
}

type PullRequest struct {
	URL                string    `json:"url"`
	ID                 int64     `json:"id"`
	NodeID             string    `json:"node_id"`
	HTMLURL            string    `json:"html_url"`
	DiffURL            string    `json:"diff_url"`
	PatchURL           string    `json:"patch_url"`
	IssueURL           string    `json:"issue_url"`
	Number             int       `json:"number"`
	State              string    `json:"state"`
	Locked             bool      `json:"locked"`
	Title              string    `json:"title"`
	User               User      `json:"user"`
	Body               string    `json:"body"`
	CreatedAt          string    `json:"created_at"`
	UpdatedAt          string    `json:"updated_at"`
	ClosedAt           string    `json:"closed_at"`
	MergedAt           *string   `json:"merged_at"`
	MergeCommitSHA     string    `json:"merge_commit_sha"`
	Assignee           *User     `json:"assignee"`
	Assignees          []User    `json:"assignees"`
	RequestedReviewers []User    `json:"requested_reviewers"`
	RequestedTeams     []Team    `json:"requested_teams"`
	Labels             []Label   `json:"labels"`
	Milestone          *string   `json:"milestone"`
	Draft              bool      `json:"draft"`
	CommitsURL         string    `json:"commits_url"`
	ReviewCommentsURL  string    `json:"review_comments_url"`
	ReviewCommentURL   string    `json:"review_comment_url"`
	CommentsURL        string    `json:"comments_url"`
	StatusesURL        string    `json:"statuses_url"`
	Head               RefBranch `json:"head"`
	Base               RefBranch `json:"base"`
}

type Team struct {
	// Puedes agregar los campos si decides procesarlos
}

type RefBranch struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	SHA   string `json:"sha"`
	User  User   `json:"user"`
	Repo  Repo   `json:"repo"`
}

type Repo struct {
	ID             int64    `json:"id"`
	NodeID         string   `json:"node_id"`
	Name           string   `json:"name"`
	FullName       string   `json:"full_name"`
	Private        bool     `json:"private"`
	HTMLURL        string   `json:"html_url"`
	Description    string   `json:"description"`
	Fork           bool     `json:"fork"`
	URL            string   `json:"url"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
	PushedAt       string   `json:"pushed_at"`
	Homepage       string   `json:"homepage"`
	Size           int      `json:"size"`
	Stargazers     int      `json:"stargazers_count"`
	Watchers       int      `json:"watchers_count"`
	Language       string   `json:"language"`
	HasIssues      bool     `json:"has_issues"`
	HasProjects    bool     `json:"has_projects"`
	HasDownloads   bool     `json:"has_downloads"`
	HasWiki        bool     `json:"has_wiki"`
	HasPages       bool     `json:"has_pages"`
	HasDiscussions bool     `json:"has_discussions"`
	Forks          int      `json:"forks"`
	OpenIssues     int      `json:"open_issues"`
	WatchersCount  int      `json:"watchers"`
	DefaultBranch  string   `json:"default_branch"`
	License        *License `json:"license"`
	Owner          User     `json:"owner"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}
