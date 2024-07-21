package entity

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func (u *User) Validate() error {
	return nil
}
