package handler

import (
	"net/http"
	"strconv"
	"todo-app/domain/entity"
	"todo-app/interface/response"
	"todo-app/usecase"

	"github.com/go-chi/chi/v5"
)

type WorkspaceHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByUserID(w http.ResponseWriter, r *http.Request)
}

type WorkspaceHandlerImpl struct {
	workspace usecase.WorkspaceUseCase
}

func NewWorkspaceHandler(workspace usecase.WorkspaceUseCase) *WorkspaceHandlerImpl {
	return &WorkspaceHandlerImpl{
		workspace: workspace,
	}
}

func (h *WorkspaceHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	workspaces, err := h.workspace.GetAll()
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toWorkspaceListResponse(workspaces))
}

func (h *WorkspaceHandlerImpl) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		response.Error(w, r, err)
	}

	workspaces, err := h.workspace.GetByUserID(userID)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toWorkspaceListResponse(workspaces))
}

func toWorkspaceResponse(workspace entity.Workspace) response.Workspace {
	return response.Workspace{
		ID:        workspace.ID,
		Title:     workspace.Title,
		Emoji:     workspace.Emoji,
		UserID:    workspace.UserID,
		CreatedAt: workspace.CreatedAt,
		UpdatedAt: workspace.UpdatedAt,
	}
}

func toWorkspaceListResponse(workspaces []entity.Workspace) response.WorkspaceList {
	res := make([]response.Workspace, 0, len(workspaces))
	for i := range workspaces {
		res = append(res, toWorkspaceResponse(workspaces[i]))
	}
	return response.WorkspaceList{
		Workspaces: res,
	}
}
