package response

type Workspace struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Emoji     string `json:"emoji"`
	UserID    int    `json:"userId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type WorkspaceList struct {
	Workspaces []Workspace `json:"workspaces"`
}
