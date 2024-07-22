package service

import (
	"todo-app/domain/entity"
	"todo-app/domain/repository"
)

type WorkspaceService interface {
	GetAll() (workspaces []entity.Workspace, err error)
	GetByUserID(userID int) (workspaces []entity.Workspace, err error)
}

type WorkspaceServiceImpl struct {
	workspace repository.WorkspaceRepository
}

func NewWorkspaceService(workspace repository.WorkspaceRepository) *WorkspaceServiceImpl {
	return &WorkspaceServiceImpl{
		workspace,
	}
}

func (srv *WorkspaceServiceImpl) GetAll() ([]entity.Workspace, error) {
	workspaces, err := srv.workspace.FindAll()
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}

func (srv *WorkspaceServiceImpl) GetByUserID(userID int) ([]entity.Workspace, error) {
	workspaces, err := srv.workspace.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
