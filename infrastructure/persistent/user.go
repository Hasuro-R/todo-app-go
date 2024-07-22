package persistent

import (
	"database/sql"
	"todo-app/domain/entity"
	"todo-app/domain/repository"
)

type UserPersistent struct {
	db *sql.DB
}

func NewUserPersistent(db *sql.DB) repository.UserRepository {
	return &UserPersistent{db}
}

func (r *UserPersistent) FindAll() ([]entity.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users, err := toEntityUsers(rows)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPersistent) FindByID(id int) (entity.User, error) {
	row := r.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id)
	user, err := toEntityUser(row)
	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}

func toEntityUser(row *sql.Row) (*entity.User, error) {
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func toEntityUsers(rows *sql.Rows) ([]entity.User, error) {
	var users []entity.User
	user := entity.User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
