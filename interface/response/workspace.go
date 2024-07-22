package response

type Workspace struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Emoji     string `json:"emoji"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type WorkspaceList struct {
	Workspaces []Workspace `json:"workspaces"`
}
