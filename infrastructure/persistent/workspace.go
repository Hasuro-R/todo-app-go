package persistent

import (
	"database/sql"
	"todo-app/domain/entity"
	"todo-app/domain/repository"
)

type WorkspacePersistent struct {
	db *sql.DB
}

func NewWorkspacePersistent(db *sql.DB) repository.WorkspaceRepository {
	return &WorkspacePersistent{db}
}

func (r *WorkspacePersistent) FindAll() ([]entity.Workspace, error) {
	rows, err := r.db.Query("SELECT * FROM workspaces")
	if err != nil {
		return nil, err
	}

	workspaces, err := toEntityWorkspaces(rows)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (r *WorkspacePersistent) FindByUserID(userID int) ([]entity.Workspace, error) {
	rows, err := r.db.Query(`SELECT * FROM workspaces WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}

	workspaces, err := toEntityWorkspaces(rows)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func toEntityWorkspaces(rows *sql.Rows) ([]entity.Workspace, error) {
	var workspaces []entity.Workspace
	workspace := entity.Workspace{}
	for rows.Next() {
		err := rows.Scan(&workspace.ID, &workspace.Title, &workspace.Emoji, &workspace.UserID, &workspace.CreatedAt, &workspace.UpdatedAt)
		if err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}
