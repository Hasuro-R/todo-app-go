package persistent

import (
	"database/sql"
	"todo-app/domain/entity"
	"todo-app/domain/repository"
)

type UserInfrastructure struct {
	db *sql.DB
}

func NewUserInfrastructure(db *sql.DB) repository.UserRepository {
	return &UserInfrastructure{db}
}

func (r *UserInfrastructure) FindAll() ([]entity.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users, err := toStructures(rows)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserInfrastructure) FindByID(id int) (entity.User, error) {
	row := r.db.QueryRow(`SELECT * FROM users WHERE id = ?`, id)
	user, err := toStructure(row)
	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}

func toStructure(row *sql.Row) (*entity.User, error) {
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func toStructures(rows *sql.Rows) ([]entity.User, error) {
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
