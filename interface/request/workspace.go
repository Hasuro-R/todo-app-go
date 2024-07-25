package request

type CreateWorkspace struct {
	Title string `json:"title"`
	Emoji string `json:"emoji"`
}
