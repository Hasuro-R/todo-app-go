package entity

type Workspace struct {
	ID        int
	Title     string
	Emoji     string
	UserID    int
	CreatedAt string
	UpdatedAt string
}

func (en *Workspace) Validate() error {
	return nil
}
