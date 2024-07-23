package repository

import "todo-app/domain/entity"

type WorkspaceRepository interface {
	FindAll() (workspaces []entity.Workspace, err error)
	FindByUserID(userID int) (workspaces []entity.Workspace, err error)
	Create(user entity.Workspace) (workspace entity.Workspace, err error)
}
